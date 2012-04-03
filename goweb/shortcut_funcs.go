package goweb

import (
	"runtime/debug"
	"log"
	"strings"
	"fmt"
	"net/http"
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

// Maps an entire RESTful set of routes to the specified RestController
// You only have to specify the methods that you require see rest_controller.go
// for the list of interfaces that can be satisfied
func MapRest(pathPrefix string, controller RestController) {

	var pathPrefixWithId string = pathPrefix + "/{id}"

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
