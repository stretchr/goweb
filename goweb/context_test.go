package goweb

import (
	"net/http"
	"net/url"
	"testing"
)

func (c *Context) assertContentType(t *testing.T, contentType string) {

	headers := c.ResponseWriter.Header()
	if headers["Content-Type"] == nil {
		t.Errorf("Content-Type expected to be '%s' but is missing :-(", contentType)
	} else if headers["Content-Type"][0] != contentType {
		t.Errorf("Content-Type expected to be '%s' not '%s'.", contentType, headers["Content-Type"][0])
	}

}

func MakeTestContext() *Context {

	ClearFormatters()
	ConfigureDefaultFormatters()

	var request *http.Request = new(http.Request)
	var responseWriter http.ResponseWriter = new(TestResponseWriter)
	var pathParams ParameterValueMap = make(ParameterValueMap)

	return makeContext(request, responseWriter, pathParams)
}

func MakeTestContextWithUrl(u string) *Context {
	ClearFormatters()
	ConfigureDefaultFormatters()

	var request *http.Request = new(http.Request)
	request.URL, _ = url.Parse(u)
	var responseWriter http.ResponseWriter
	var pathParams ParameterValueMap = make(ParameterValueMap)

	return makeContext(request, responseWriter, pathParams)
}

func TestMakeContext(t *testing.T) {

	var request *http.Request = new(http.Request)
	var responseWriter http.ResponseWriter
	var pathParams ParameterValueMap = make(ParameterValueMap)
	pathParams["something"] = "this"

	context := makeContext(request, responseWriter, pathParams)

	if context.Request != request {
		t.Errorf("context.Request incorrect")
	}
	if context.ResponseWriter != responseWriter {
		t.Errorf("context.ResponseWriter incorrect")
	}
	if context.PathParams["something"] != "this" {
		t.Errorf("context.PathParams incorrect")
	}

}

func TestContextFormat(t *testing.T) {

	var context *Context

	context = MakeTestContextWithUrl(testDomain + "/people/123.json")
	if context.Format != JSON_FORMAT {
		t.Errorf("Format should be JSON not %s", context.Format)
	}

	context = MakeTestContextWithUrl(testDomain + "/people/123.xml")
	if context.Format != XML_FORMAT {
		t.Errorf("Format should be XML")
	}

	context = MakeTestContextWithUrl(testDomain + "/people/123.html")
	if context.Format != HTML_FORMAT {
		t.Errorf("Format should be HTML")
	}

	context = MakeTestContextWithUrl(testDomain + "/people/123")
	if context.Format != DEFAULT_FORMAT {
		t.Errorf("Format should be the default format")
	}
}

func TestWriteResponsePassesTheRightThingsToTheFormatter(t *testing.T) {

	context := MakeTestContextWithUrl(testDomain + "/people/123.json")
	ClearFormatters()
	context.ResponseWriter = new(TestResponseWriter)

	data := "This is the data"

	// create a test formatter that will always be used
	testFormatter := new(TestFormatter)
	AddFormatter(testFormatter)

	// write something
	context.WriteResponse(data, 200)

	if testFormatter.LastContext != context {
		t.Errorf("Correct context object was not passed to the formatter. ", testFormatter.LastContext, context)
	}

	if testFormatter.LastInput != data {
		t.Error("Correct input (data) object was not passed to the formatter")
	}

}

/*

	Quick method checking helpers

*/

func AssertNotGet(c *Context, t *testing.T) {
	if c.IsGet() {
		t.Errorf("IsGet should be false for '%s' method.", c.Request.Method)
	}
}
func AssertNotPost(c *Context, t *testing.T) {
	if c.IsPost() {
		t.Errorf("IsPost should be false for '%s' method.", c.Request.Method)
	}
}
func AssertNotPut(c *Context, t *testing.T) {
	if c.IsPut() {
		t.Errorf("IsPut should be false for '%s' method.", c.Request.Method)
	}
}
func AssertNotDelete(c *Context, t *testing.T) {
	if c.IsDelete() {
		t.Errorf("IsDelete should be false for '%s' method.", c.Request.Method)
	}
}
func AssertNotOptions(c *Context, t *testing.T) {
	if c.IsOptions() {
		t.Errorf("IsOptions should be false for '%s' method.", c.Request.Method)
	}
}
func AssertNotHead(c *Context, t *testing.T) {
	if c.IsHead() {
		t.Errorf("IsHead should be false for '%s' method.", c.Request.Method)
	}
}
func AssertGet(c *Context, t *testing.T) {
	if !c.IsGet() {
		t.Errorf("IsGet should be true for '%s' method.", c.Request.Method)
	}
}
func AssertPost(c *Context, t *testing.T) {
	if !c.IsPost() {
		t.Errorf("IsPost should be true for '%s' method.", c.Request.Method)
	}
}
func AssertPut(c *Context, t *testing.T) {
	if !c.IsPut() {
		t.Errorf("IsPut should be true for '%s' method.", c.Request.Method)
	}
}
func AssertDelete(c *Context, t *testing.T) {
	if !c.IsDelete() {
		t.Errorf("IsDelete should be true for '%s' method.", c.Request.Method)
	}
}
func AssertOptions(c *Context, t *testing.T) {
	if !c.IsOptions() {
		t.Errorf("IsOptions should be true for '%s' method.", c.Request.Method)
	}
}
func AssertHead(c *Context, t *testing.T) {
	if !c.IsHead() {
		t.Errorf("IsHead should be true for '%s' method.", c.Request.Method)
	}
}

func TestIsGet(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = GET_HTTP_METHOD

	AssertGet(context, t)
	AssertNotPost(context, t)
	AssertNotPut(context, t)
	AssertNotDelete(context, t)

}
func TestIsPost(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = POST_HTTP_METHOD

	AssertNotGet(context, t)
	AssertPost(context, t)
	AssertNotPut(context, t)
	AssertNotDelete(context, t)
	AssertNotOptions(context, t)
	AssertNotHead(context, t)

}
func TestIsPut(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = PUT_HTTP_METHOD

	AssertNotGet(context, t)
	AssertNotPost(context, t)
	AssertPut(context, t)
	AssertNotDelete(context, t)
	AssertNotOptions(context, t)
	AssertNotHead(context, t)

}
func TestIsDelete(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = DELETE_HTTP_METHOD

	AssertNotGet(context, t)
	AssertNotPost(context, t)
	AssertNotPut(context, t)
	AssertDelete(context, t)
	AssertNotOptions(context, t)
	AssertNotHead(context, t)

}
func TestIsOptions(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = OPTIONS_HTTP_METHOD

	AssertNotGet(context, t)
	AssertNotPost(context, t)
	AssertNotPut(context, t)
	AssertNotDelete(context, t)
	AssertOptions(context, t)
	AssertNotHead(context, t)

}
func TestIsHead(t *testing.T) {

	context := MakeTestContext()

	// set the request method
	context.Request.Method = HEAD_HTTP_METHOD

	AssertNotGet(context, t)
	AssertNotPost(context, t)
	AssertNotPut(context, t)
	AssertNotDelete(context, t)
	AssertNotOptions(context, t)
	AssertHead(context, t)

}

/*
	Form data and parameter readers
*/
func TestRequestContextValue(t *testing.T) {

	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?context=123")
	if context.GetRequestContext() != "123" {
		t.Errorf("GetRequestContext() should return the correct request context")
	}

	context = MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	if context.GetRequestContext() != "" {
		t.Errorf("GetRequestContext() should return an empty string if no context is present")
	}

}

func TestCallbackValue(t *testing.T) {

	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?callback=myFunc")
	if context.GetCallback() != "myFunc" {
		t.Errorf("GetCallback() should return the correct request callback value")
	}

	context = MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	if context.GetCallback() != "" {
		t.Errorf("GetCallback() should return an empty string if no callback is present")
	}

}

/*
	API Responders
*/
func TestRespondContentType(t *testing.T) {

	data := "This is the data"
	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	context.ResponseWriter = response

	context.RespondWithData(data)

	context.assertContentType(t, "application/json")

}

func TestRespondWithData(t *testing.T) {

	data := "This is the data"
	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	context.ResponseWriter = response

	context.RespondWithData(data)

	if response.WrittenHeaderInt != 200 {
		t.Errorf("RespondWithData should have written 200 status (not %d)", response.WrittenHeaderInt)
	}

	assertEqual(t, response.Output, "{\"C\":\"\",\"S\":200,\"D\":\"This is the data\",\"E\":null}", "RespondWithData wrong")

}

func TestRespondWithError(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	context.ResponseWriter = response

	context.RespondWithError(http.StatusNotImplemented)

	if response.WrittenHeaderInt != http.StatusNotImplemented {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusNotImplemented, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"\",\"S\":501,\"D\":null,\"E\":[\"Not Implemented\"]}", "for TestRespondWithError")

}

func TestRespondWithErrorMessage(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	context.ResponseWriter = response

	context.RespondWithErrorMessage("Something went wrong", http.StatusNotImplemented)

	if response.WrittenHeaderInt != http.StatusNotImplemented {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusNotImplemented, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"\",\"S\":501,\"D\":null,\"E\":[\"Something went wrong\"]}", "for TestRespondWithError")

}

func TestRespondWithStatus(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json")
	context.ResponseWriter = response

	context.RespondWithStatus(http.StatusNotImplemented)

	if response.WrittenHeaderInt != http.StatusNotImplemented {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusNotImplemented, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"\",\"S\":501,\"D\":null,\"E\":null}", "for TestRespondWithStatus")

}

func TestRespondWithNotFound(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?" + REQUEST_CONTEXT_PARAMETER + "=123")
	context.ResponseWriter = response

	context.RespondWithNotFound()

	if response.WrittenHeaderInt != http.StatusNotFound {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusNotFound, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"123\",\"S\":404,\"D\":null,\"E\":[\"Not Found\"]}", "for TestRespondWithStatus")

}

func TestRespondWithOK(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?" + REQUEST_CONTEXT_PARAMETER + "=123")
	context.ResponseWriter = response

	context.RespondWithOK()

	if response.WrittenHeaderInt != http.StatusOK {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusOK, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"123\",\"S\":200,\"D\":null,\"E\":null}", "for TestRespondWithStatus")

}

func TestRespondWithNotImplemented(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?" + REQUEST_CONTEXT_PARAMETER + "=123")
	context.ResponseWriter = response

	context.RespondWithNotImplemented()

	if response.WrittenHeaderInt != http.StatusNotImplemented {
		t.Errorf("RespondWithData should have written %d status (not %d)", http.StatusNotImplemented, response.WrittenHeaderInt)
	}
	assertEqual(t, response.Output, "{\"C\":\"123\",\"S\":501,\"D\":null,\"E\":[\"Not Implemented\"]}", "for TestRespondWithStatus")

}

func TestRespondWithObject_ContextIsPassedThrough(t *testing.T) {

	requestContext := "this-is-the-context"
	gowebContext := MakeTestContextWithUrl(testDomain + "/people.json?" + REQUEST_CONTEXT_PARAMETER + "=" + requestContext)

	response := new(TestResponseWriter)
	gowebContext.ResponseWriter = response

	gowebContext.RespondWithData("data")

	assertEqual(t, response.Output, "{\"C\":\""+requestContext+"\",\"S\":200,\"D\":\"data\",\"E\":null}", "for TestRespondWithObject_ContextIsPassedThrough")

}

func TestRespondWithData_WithCallbackFunction(t *testing.T) {

	data := "This is the data"
	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?callback=doSomething")
	context.ResponseWriter = response

	context.RespondWithData(data)

	if response.WrittenHeaderInt != 200 {
		t.Errorf("RespondWithData should have written 200 status (not %d)", response.WrittenHeaderInt)
	}

	assertEqual(t, response.Output, "doSomething({\"C\":\"\",\"S\":200,\"D\":\"This is the data\",\"E\":null})", "TestRespondWithData_WithCallbackFunction wrong")

	context.assertContentType(t, JSONP_CONTENT_TYPE)

}
func TestRespondWithData_WithCallbackFunctionAndContext(t *testing.T) {

	data := "This is the data"
	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?callback=doSomething&context=123")
	context.ResponseWriter = response

	context.RespondWithData(data)

	if response.WrittenHeaderInt != 200 {
		t.Errorf("RespondWithData should have written 200 status (not %d)", response.WrittenHeaderInt)
	}

	assertEqual(t, response.Output, "doSomething({\"C\":\"123\",\"S\":200,\"D\":\"This is the data\",\"E\":null}, \"123\")", "TestRespondWithData_WithCallbackFunction wrong")

}

func TestAlways200(t *testing.T) {

	response := new(TestResponseWriter)
	context := MakeTestContextWithUrl(testDomain + "/people/123/groups/456.json?callback=doSomething&always200&context=123")
	context.ResponseWriter = response

	// send an internal server error
	context.RespondWithStatus(http.StatusInternalServerError)

	if response.WrittenHeaderInt != 200 {
		t.Errorf("RespondWithData should have written 200 status (not %d) with ?always200", response.WrittenHeaderInt)
	}

	assertEqual(t, response.Output, "doSomething({\"C\":\"123\",\"S\":500,\"D\":null,\"E\":null}, \"123\")", "TestAlways200 wrong")

}
