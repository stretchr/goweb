package goweb

import (
	"os"
	"http"
	"testing"
)

// test route
var routePath string = "/people/{id}/groups/{group_id}.json"
var routePathWithoutExtension string = "/people/{id}/groups/{group_id}"

// expected test route regex
var routePathRegex string = "/people/" + ROUTE_REGEX_PLACEHOLDER + "/groups/" + ROUTE_REGEX_PLACEHOLDER

// Domain used for testing
var testDomain string = "http://test.matryer.com"

/*
	Test controller
*/

// test controller
type TestController struct {
	HandleRequestWasCalled bool
	WriteHeaderWasCalled bool
	LastContext *Context
}
func (handler *TestController) HandleRequest(c *Context) {
	
	// this method has been called
	handler.HandleRequestWasCalled = true
	
	// save the request and response objects
	handler.LastContext = c
	
}
func (handler *TestController) WriteHeader(statusCode int) {
	handler.WriteHeaderWasCalled = true
}

/*
	Test RestController
*/
type TestRestController struct {
	lastCall string
	lastId string
}
func (cr *TestRestController) Create(cx *Context) { cr.lastCall = "Create"; cr.lastId = "(none)" }
func (cr *TestRestController) Delete(id string, cx *Context) { cr.lastCall = "Delete"; cr.lastId = id }
func (cr *TestRestController) DeleteMany(cx *Context) { cr.lastCall = "DeleteMany"; cr.lastId = "(none)" }
func (cr *TestRestController) Read(id string, cx *Context) { cr.lastCall = "Read"; cr.lastId = id }
func (cr *TestRestController) ReadMany(cx *Context) { cr.lastCall = "ReadMany"; cr.lastId = "(none)" }
func (cr *TestRestController) Update(id string, cx *Context) { cr.lastCall = "Update"; cr.lastId = id }
func (cr *TestRestController) UpdateMany(cx *Context) { cr.lastCall = "UpdateMany"; cr.lastId = "(none)" }

/*
	Test ResponseWriter
*/
type TestResponseWriter struct {
	WrittenHeaderInt int
	Output string
	header http.Header
}
func (rw *TestResponseWriter) Header() http.Header {
	
	if rw.header == nil {
		rw.header = make(http.Header)
	}
	
	return rw.header
}
func (rw *TestResponseWriter) Write(bytes []byte) (int, os.Error) {
	
	// add these bytes to the output string
	rw.Output = rw.Output + string(bytes)
	
	// return normal values
	return 0, nil
	
}
func (rw *TestResponseWriter) WriteHeader(i int) {
	rw.WrittenHeaderInt = i
}


func RouteMatcherFunc_Match(c *Context) RouteMatcherFuncValue {
	return Match
}
func RouteMatcherFunc_NoMatch(c *Context) RouteMatcherFuncValue {
	return NoMatch
}
func RouteMatcherFunc_DontCare(c *Context) RouteMatcherFuncValue {
	return DontCare
}




type TestFormatter struct {}

func (f *TestFormatter) Format(input interface{}) ([]uint8, os.Error) {
	return []uint8(""), nil
}
func (f *TestFormatter) ContentType() string {
	return "text/plain"
}




/*

	Test helper functions

*/
func assertEqual(t *testing.T, actual interface{}, expected interface{}, message string) {
	msg := "Objects not equal.\n\tExpected: %s\n\tbut was : %s.\n\t" + message
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}


