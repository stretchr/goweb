package goweb

import (
	"testing"
	"net/http"
	"net/url"
)

func handleRequest(path string, httpMethod string) {
	var request *http.Request = new(http.Request)
	request.Method = httpMethod
	request.URL, _ = url.Parse(testDomain + path)
	DefaultHttpHandler.ServeHTTP(new(TestResponseWriter), request)
}
func assertLastId(t *testing.T, controller *TestRestController, expectedLastId string) {
	if controller.lastId != expectedLastId {
		t.Errorf("The last ID should be \"%s\" but was \"%s\".", expectedLastId, controller.lastId)
	}
}
func assertNoLastId(t *testing.T, controller *TestRestController) {
	assertLastId(t, controller, "(none)")
}
func assertLastCall(t *testing.T, controller *TestRestController, expectedLastCall string) {
	if controller.lastCall != expectedLastCall {
		t.Errorf("The last call should be \"%s\" but was \"%s\".", expectedLastCall, controller.lastCall)
	}
}
func assertRoute(t *testing.T, r *Route, expectedPath string, description string, matcherFuncs ...RouteMatcherFunc) {

	if r.Path != expectedPath {
		t.Errorf("route.Path inorrect. Expected \"%s\" but was \"%s\".", expectedPath, r.Path)
	}

	// ensue each matcher function is present
	for _, expected := range matcherFuncs {

		var found bool = false

		for _, existing := range r.MatcherFuncs {
			if existing(nil) == expected(nil) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Matcher function not found: (%s) (for test: %s '%s')", expected, description, expectedPath)
		}

	}

}

func TestMapRest(t *testing.T) {

	DefaultRouteManager.ClearRoutes()

	var testRestController RestController = new(TestRestController)

	MapRest("/people", testRestController)

	if len(DefaultRouteManager.routes) != 7 {
		t.Errorf("There shouldn't be %d route(s)", len(DefaultRouteManager.routes))
	} else {

		var route *Route

		// GET /people/1
		route = DefaultRouteManager.routes[0]
		assertRoute(t, route, "/people/{id}", "GET", GetMethod)
		handleRequest("/people/123", "GET")
		assertLastCall(t, testRestController.(*TestRestController), "Read")
		assertLastId(t, testRestController.(*TestRestController), "123")

		// GET /people
		route = DefaultRouteManager.routes[1]
		assertRoute(t, route, "/people", "GET", GetMethod)
		handleRequest("/people", "GET")
		assertLastCall(t, testRestController.(*TestRestController), "ReadMany")
		assertNoLastId(t, testRestController.(*TestRestController))

		// UPDATE /people/1
		route = DefaultRouteManager.routes[2]
		assertRoute(t, route, "/people/{id}", "PUT", PutMethod)
		handleRequest("/people/123", "PUT")
		assertLastCall(t, testRestController.(*TestRestController), "Update")
		assertLastId(t, testRestController.(*TestRestController), "123")

		// UPDATE /people
		route = DefaultRouteManager.routes[3]
		assertRoute(t, route, "/people", "PUT", PutMethod)
		handleRequest("/people", "PUT")
		assertLastCall(t, testRestController.(*TestRestController), "UpdateMany")
		assertNoLastId(t, testRestController.(*TestRestController))

		// DELETE /people/1
		route = DefaultRouteManager.routes[4]
		assertRoute(t, route, "/people/{id}", "DELETE", DeleteMethod)
		handleRequest("/people/123", "DELETE")
		assertLastCall(t, testRestController.(*TestRestController), "Delete")
		assertLastId(t, testRestController.(*TestRestController), "123")

		// DELETE /people
		route = DefaultRouteManager.routes[5]
		assertRoute(t, route, "/people", "DELETE", DeleteMethod)
		handleRequest("/people", "DELETE")
		assertLastCall(t, testRestController.(*TestRestController), "DeleteMany")
		assertNoLastId(t, testRestController.(*TestRestController))

		// CREATE /people
		route = DefaultRouteManager.routes[6]
		assertRoute(t, route, "/people", "POST", PostMethod)
		handleRequest("/people", "POST")
		assertLastCall(t, testRestController.(*TestRestController), "Create")
		assertNoLastId(t, testRestController.(*TestRestController))

	}

}

func TestMapShortcut(t *testing.T) {

	DefaultRouteManager.ClearRoutes()

	var manager *RouteManager = DefaultRouteManager
	var controller *TestController = new(TestController)

	var createdRoute *Route = Map(routePath, controller, RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare, RouteMatcherFunc_DontCare, RouteMatcherFunc_Match)

	if len(manager.routes) != 1 {
		t.Errorf(".Map should have created a new route (not %d)", len(manager.routes))
	} else {

		// get the route
		var firstRoute *Route = manager.routes[0]

		if firstRoute != createdRoute {
			t.Errorf(".Map should return the same route it adds to .routes")
		}

		if firstRoute.Path != routePath {
			t.Errorf(".Map should have set the right path")
		}

		if firstRoute.Controller != controller {
			t.Errorf(".Map should have set the right controller")
		}

		if len(firstRoute.MatcherFuncs) != 4 {
			t.Errorf("MatcherFuncs should be 4 (not %d)", len(firstRoute.MatcherFuncs))
		} else {

			if firstRoute.MatcherFuncs[0](nil) != RouteMatcherFunc_NoMatch(nil) {
				t.Errorf("firstRoute.MatcherFuncs[0] should be RouteMatcherFunc_NoMatch")
			}
			if firstRoute.MatcherFuncs[1](nil) != RouteMatcherFunc_DontCare(nil) {
				t.Errorf("firstRoute.MatcherFuncs[1] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[2](nil) != RouteMatcherFunc_DontCare(nil) {
				t.Errorf("firstRoute.MatcherFuncs[2] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[3](nil) != RouteMatcherFunc_Match(nil) {
				t.Errorf("firstRoute.MatcherFuncs[3] should be RouteMatcherFunc_Match")
			}

		}

	}

}

func TestMapFuncShortcut(t *testing.T) {

	DefaultRouteManager.ClearRoutes()

	var manager *RouteManager = DefaultRouteManager
	var controller *TestController = new(TestController)
	var functionWrapper func(c *Context) = func(c *Context) {
		controller.HandleRequest(c)
	}

	var createdRoute *Route = MapFunc(routePath, functionWrapper, RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare, RouteMatcherFunc_DontCare, RouteMatcherFunc_Match)

	if len(manager.routes) != 1 {
		t.Errorf(".Map should have created a new route (not %d)", len(manager.routes))
	} else {

		// get the route
		var firstRoute *Route = manager.routes[0]

		if firstRoute != createdRoute {
			t.Errorf(".Map should return the same route it adds to .routes")
		}

		if firstRoute.Path != routePath {
			t.Errorf(".Map should have set the right path")
		}

		if len(firstRoute.MatcherFuncs) != 4 {
			t.Errorf("MatcherFuncs should be 4 (not %d)", len(firstRoute.MatcherFuncs))
		} else {

			if firstRoute.MatcherFuncs[0](nil) != RouteMatcherFunc_NoMatch(nil) {
				t.Errorf("firstRoute.MatcherFuncs[0] should be RouteMatcherFunc_NoMatch")
			}
			if firstRoute.MatcherFuncs[1](nil) != RouteMatcherFunc_DontCare(nil) {
				t.Errorf("firstRoute.MatcherFuncs[1] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[2](nil) != RouteMatcherFunc_DontCare(nil) {
				t.Errorf("firstRoute.MatcherFuncs[2] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[3](nil) != RouteMatcherFunc_Match(nil) {
				t.Errorf("firstRoute.MatcherFuncs[3] should be RouteMatcherFunc_Match")
			}

		}

	}

}
