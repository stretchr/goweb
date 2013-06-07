package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/controllers"
	"github.com/stretchrcom/goweb/http"
	"github.com/stretchrcom/goweb/paths"
	stewstrings "github.com/stretchrcom/stew/strings"
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
	RestfulIDParameterName string = "id"
)

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
	var matcherFuncs []MatcherFunc
	for i := matcherFuncStartPos; i < len(options); i++ {
		if matcherFunc, ok := options[i].(MatcherFunc); ok {
			matcherFuncs = append(matcherFuncs, matcherFunc)
		} else {
			panic(fmt.Sprintf("goweb: Argument %d (index %d) passed to Map must be of type MatcherFunc but was %s.", i+1, i, options[i]))
		}
	}

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

	var path string
	var controller interface{}

	switch len(options) {
	case 0: // ()
		// no arguments is an error
		panic("goweb: Cannot call MapController with no arguments")
	case 1: // (controller)
		if restfulController, ok := options[0].(controllers.RestfulController); ok {
			controller = restfulController
			path = restfulController.Path()
		} else {
			// use the default path
			controller = options[0]
			path = paths.PathPrefixForClass(options[0])
		}
		break
	case 2: // (path, controller)
		path = options[0].(string)
		controller = options[1]
	}

	// get the specialised paths that we might need
	pathWithID := stewstrings.MergeStrings(path, "/{", RestfulIDParameterName, "}")         // e.g.  people/123
	pathWithOptionalID := stewstrings.MergeStrings(path, "/[", RestfulIDParameterName, "]") // e.g.  people/[123]

	// get the HTTP methods that we will end up mapping
	collectiveMethods := controllers.OptionsListForResourceCollection(controller)
	singularMethods := controllers.OptionsListForSingleResource(controller)

	// BeforeHandler
	if beforeController, ok := controller.(controllers.BeforeHandler); ok {

		// map the collective before handler
		h.MapBefore(collectiveMethods, path, beforeController.Before)

		// map the singular before handler
		h.MapBefore(singularMethods, pathWithID, beforeController.Before)

	}

	// AfterHandler
	if afterController, ok := controller.(controllers.AfterHandler); ok {

		// map the collective after handler
		h.MapAfter(collectiveMethods, path, afterController.After)

		// map the singular after handler
		h.MapAfter(singularMethods, pathWithID, afterController.After)

	}

	// POST /resource  -  Create
	if restfulController, ok := controller.(controllers.RestfulCreator); ok {
		h.Map(http.MethodPost, path, restfulController.Create)
	}

	// GET /resource/{id}  -  Read
	if restfulController, ok := controller.(controllers.RestfulReader); ok {
		h.Map(http.MethodGet, pathWithID, func(ctx context.Context) error {
			return restfulController.Read(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// GET /resource  -  ReadMany
	if restfulController, ok := controller.(controllers.RestfulManyReader); ok {
		h.Map(http.MethodGet, path, restfulController.ReadMany)
	}

	// DELETE /resource/{id}  -  Delete
	if restfulController, ok := controller.(controllers.RestfulDeletor); ok {
		h.Map(http.MethodDelete, pathWithID, func(ctx context.Context) error {
			return restfulController.Delete(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// DELETE /resource  -  DeleteMany
	if restfulController, ok := controller.(controllers.RestfulManyDeleter); ok {
		h.Map(http.MethodDelete, path, restfulController.DeleteMany)
	}

	// PUT /resource/{id}  -  Update
	if restfulController, ok := controller.(controllers.RestfulUpdater); ok {
		h.Map(http.MethodPut, pathWithID, func(ctx context.Context) error {
			return restfulController.Update(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// PUT /resource  -  UpdateMany
	if restfulController, ok := controller.(controllers.RestfulManyUpdater); ok {
		h.Map(http.MethodPut, path, restfulController.UpdateMany)
	}

	// POST /resource/{id}  -  Replace
	if restfulController, ok := controller.(controllers.RestfulReplacer); ok {
		h.Map(http.MethodPost, pathWithID, func(ctx context.Context) error {
			return restfulController.Replace(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// HEAD /resource/[id]  -  Head
	if restfulController, ok := controller.(controllers.RestfulHead); ok {
		h.Map(http.MethodHead, pathWithOptionalID, restfulController.Head)
	}

	// OPTIONS /resource/[id]  -  Options
	if restfulController, ok := controller.(controllers.RestfulOptions); ok {

		h.Map(http.MethodOptions, pathWithOptionalID, restfulController.Options)

	} else {

		// use the default options implementation

		h.Map(http.MethodOptions, path, func(ctx context.Context) error {
			ctx.HttpResponseWriter().Header().Set("Allow", strings.Join(collectiveMethods, ","))
			ctx.HttpResponseWriter().WriteHeader(200)
			return nil
		})

		h.Map(http.MethodOptions, pathWithID, func(ctx context.Context) error {
			ctx.HttpResponseWriter().Header().Set("Allow", strings.Join(singularMethods, ","))
			ctx.HttpResponseWriter().WriteHeader(200)
			return nil
		})

	}

	// everything ok
	return nil

}

// MapStaticFile maps a static file from the specified staticFilePath to the
// specified publicPath.
//
//     goweb.MapStaticFile("favicon.ico", "/location/on/system/to/icons/favicon.ico")
func (h *HttpHandler) MapStaticFile(publicPath, staticFilePath string) (Handler, error) {

	handler, mapErr := h.Map(http.MethodGet, publicPath, func(ctx context.Context) error {

		nethttp.ServeFile(ctx.HttpResponseWriter(), ctx.HttpRequest(), staticFilePath)

		return nil

	})

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
func (h *HttpHandler) MapStatic(publicPath, systemPath string) (Handler, error) {

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

	})

	if mapErr != nil {
		return handler, mapErr
	}

	// set the handler description
	handler.(*PathMatchHandler).Description = fmt.Sprintf("Static files from: %s", systemPath)

	return handler, nil

}
