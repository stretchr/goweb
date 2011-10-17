package goweb

import (
	"runtime/debug"
	"log"
	"strings"
	"fmt"
	"http"
)

// Wraps a controllerFunc to catch any panics, log them and
// respond with an appropriate error
func safeControllerFunc(controllerFunc func(*Context)) func(*Context) {
	return func(cx *Context) {
		defer func() {
			if err := recover(); err != nil {
				lines := strings.Split(string(debug.Stack()), "\n")[4:]
				log.Print("panic: ", err, "\n", strings.Join(lines, "\n"))
				cx.RespondWithError(500)
			}
		}()
		controllerFunc(cx)
	}
}

// Maps a new route to a controller (with optional RouteMatcherFuncs)
// and returns the new route
func Map(path string, controller Controller, matcherFuncs ...RouteMatcherFunc) *Route {
	return DefaultRouteManager.Map(path, controller, matcherFuncs...)
}

// Maps a new route to a function (with optional RouteMarcherFuncs)
// and returns the new route
func MapFunc(path string, controllerFunc func(*Context), matcherFuncs ...RouteMatcherFunc) *Route {
	return DefaultRouteManager.MapFunc(path, safeControllerFunc(controllerFunc), matcherFuncs...)
}

type MapContext struct {
	pathPrefix string
}

//type RestNester func(*MapContext)

func (mc *MapContext) MapRest(pathPrefix string, controller RestController) {
	pathPrefix = mc.pathPrefix + pathPrefix
	pathPrefixWithId := pathPrefix + `/{id2}`

	// GET /resource1/{id}/resource2/{id2}
	if rc, ok := controller.(NestedRestReader); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Read(c.PathParams["id"], c.PathParams["id2"], c)
		}, GetMethod)
	}
	// GET /resource1/{id}/resource2
	if rc, ok := controller.(NestedRestManyReader); ok {
		fmt.Println("NestedRestManyReader", pathPrefix)
		MapFunc(pathPrefix, func(c *Context) {
			rc.ReadMany(c.PathParams["id"], c)
		}, GetMethod)
	}
	// PUT /resource1/{id}/resource2/{id2}
	if rc, ok := controller.(NestedRestUpdater); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Update(c.PathParams["id"], c.PathParams["id2"], c)
		}, PutMethod)
	}
	// PUT /resource1/{id}/resource2
	if rc, ok := controller.(NestedRestManyUpdater); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.UpdateMany(c.PathParams["id"], c)
		}, PutMethod)
	}
	// DELETE /resource1/{id}/resource2/{id2}
	if rc, ok := controller.(NestedRestDeleter); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Delete(c.PathParams["id"], c.PathParams["id2"], c)
		}, DeleteMethod)
	}
	// DELETE /resource1/{id}/resource2
	if rc, ok := controller.(NestedRestManyDeleter); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.DeleteMany(c.PathParams["id"], c)
		}, DeleteMethod)
	}
	// CREATE /resource1/{id}/resource2
	if rc, ok := controller.(NestedRestCreator); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.Create(c.PathParams["id"], c)
		}, PostMethod)
	}
}

// Maps an entire RESTful set of routes to the specified RestController
// You only have to specify the methods that you require see rest_controller.go
// for the list of interfaces that can be satisfied
func MapRest(pathPrefix string, controller RestController, nested ...func(*MapContext)) {

	var pathPrefixWithId string = pathPrefix + "/{id}"

	// apply nested routes for /resource/{id}
	mc := &MapContext{pathPrefixWithId}
	for _, f := range nested {
		f(mc)
	}

	// GET /resource/{id}
	if rc, ok := controller.(RestReader); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Read(c.PathParams["id"], c)
		}, GetMethod)
	}
	// GET /resource
	if rc, ok := controller.(RestManyReader); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.ReadMany(c)
		}, GetMethod)
	}
	// PUT /resource/{id}
	if rc, ok := controller.(RestUpdater); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Update(c.PathParams["id"], c)
		}, PutMethod)
	}
	// PUT /resource
	if rc, ok := controller.(RestManyUpdater); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.UpdateMany(c)
		}, PutMethod)
	}
	// DELETE /resource/{id}
	if rc, ok := controller.(RestDeleter); ok {
		MapFunc(pathPrefixWithId, func(c *Context) {
			rc.Delete(c.PathParams["id"], c)
		}, DeleteMethod)
	}
	// DELETE /resource
	if rc, ok := controller.(RestManyDeleter); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.DeleteMany(c)
		}, DeleteMethod)
	}
	// CREATE /resource
	if rc, ok := controller.(RestCreator); ok {
		MapFunc(pathPrefix, func(c *Context) {
			rc.Create(c)
		}, PostMethod)
	}
}

// Maps a path to a static directory
func MapStatic(pathPrefix string, rootDirectory string) {
	MapFunc(pathPrefix, func(cx *Context) {
		path := cx.Request.URL.Path
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		path = rootDirectory + path
		fmt.Println("static:", path)
		http.ServeFile(cx.ResponseWriter, cx.Request, path)
	})
}
