package goweb

/*
	RouteManager
*/

// Manages routes and matching
type RouteManager struct {
	routes []*Route
}

// Creates a route that maps the specified path to the specified controller
// along with any optional RouteMatcherFunc modifiers
func (manager *RouteManager) Map(path string, controller Controller, matcherFuncs ...RouteMatcherFunc) *Route {

	// create the route (from the path)
	route := makeRouteFromPath(path)

	// set the controller
	route.Controller = controller

	// set the matcher funcs
	route.MatcherFuncs = matcherFuncs

	// add the new route to the default 
	manager.AddRoute(route)

	// return the new route
	return route

}

func (manager *RouteManager) MapFunc(path string, contorllerFunction func(*Context), matcherFuncs ...RouteMatcherFunc) *Route {

	// create the route (from the path)
	route := makeRouteFromPath(path)

	// set the controller
	route.Controller = ControllerFunc(contorllerFunction)

	// set the matcher funcs
	route.MatcherFuncs = matcherFuncs

	// add the new route to the default 
	manager.AddRoute(route)

	// return the new route
	return route

}

// Adds a route to the manager
func (manager *RouteManager) AddRoute(route *Route) {
	manager.routes = append(manager.routes, route)
}

// Clears all routes
func (manager *RouteManager) ClearRoutes() {
	manager.routes = make([]*Route, 0)
}

// Default instance of the RouteManager
var DefaultRouteManager *RouteManager = new(RouteManager)
