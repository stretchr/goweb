package goweb

import "testing"

func TestControllerFunc(t *testing.T) {
	
	var actualContext *Context = nil
	var expectedContext *Context = new(Context)
	
	// create a controller object using ControllerFunc
	var controller Controller = ControllerFunc(func(c *Context){ actualContext = c })
	
	// use the controller as an object
	controller.HandleRequest(expectedContext)
	
	// ensure the function was called with the correct parameter
	if actualContext != expectedContext {
		t.Errorf(".HandleRequest on controller object (created by ControllerFunc()) didn't call function with correct parameters")
	}
	
}