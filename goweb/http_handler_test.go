package goweb

import (
	"testing"
	"http"
	"strings"
)

/*
	HttpHandler
*/

func TestGetMathingRoute(t *testing.T) {
	
	// make a test request
	var testRequest *http.Request = new(http.Request)
	
	DefaultRouteManager.ClearRoutes()
	
	var route1 *Route = makeRouteFromPath("/people/{id}")
	var route2 *Route = makeRouteFromPath("/people/{person_id}/groups/{group_id}")
	var route3 *Route = makeRouteFromPath("*")
	
	DefaultRouteManager.AddRoute(route1)
	DefaultRouteManager.AddRoute(route2)
	DefaultRouteManager.AddRoute(route3)
	
	testRequest.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	_, route, _ := DefaultHttpHandler.GetMathingRoute(nil, testRequest)
	
	if route != route1 {
		t.Errorf("Route1 expected")
	}
	
}

func TestGetMatchingRoute_WithMethodOverrideParameter(t *testing.T) {
	
	var lastMethod string = ""
	
	DefaultRouteManager.ClearRoutes()
	
	MapFunc("/api", func(c *Context){
		lastMethod = "GET"
	}, GetMethod)
	MapFunc("/api", func(c *Context){
		lastMethod = "POST"
	}, PostMethod)
	MapFunc("/api", func(c *Context){
		lastMethod = "PUT"
	}, PutMethod)
	MapFunc("/api", func(c *Context){
		lastMethod = "DELETE"
	}, DeleteMethod)
	
	// make test objects
	var testResponse http.ResponseWriter = new(TestResponseWriter)
	var testRequest *http.Request = new(http.Request)
	var url *http.URL
	
	testRequest.Method = "GET"
	
	// handle the request
	url, _ = http.ParseURL(testDomain + "/api")
	testRequest, _ = http.NewRequest("GET", url.Raw, nil)
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	if lastMethod != "GET" {
		t.Errorf("ServeHTTP with no method override parameter should the actual HTTP method.  GET expected but was '%s'", lastMethod)
	}
	
	url, _ = http.ParseURL(testDomain + "/api?" + REQUEST_METHOD_OVERRIDE_PARAMETER + "=post")
	testRequest, _ = http.NewRequest("GET", url.Raw, nil)
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	if lastMethod != "POST" {
		t.Errorf("ServeHTTP with method override parameter should use that method instead.  POST expected but was '%s'", lastMethod)
	}
	
	url, _ = http.ParseURL(testDomain + "/api?" + REQUEST_METHOD_OVERRIDE_PARAMETER + "=Put")
	testRequest, _ = http.NewRequest("GET", url.Raw, nil)
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	if lastMethod != "PUT" {
		t.Errorf("ServeHTTP with method override parameter should use that method instead.  PUT expected but was '%s'", lastMethod)
	}
	
	url, _ = http.ParseURL(testDomain + "/api?" + REQUEST_METHOD_OVERRIDE_PARAMETER + "=DELETE")
	testRequest, _ = http.NewRequest("GET", url.Raw, nil)
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	if lastMethod != "DELETE" {
		t.Errorf("ServeHTTP with method override parameter should use that method instead.  DELETE expected but was '%s'", lastMethod)
	}
	
}

func TestServeHTTP(t *testing.T) {
	
	// clear routes
	DefaultRouteManager.ClearRoutes()

	// make some routes
	var route1 *Route = makeRouteFromPath("/people/{id}/groups/{group_id}")
	var route2 *Route = makeRouteFromPath("/animals/{id}/groups/{group_id}")
	DefaultRouteManager.AddRoute(route1)
	DefaultRouteManager.AddRoute(route2)
	
	// setup controllers
	var controller1 *TestController = new(TestController)
	var controller2 *TestController = new(TestController)

	// bind the route to the controller
	route1.Controller = controller1
	route2.Controller = controller2

	// make test objects
	var testResponse http.ResponseWriter = new(TestResponseWriter)
	var testRequest *http.Request = new(http.Request)
	
	// set the request URL
	testRequest.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	
	// handle the request
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)

	// make sure the routes were parsed correctly and
	// the relevant functions called
	if !controller1.HandleRequestWasCalled {
		t.Errorf("HandleRequest on controller1 should have been called")
	} else {

		if controller1.LastContext.ResponseWriter != testResponse {
			t.Errorf("Incorrect ResponseWriter passed to HandleRequest of controller1")
		}
		if controller1.LastContext.Request != testRequest {
			t.Errorf("Incorrect Request (%s) passed to HandleRequest of controller1. Should be (%s).", controller1.LastContext.Request, testRequest)
		}
	
		// ensure the path parameters are correct
		if controller1.LastContext.PathParams["id"] != "123" {
			t.Errorf("Context.PathParameters['id'] should be '123' not '%s'", controller1.LastContext.PathParams["id"])
		}
		if controller1.LastContext.PathParams["group_id"] != "456" {
			t.Errorf("Context.PathParameters['group_id'] should be '456' not '%s'", controller1.LastContext.PathParams["group_id"])
		}

	}

}

func TestNoControllerReturnsError(t *testing.T) {
	
	DefaultRouteManager.ClearRoutes()
	
	var route1 *Route = makeRouteFromPath("/people/{id}/groups/{group_id}")
	
	// use NO controller
	route1.Controller = nil
	
	// add the route
	DefaultRouteManager.AddRoute(route1)
	
	// make test objects
	var testResponse *TestResponseWriter = new(TestResponseWriter)
	var testRequest *http.Request = new(http.Request)
	
	// set the request URL
	testRequest.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	
	// handle the request
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	
	if testResponse.WrittenHeaderInt != http.StatusInternalServerError {
		t.Errorf("Errors should return 500 status code")
	}
	
	if strings.Index(testResponse.Output, ERR_STANDARD_PREFIX) == -1 {
		t.Errorf("Output expected to contain ERR_STANDARD_PREFIX. '%s'", testResponse.Output)
	}
	if strings.Index(testResponse.Output, ERR_NO_CONTROLLER) == -1 {
		t.Errorf("Output expected to contain ERR_NO_CONTROLLER. '%s'", testResponse.Output)
	}
	
}

func TestNoRouteFoundError(t *testing.T) {
	
	DefaultRouteManager.ClearRoutes()
	
	// make test objects
	var testResponse *TestResponseWriter = new(TestResponseWriter)
	var testRequest *http.Request = new(http.Request)
	
	// set the request URL
	testRequest.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	
	// handle the request
	DefaultHttpHandler.ServeHTTP(testResponse, testRequest)
	
	if testResponse.WrittenHeaderInt != http.StatusInternalServerError {
		t.Errorf("Errors should return 500 status code")
	}

	if strings.Index(testResponse.Output, ERR_STANDARD_PREFIX) == -1 {
		t.Errorf("Output expected to contain ERR_STANDARD_PREFIX. '%s'", testResponse.Output)
	}
	if strings.Index(testResponse.Output, ERR_NO_MATCHING_ROUTE) == -1 {
		t.Errorf("Output expected to contain ERR_NO_MATCHING_ROUTE. '%s'", testResponse.Output)
	}
	
}


