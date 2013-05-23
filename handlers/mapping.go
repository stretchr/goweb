package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/controllers"
	"github.com/stretchrcom/goweb/http"
	"github.com/stretchrcom/goweb/paths"
	stewstrings "github.com/stretchrcom/stew/strings"
)

var (
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
	var method string
	var path string
	var executor HandlerExecutionFunc

	switch options[0].(type) {
	case string:

		switch options[1].(type) {
		case nil:
			panic("goweb: Cannot call Map with 2nd argument nil.")
		case string: // (method, path, executor, ...)
			method = options[0].(string)
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
	if len(method) > 0 {
		handler.HttpMethods = []string{method}
	}

	// do we have any MatcherFuncs?
	handler.MatcherFuncs = matcherFuncs

	// return the handler
	return handler, nil

}

// Map maps a handler function to a specified path and optional HTTP method.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) Map(options ...interface{}) error {

	handler, err := h.handlerForOptions(options...)

	if err != nil {
		return err
	}

	// append the handler
	h.AppendHandler(handler)

	return nil

}

// Map maps a handler function to a specified path and optional HTTP method 
// to be executed before any other handlers.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) MapBefore(options ...interface{}) error {

	handler, err := h.handlerForOptions(options...)

	if err != nil {
		return err
	}

	// append the handler
	h.AppendPreHandler(handler)

	return nil

}

// Map maps a handler function to a specified path and optional HTTP method 
// to be executed after any other handlers.
//
// For usage information, see goweb.Map.
func (h *HttpHandler) MapAfter(options ...interface{}) error {

	handler, err := h.handlerForOptions(options...)

	if err != nil {
		return err
	}

	// append the handler
	h.AppendPostHandler(handler)

	return nil

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
		break
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

	pathWithID := stewstrings.MergeStrings(path, "/{", RestfulIDParameterName, "}")         // e.g.  people/123
	pathWithOptionalID := stewstrings.MergeStrings(path, "/[", RestfulIDParameterName, "]") // e.g.  people/[123]

	// POST /resource  -  Create
	if restfulController, ok := controller.(controllers.RestfulCreator); ok {
		h.Map(http.MethodPost, path, func(ctx context.Context) error { return restfulController.Create(ctx) })
	}

	// GET /resource/{id}  -  ReadOne
	if restfulController, ok := controller.(controllers.RestfulReader); ok {
		h.Map(http.MethodGet, pathWithID, func(ctx context.Context) error {
			return restfulController.Read(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// GET /resource  -  ReadMany
	if restfulController, ok := controller.(controllers.RestfulManyReader); ok {
		h.Map(http.MethodGet, path, func(ctx context.Context) error { return restfulController.ReadMany(ctx) })
	}

	// DELETE /resource/{id}  -  DeleteOne
	if restfulController, ok := controller.(controllers.RestfulDeletor); ok {
		h.Map(http.MethodDelete, pathWithID, func(ctx context.Context) error {
			return restfulController.Delete(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// DELETE /resource  -  DeleteMany
	if restfulController, ok := controller.(controllers.RestfulManyDeleter); ok {
		h.Map(http.MethodDelete, path, func(ctx context.Context) error {
			return restfulController.DeleteMany(ctx)
		})
	}

	// PUT /resource/{id}  -  Update
	if restfulController, ok := controller.(controllers.RestfulUpdater); ok {
		h.Map(http.MethodPut, pathWithID, func(ctx context.Context) error {
			return restfulController.Update(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// PUT /resource  -  UpdateMany
	if restfulController, ok := controller.(controllers.RestfulManyUpdater); ok {
		h.Map(http.MethodPut, path, func(ctx context.Context) error {
			return restfulController.UpdateMany(ctx)
		})
	}

	// POST /resource/{id}  -  Replace
	if restfulController, ok := controller.(controllers.RestfulReplacer); ok {
		h.Map(http.MethodPost, pathWithID, func(ctx context.Context) error {
			return restfulController.Replace(ctx.PathParams().Get(RestfulIDParameterName).(string), ctx)
		})
	}

	// HEAD /resource/[id]  -  Head
	if restfulController, ok := controller.(controllers.RestfulHead); ok {
		h.Map(http.MethodHead, pathWithOptionalID, func(ctx context.Context) error {
			return restfulController.Head(ctx)
		})
	}

	// HEAD /resource/[id]  -  Options
	if restfulController, ok := controller.(controllers.RestfulOptions); ok {
		h.Map(http.MethodOptions, pathWithOptionalID, func(ctx context.Context) error {
			return restfulController.Options(ctx)
		})
	}

	// everything ok
	return nil

}
