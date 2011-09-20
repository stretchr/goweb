package goweb

// Maps a new route to a controller (with optional RouteMatcherFuncs)
// and returns the new route
func Map(path string, controller Controller, matcherFuncs ...RouteMatcherFunc) *Route {
	return DefaultRouteManager.Map(path, controller, matcherFuncs...)
}

// Maps a new route to a function (with optional RouteMarcherFuncs)
// and returns the new route
func MapFunc(path string, controllerFunc func(*Context), matcherFuncs ...RouteMatcherFunc) *Route {
	return DefaultRouteManager.MapFunc(path, controllerFunc, matcherFuncs...)
}

// Maps an entire RESTful set of routes to the specified RestController
func MapRest(pathPrefix string, controller RestController) {
	
	var pathPrefixWithId string = pathPrefix + "/{id}"
	
	// GET /resource/{id}
	MapFunc(pathPrefixWithId, func(c *Context) {
		controller.Read(c.PathParams["id"], c)
	}, GetMethod)
	
	// GET /resource
	MapFunc(pathPrefix, func(c *Context){
		controller.ReadMany(c)
	}, GetMethod)
	
	// PUT /resource/{id}
	MapFunc(pathPrefixWithId, func(c *Context) {
		controller.Update(c.PathParams["id"], c)
	}, PutMethod)
	
	// PUT /resource
	MapFunc(pathPrefix, func(c *Context){
		controller.UpdateMany(c)
	}, PutMethod)
	
	// DELETE /resource/{id}
	MapFunc(pathPrefixWithId, func(c *Context) {
		controller.Delete(c.PathParams["id"], c)
	}, DeleteMethod)
	
	// DELETE /resource
	MapFunc(pathPrefix, func(c *Context){
		controller.DeleteMany(c)
	}, DeleteMethod)
	
	// CREATE /resource
	MapFunc(pathPrefix, func(c *Context){
		controller.Create(c)
	}, PostMethod);
	
}