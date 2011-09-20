package goweb

import (
	"testing"
	"http"
)

/*
	Route
*/

func TestMakeRouteFromPath(t *testing.T) {

	var route *Route
	route = makeRouteFromPath(routePath)

	// test the pattern
	if route.pattern != routePathRegex {
		t.Errorf(".pattern expected to be \"%s\" but was \"%s\"", routePathRegex, route.pattern)
	}

	// test the parameter keys
	var paramKeys ParameterKeyMap = route.parameterKeys

	if (len(paramKeys) != 2) {
		t.Errorf("paramKeys should have 2 items")
	}
	if (paramKeys["id"] != 1) {
		t.Errorf("paramKeys['id'] expected to be 1, but was %s", paramKeys["id"])
	}
	if (paramKeys["group_id"] != 3) {
		t.Errorf("paramKeys['group_id'] expected to be 3, but was %s", paramKeys["id"])
	}

	// test the path
	if route.Path != routePath {
		t.Errorf("Original path should be set")
	}

}

func TestRouteGetParameterValueMapFromPath(t *testing.T) {
	
	var route *Route = makeRouteFromPath(routePath)
	var paramValues ParameterValueMap = route.getParameterValueMap("/people/123/groups/456")
	
	if (len(paramValues) != 2) {
		t.Errorf("paramValues should have 2 items")
	}
	if (paramValues["id"] != "123") {
		t.Errorf("paramKeys['id'] expected to be '123', but was %s", paramValues["id"])
	}
	if (paramValues["group_id"] != "456") {
		t.Errorf("paramKeys['group_id'] expected to be '456', but was %s", paramValues["group_id"])
	}
	
}

func TestRouteDoesMatchPath(t *testing.T) {
	
	var route1 *Route = makeRouteFromPath(routePathWithoutExtension)
	var route2 *Route = makeRouteFromPath("/something-else/{id}/groups/{group_id}")

	if !route1.DoesMatchPath("/people/123/groups/456") {
		t.Errorf("Route should match given path '/people/123/groups/456'")
	}
	if !route1.DoesMatchPath("/people/123/groups/456/comments/13579") {
		t.Errorf("Route should match given path '/people/123/groups/456/comments/13579'")
	}
	if route2.DoesMatchPath("/people/123/groups/456") {
		t.Errorf("Route should NOT match given path")
	}

}

func TestRouteDoesMatchPath_WithExtension(t *testing.T) {
	
	var route1 *Route = makeRouteFromPath(routePath)
	
	if !route1.DoesMatchPath("/people/123/groups/456.json") {
		t.Errorf("Route should match given path '/people/123/groups/456.json'")
	}
	if route1.DoesMatchPath("/people/123/groups/456/comments/13579") {
		t.Errorf("Route should NOT match given path '/people/123/groups/456/comments/13579'")
	}
	if route1.DoesMatchPath("/people/123/groups/456.xml") {
		t.Errorf("Route should NOT match given path '/people/123/groups/456.xml'")
	}

}

func TestCatchAllRouteDoesMatchPath(t *testing.T) {
	
	var route1 *Route = makeRouteFromPath("*")

	if !route1.DoesMatchPath("/people/123/groups/456") {
		t.Errorf("'*' route should match EVERY path")
	}
	if !route1.DoesMatchPath("/") {
		t.Errorf("'*' route should match EVERY path")
	}
	if !route1.DoesMatchPath("") {
		t.Errorf("'*' route should match EVERY path")
	}
	if !route1.DoesMatchPath("/something") {
		t.Errorf("'*' route should match EVERY path")
	}
	if !route1.DoesMatchPath("/this/is/quite/a/long/and/specific/path") {
		t.Errorf("'*' route should match EVERY path")
	}
	
}

func TestRouteDoesMatchContext(t *testing.T) {

	// make a route
	var route1 *Route = makeRouteFromPath(routePath)
	
	// make a test request
	var request *http.Request = new(http.Request)
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	var context *Context = new(Context)
	context.Request = request
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_Match }

	if !route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be true with RouteMatcherFunc_Match")
	}
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_DontCare }

	if route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be false with RouteMatcherFunc_DontCare since default when matcher funcs are present is NOT to match")
	}
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_Match, RouteMatcherFunc_DontCare }

	if !route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be true with RouteMatcherFunc_Match, RouteMatcherFunc_DontCare")
	}
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_NoMatch }

	if route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be false with RouteMatcherFunc_NoMatch")
	}
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare }

	if route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be false with RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare")
	}
	
	route1.MatcherFuncs = []RouteMatcherFunc{ RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare }

	if route1.DoesMatchContext(context) {
		t.Errorf("DoesMatchPath should be false with RouteMatcherFunc_NoMatch, RouteMatcherFunc_DontCare")
	}
	
}


