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

/*
  Map maps an executor to the specified PathPattern.

	(func [, matcherFuncs])
  (path, func [, matcherFuncs])
  (method, path, func [, matcherFuncs])
*/
func (h *HttpHandler) Map(options ...interface{}) error {

	if len(options) == 0 {
		// no arguments is an error
		panic("goweb: Cannot call Map with no arguments.")
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
		return pathErr
	}

	handler := NewPathMatchHandler(pathPattern, executor)

	// did they specify a method?
	if len(method) > 0 {
		handler.HttpMethods = []string{method}
	}

	// do we have any MatcherFuncs?
	handler.MatcherFuncs = matcherFuncs

	// append the handler
	h.AppendHandler(handler)

	return nil

}

// MapController maps a controller in the handler.
//
// MapController(controller) will map the controller and use the Path() method on the
// controller to determine the path prefix.
//
// MapController(path, controller) will map the controller to the specified path prefix regardless
// of what the Path() methods returns.
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
			// Invalid type for a controller
			panic("goweb: MapController expects a single argument to implement the controllers.RestfulController interface.")
		}
		break
	case 2: // (path, controller)
		path = options[0].(string)
		controller = options[1]
	}

	pathWithID := stewstrings.MergeStrings(path, "/{", RestfulIDParameterName, "}") // e.g.  people/123

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
		h.Map(http.MethodHead, path, func(ctx context.Context) error {
			return restfulController.Head(ctx)
		})
	}

	// HEAD /resource/[id]  -  Options
	if restfulController, ok := controller.(controllers.RestfulOptions); ok {
		h.Map(http.MethodOptions, path, func(ctx context.Context) error {
			return restfulController.Options(ctx)
		})
	}

	// everything ok
	return nil

}
