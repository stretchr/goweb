package goweb

import (
	"testing"
)

/*
	RouteManager
*/

func TestAddRoute(t *testing.T) {
	
	var route1 *Route = makeRouteFromPath(routePath)
	var route2 *Route = makeRouteFromPath("/something-else/{id}/groups/{group_id}")

	manager := new(RouteManager)

	manager.AddRoute(route1)

	if manager.routes.Len() != 1 {
		t.Errorf("len manager.Routes should be 1 not %d", manager.routes.Len())
	}

	manager.AddRoute(route2)

	if manager.routes.Len() != 2 {
		t.Errorf("len manager.Routes should be 2 not %d", manager.routes.Len())
	}

}

func TestMap(t *testing.T) {
	
	manager := new(RouteManager)
	controller := new(TestController)
	
	var createdRoute *Route = manager.Map(routePath, controller)
	
	if len(manager.routes) != 1 {
		t.Errorf(".Map should have created a new route")
	} else {
	
		// get the route
		var firstRoute *Route = manager.routes[0].(*Route)
	
		if firstRoute != createdRoute {
			t.Errorf(".Map should return the same route it adds to .routes")
		}
	
		if firstRoute.Path != routePath {
			t.Errorf(".Map should have set the right path")
		}
	
		if firstRoute.Controller != controller {
			t.Errorf(".Map should have set the right controller")
		}
	
	}
	
}

func TestMapWithMatcherFuncs(t *testing.T) {
	
	manager := new(RouteManager)
	controller := new(TestController)
	
	var createdRoute *Route = manager.Map(routePath, controller, RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare, RouteMatcherFunc_DontCare, RouteMatcherFunc_Match)
	
	if len(manager.routes) != 1 {
		t.Errorf(".Map should have created a new route")
	} else {
	
		// get the route
		var firstRoute *Route = manager.routes[0].(*Route)
	
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
			
			if firstRoute.MatcherFuncs[0] != RouteMatcherFunc_NoMatch {
				t.Errorf("firstRoute.MatcherFuncs[0] should be RouteMatcherFunc_NoMatch")
			}
			if firstRoute.MatcherFuncs[1] != RouteMatcherFunc_DontCare {
				t.Errorf("firstRoute.MatcherFuncs[1] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[2] != RouteMatcherFunc_DontCare {
				t.Errorf("firstRoute.MatcherFuncs[2] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[3] != RouteMatcherFunc_Match {
				t.Errorf("firstRoute.MatcherFuncs[3] should be RouteMatcherFunc_Match")
			}
			
		}
	
	}
	
}

func TestMapFunc(t *testing.T) {
	
	manager := new(RouteManager)
	
	var expectedContext *Context = new(Context)
	var lastCalledContext *Context = nil
	var createdRoute *Route
	
	// map a new route using a function
	createdRoute = manager.MapFunc(routePath, func(c *Context){ lastCalledContext = c }, RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare, RouteMatcherFunc_DontCare, RouteMatcherFunc_Match)
	
	// make the contorller handle the request
	createdRoute.Controller.HandleRequest(expectedContext)
	
	// make sure the function was called
	if expectedContext != lastCalledContext {
		t.Errorf("Mapped function was not called by HandleRequst with correct parameter")
	}
	
	// and check the other stuff about the route
	if len(manager.routes) != 1 {
		
		t.Errorf(".Map should have created a new route")
		
	} else {

		// get the route
		var firstRoute *Route = manager.routes[0].(*Route)

		if firstRoute != createdRoute {
			t.Errorf(".Map should return the same route it adds to .routes")
		}

		if firstRoute.Path != routePath {
			t.Errorf(".Map should have set the right path")
		}
	
		if len(firstRoute.MatcherFuncs) != 4 {
			t.Errorf("MatcherFuncs should be 4 (not %d)", len(firstRoute.MatcherFuncs))
		} else {
		
			if firstRoute.MatcherFuncs[0] != RouteMatcherFunc_NoMatch {
				t.Errorf("firstRoute.MatcherFuncs[0] should be RouteMatcherFunc_NoMatch")
			}
			if firstRoute.MatcherFuncs[1] != RouteMatcherFunc_DontCare {
				t.Errorf("firstRoute.MatcherFuncs[1] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[2] != RouteMatcherFunc_DontCare {
				t.Errorf("firstRoute.MatcherFuncs[2] should be RouteMatcherFunc_DontCare")
			}
			if firstRoute.MatcherFuncs[3] != RouteMatcherFunc_Match {
				t.Errorf("firstRoute.MatcherFuncs[3] should be RouteMatcherFunc_Match")
			}
		
		}

	}
	
}

func TestClearRoutes(t *testing.T) {
	
	var rm *RouteManager = new(RouteManager)
	
	var route1 *Route = makeRouteFromPath("/something-else/{id}/groups/{group_id}")
	var route2 *Route = makeRouteFromPath(routePath)

	rm.AddRoute(route1)
	rm.AddRoute(route2)
	
	rm.ClearRoutes()
	
	if len(rm.routes) != 0 {
		t.Errorf("ClearRoutes should have cleared the .routes")
	}
	
}
