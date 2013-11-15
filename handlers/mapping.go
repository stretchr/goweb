package handlers

import (
	"fmt"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/controllers"
	"github.com/stretchr/goweb/http"
	"github.com/stretchr/goweb/paths"
	stewstrings "github.com/stretchr/stew/strings"
	nethttp "net/http"
	"strings"
)

var (
	// RestfulIDParameterName is the name Goweb will use when mapping the ID
	// parameter in RESTful mappings.
	//
	//     ctx.PathParam(RestfulIDParameterName)
	//
	// By default, "id" is used.
	//
	// To specify you own, change this value before mapping your
	// controllers.
	//
	//     handlers.RestfulIDParameterName = "resourceId"
	//
	RestfulIDParameterName string = "id"
)

// Look for supported MatcherFunc option types and return a full list of
// all MatcherFuncs found.
func findMatcherFuncs(options ...interface{}) []MatcherFunc {
	var matcherFuncs []MatcherFunc
	for i := 0; i < len(options); i++ {
		switch options[i].(type) {
		case func(context.Context) (MatcherFuncDecision, error):
			matcher := options[i].(func(context.Context) (MatcherFuncDecision, error))
			matcherFuncs = append(matcherFuncs, MatcherFunc(matcher))
		case MatcherFunc:
			matcher := options[i].(MatcherFunc)
			matcherFuncs = append(matcherFuncs, matcher)
		case []MatcherFunc:
			matchers := options[i].([]MatcherFunc)
			matcherFuncs = append(matcherFuncs, matchers...)
		default:
			panic(fmt.Sprintf("goweb: Argument %d (index %d) passed to Map must be of type MatcherFunc or []MatcherFunc, but was %s.", i+1, i, options[i]))
		}
	}
	return matcherFuncs
}

// handlerForOptions gets or creates a Handler object based on the specified
// options.  See goweb.Map for details of valid options.
func (h *HttpHandler) handlerForOptions(options ...interface{}) (Handler, error) {

	if len(options) == 0 {
		// no arguments is an error
		panic("goweb: Cannot call Map functions with no arguments.")
	}

	var matcherFuncStartPos int = -1
	var methods []string
	var path string
	var executor HandlerExecutionFunc

	switch options[0].(type) {
	case string, []string:

		switch options[1].(type) {
		case nil:
			panic("goweb: Cannot call Map with 2nd argument nil.")
		case string: // (method|methods, path, executor, ...)

			// get the methods from the arguments
			switch options[0].(type) {
			case []string:
				methods = options[0].([]string)
			case string:
				methods = []string{options[0].(string)}
			}

			path = options[1].(string)
			executor = options[2].(func(context.Context) error)
			matcherFuncStartPos = 3
		default: // (path, executor, ...)
			path = options[0].(string)
			executor = options[1].(func(context.Context) error)
			matcherFuncStartPos = 2
		}
	case Handler: // actual handler object
		return options[0].(Handler), nil
	default: // (executor)
		matcherFuncStartPos = 1
		path = "***"
		executor = options[0].(func(context.Context) error)
	}

	// collect the matcher funcs
	var matcherFuncs []MatcherFunc = findMatcherFuncs(options[matcherFuncStartPos:]...)

	pathPattern, pathErr := paths.NewPathPattern(path)

	if pathErr != nil {
		return nil, pathErr
	}

	handler := NewPathMatchHandler(pathPattern, executor)

	// did they specify a method?
	if len(methods) > 0 {
		handler.HttpMethods = methods
	}

	// do we have any MatcherFuncs?
	handler.MatcherFuncs = matcherFuncs

	// return the handler
	return handler, nil

}

// Map maps a handler function to a specified path and optional HTTP method.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) Map(options ...interface{}) (Handler, error) {

	handler, err := h.handlerForOptions(options...)

	// normally mapped handlers should only execute one,
	// then it should break
	if pathMatchHandler, ok := handler.(*PathMatchHandler); ok {
		pathMatchHandler.BreakCurrentPipeline = true
	}

	if err != nil {
		return nil, err
	}

	// append the handler
	h.AppendHandler(handler)

	return handler, nil

}

// Map maps a handler function to a specified path and optional HTTP method
// to be executed before any other handlers.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) MapBefore(options ...interface{}) (Handler, error) {

	handler, err := h.handlerForOptions(options...)

	if err != nil {
		return nil, err
	}

	// append the handler
	h.AppendPreHandler(handler)

	return handler, nil

}

// Map maps a handler function to a specified path and optional HTTP method
// to be executed after any other handlers.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) MapAfter(options ...interface{}) (Handler, error) {

	handler, err := h.handlerForOptions(options...)

	if err != nil {
		return nil, err
	}

	// append the handler
	h.AppendPostHandler(handler)

	return handler, nil

}

// MapController maps a controller to a specified path prefix.
//
// For more information, see goweb.MapController.
func (h *HttpHandler) MapController(options ...interface{}) error {

	var matcherFuncStartPos int = -1
	var path string
	var controller interface{}

	if len(options) == 0 {
		// no arguments is an error
		panic("goweb: Cannot call MapController with no arguments")
	}

	switch options[0].(type) {
	case string: // (path, controller)
		if len(options) == 1 {
			// we need more than just a string
			panic("goweb: Cannot call MapController without a Controller")
		}
		path = options[0].(string)
		controller = options[1]
		matcherFuncStartPos = 2
	default: // (controller)
		if restfulController, ok := options[0].(controllers.RestfulController); ok {
			controller = restfulController
			path = restfulController.Path()
		} else {
			// use the default path
			controller = options[0]
			path = paths.PathPrefixForClass(options[0])
		}
		matcherFuncStartPos = 1
	}

	// store the matcher function slice
	var matcherFuncs []MatcherFunc = findMatcherFuncs(options[matcherFuncStartPos:]...)

	// get the specialised paths that we might need
	pathWithID := stewstrings.MergeStrings(path, "/{", RestfulIDParameterName, "}")         // e.g.  people/123
	pathWithOptionalID := stewstrings.MergeStrings(path, "/[", RestfulIDParameterName, "]") // e.g.  people/[123]

	// get the HTTP methods that we will end up mapping
	collectiveMethods := optionsListForResourceCollection(h, controller)
	singularMethods := optionsListForSingleResource(h, controller)

	// BeforeHandler
	if beforeController, ok := controller.(controllers.BeforeHandler); ok {

		// map the collective before handler
		h.MapBefore(collectiveMethods, path, beforeController.Before, matcherFuncs)

		// map the singular before handler
		h.MapBefore(singularMethods, pathWithID, beforeController.Before, matcherFuncs)

	}

	// AfterHandler
	if afterController, ok := controller.(controllers.AfterHandler); ok {

		// map the collective after handler
		h.MapAfter(collectiveMethods, path, afterController.After, matcherFuncs)

		// map the singular after handler
		h.MapAfter(singularMethods, pathWithID, afterController.After, matcherFuncs)

	}

	// POST /resource  -  Create
	if restfulController, ok := controller.(controllers.RestfulCreator); ok {
		h.Map(h.HttpMethodForCreate, path, restfulController.Create, matcherFuncs)
	}

	// GET /resource/{id}  -  Read
	if restfulController, ok := controller.(controllers.RestfulReader); ok {
		h.Map(h.HttpMethodForReadOne, pathWithID, func(ctx context.Context) error {
			return restfulController.Read(ctx.PathParams().Get(RestfulIDParameterName).Str(), ctx)
		}, matcherFuncs)
	}

	// GET /resource  -  ReadMany
	if restfulController, ok := controller.(controllers.RestfulManyReader); ok {
		h.Map(h.HttpMethodForReadMany, path, restfulController.ReadMany, matcherFuncs)
	}

	// DELETE /resource/{id}  -  Delete
	if restfulController, ok := controller.(controllers.RestfulDeletor); ok {
		h.Map(h.HttpMethodForDeleteOne, pathWithID, func(ctx context.Context) error {
			return restfulController.Delete(ctx.PathParams().Get(RestfulIDParameterName).Str(), ctx)
		}, matcherFuncs)
	}

	// DELETE /resource  -  DeleteMany
	if restfulController, ok := controller.(controllers.RestfulManyDeleter); ok {
		h.Map(h.HttpMethodForDeleteMany, path, restfulController.DeleteMany, matcherFuncs)
	}

	// PATCH /resource/{id}  -  Update
	if restfulController, ok := controller.(controllers.RestfulUpdater); ok {
		h.Map(h.HttpMethodForUpdateOne, pathWithID, func(ctx context.Context) error {
			return restfulController.Update(ctx.PathParams().Get(RestfulIDParameterName).Str(), ctx)
		}, matcherFuncs)
	}

	// PATCH /resource  -  UpdateMany
	if restfulController, ok := controller.(controllers.RestfulManyUpdater); ok {
		h.Map(h.HttpMethodForUpdateMany, path, restfulController.UpdateMany, matcherFuncs)
	}

	// PUT /resource/{id}  -  Replace
	if restfulController, ok := controller.(controllers.RestfulReplacer); ok {
		h.Map(h.HttpMethodForReplace, pathWithID, func(ctx context.Context) error {
			return restfulController.Replace(ctx.PathParams().Get(RestfulIDParameterName).Str(), ctx)
		}, matcherFuncs)
	}

	// HEAD /resource/[id]  -  Head
	if restfulController, ok := controller.(controllers.RestfulHead); ok {
		h.Map(h.HttpMethodForHead, pathWithOptionalID, restfulController.Head, matcherFuncs)
	}

	// OPTIONS /resource/[id]  -  Options
	if restfulController, ok := controller.(controllers.RestfulOptions); ok {

		h.Map(h.HttpMethodForOptions, pathWithOptionalID, restfulController.Options, matcherFuncs)

	} else {

		// use the default options implementation

		h.Map(http.MethodOptions, path, func(ctx context.Context) error {
			ctx.HttpResponseWriter().Header().Set("Allow", strings.Join(collectiveMethods, ","))
			ctx.HttpResponseWriter().WriteHeader(200)
			return nil
		}, matcherFuncs)

		h.Map(http.MethodOptions, pathWithID, func(ctx context.Context) error {
			ctx.HttpResponseWriter().Header().Set("Allow", strings.Join(singularMethods, ","))
			ctx.HttpResponseWriter().WriteHeader(200)
			return nil
		}, matcherFuncs)

	}

	// everything ok
	return nil

}

// MapStaticFile maps a static file from the specified staticFilePath to the
// specified publicPath.
//
//     goweb.MapStaticFile("favicon.ico", "/location/on/system/to/icons/favicon.ico")
func (h *HttpHandler) MapStaticFile(publicPath, staticFilePath string, matcherFuncs ...MatcherFunc) (Handler, error) {

	handler, mapErr := h.Map(http.MethodGet, publicPath, func(ctx context.Context) error {

		nethttp.ServeFile(ctx.HttpResponseWriter(), ctx.HttpRequest(), staticFilePath)

		return nil

	}, matcherFuncs)

	if mapErr != nil {
		return handler, mapErr
	}

	// set the handler description
	handler.(*PathMatchHandler).Description = fmt.Sprintf("Static file from: %s", staticFilePath)

	return nil, nil

}

// MapStatic maps static files from the specified systemPath to the
// specified publicPath.
//
//     goweb.MapStatic("/static", "/location/on/system/to/files")
func (h *HttpHandler) MapStatic(publicPath, systemPath string, matcherFuncs ...MatcherFunc) (Handler, error) {

	path := paths.NewPath(publicPath)
	var dynamicPath string = path.RawPath

	// ensure the path ends in "***"
	segments := path.Segments()
	if segments[len(segments)-1] != "***" {
		dynamicPath = fmt.Sprintf("%s/***", path.RawPath)
	}

	handler, mapErr := h.Map(http.MethodGet, dynamicPath, func(ctx context.Context) error {

		// get the non-system part of the path
		thePath := path.RealFilePath(systemPath, ctx.Path().RawPath)

		nethttp.ServeFile(ctx.HttpResponseWriter(), ctx.HttpRequest(), thePath)

		return nil

	}, matcherFuncs)

	if mapErr != nil {
		return handler, mapErr
	}

	// set the handler description
	handler.(*PathMatchHandler).Description = fmt.Sprintf("Static files from: %s", systemPath)

	return handler, nil

}
