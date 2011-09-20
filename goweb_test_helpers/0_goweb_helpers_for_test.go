/*
	
	Helpful methods for writing tests for goweb code.
	by Mat Ryer
	
*/

// TODO: insert your package here
package replace_me_with_your_project_name

import (
	"testing"
	"strings"
  "goweb"
  "http"
	"os"
)

var currentRequest *http.Request
var currentResponse *TestResponseWriter
var currentContext *goweb.Context

/*
	Make a fake goweb.Context object
*/
func MakeTestContext(method, url string) *goweb.Context {

  // make a fake http request
  currentRequest, _ = http.NewRequest(method, url, nil)

	// make a test response 
	currentResponse = new(TestResponseWriter)

	currentContext = new(goweb.Context)
	currentContext.Format = goweb.JSON_FORMAT
	currentContext.Request = currentRequest
	currentContext.ResponseWriter = currentResponse
	currentContext.PathParams = make(goweb.ParameterValueMap)

	return currentContext
}

/*
	Asserts that the last response has the specified HTTP Status Code
*/
func assertResponseStatus(t *testing.T, status int) {
	if currentResponse.WrittenHeaderInt != status {
		t.Errorf("HTTP status code expected to be %d, not %d", status, currentResponse.WrittenHeaderInt)
	}
}

/*
	Asserts that the last response was a redirection to the
	specified URL
*/
func assertResponseRedirected(t *testing.T, url string) {
	assertResponseStatus(t, http.StatusFound)
	if len(currentResponse.Header()["Location"]) != 1 {
		t.Errorf("'Location' header should be there.")
	} else {
		
		location := currentResponse.Header()["Location"][0]
		expectedLocation := url
		
		if location != expectedLocation {
			t.Errorf("Location header should be '%s' not '%s'.", expectedLocation, location)
		}
		
	}
}

/*
	Asserts that the last response body was correct
*/
func assertResponse(t *testing.T, response string) {
	if currentResponse.Output != response {
		t.Errorf("Response incorrect.\n\tExpected: \"%s\"\n\tActual:   \"%s\"\t", response, currentResponse.Output)
	}
}

/*
	Asserts that a string contains another
*/
func assertContains(t *testing.T, s string, contains string) {
	if strings.Index(s, contains) == -1 {
		t.Errorf("String expected to contain: '%s' but didn't:  \"%s\"", contains, s)
	}
}

/*
	Asserts that the last response contains the specified string
*/
func assertResponseContains(t *testing.T, contains string) {
	assertContains(t, currentResponse.Output, contains)
}

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
	
	//log.Printf("TestResponseWriter: Write: %s", string(bytes))
	
	// add these bytes to the output string
	rw.Output = rw.Output + string(bytes)
	
	// return normal values
	return 0, nil
	
}
func (rw *TestResponseWriter) WriteHeader(i int) {
	//log.Printf("TestResponseWriter: WriteHeader: %d", i)
	rw.WrittenHeaderInt = i
}
