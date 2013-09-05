// +build !appengine

package goweb

import (
	"github.com/stretchr/goweb/handlers"
	"github.com/stretchr/testify/assert"
	testifyhttp "github.com/stretchr/testify/http"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	// TestHttpRequest represents the current http.Request object
	// that is being tested.
	TestHttpRequest *http.Request

	// TestResponseWriter is the ResponseWriter that goweb.Test will
	// write to, in order to allow for testing.
	TestResponseWriter *testifyhttp.TestResponseWriter
)

// RequestBuilderFunc is a function that builds a TestRequest.
type RequestBuilderFunc func() *http.Request

// goweb.Test tests some functionality.  You will need to include the
// github.com/stretchr/testify/http package in order to make use of the
// test functionality.
//
// Argument signatures
//
// The Test function accepts many different signatures, following some simple
// rules:
//
//     * The first argument is always the testing.T object
//     * The second argument must be either a string, or a function that will
//       build the TestRequest (a RequestBuilderFunc).
//     * If the second argument is a string, an optional third argument would
//       be either another string or []byte array representing the body.
//     * The final argument must always be a func of type func(*testing.T, *testifyhttp.TestResponseWriter)
//
// For example:
//
//    import (
//      // import testify
//      testifyhttp "github.com/stretchr/testify/http"
//    )
//
//     // Test(t, string, func(*testing.T, *testifyhttp.TestResponseWriter))
//     // Makes a request with the specified method and path, and calls
//     // the function to make the appropriate assertions.
//     goweb.Test(t, "METHOD path", func(t *testing.T, response *testifyhttp.TestResponseWriter) {
//
//       /* assertions on the response go here */
//
//     })
//
//     // Test(t, RequestBuilderFunc, func(*testing.T, *testifyhttp.TestResponseWriter))
//     // Makes a request by calling the RequestBuilderFunc, and calls
//     // the function to make the appropriate assertions.
//     goweb.Test(t, func() *http.Request {
//
//       /* build and return the http.Request */
//
//     }, func(t *testing.T, response *testifyhttp.TestResponseWriter) {
//
//       /* assertions on the response go here */
//
//     })
//
func Test(t *testing.T, options ...interface{}) {
	TestOn(t, DefaultHttpHandler(), options...)
}

// TestOn is the same as the goweb.Test function, except it allows you
// to explicitly specify the HttpHandler on which to run the tests.
func TestOn(t *testing.T, handler *handlers.HttpHandler, options ...interface{}) {

	/*
	   Get the request builder function
	*/

	var requestBuilder RequestBuilderFunc

	switch options[0].(type) {
	case string:
		// Test(t, "GET people/123", func)

		// split out the method and path
		methodAndPath := strings.Split(options[0].(string), " ")

		// make sure we have a method and a path
		if !assert.Equal(t, 2, len(methodAndPath), "goweb: First options argument of goweb.Test, if a string, must follow the format \"METHOD path\", and cannot be \"%s\".", options[0]) {
			return
		}

		var method string = methodAndPath[0]
		var path string = methodAndPath[1]
		var body string

		// do we have a body?
		switch options[1].(type) {
		case []byte:
			body = string(options[1].([]byte))
		case string:
			body = options[1].(string)
		}

		// set the builder
		requestBuilder = func() *http.Request {

			httpRequest, httpRequestError := http.NewRequest(method, path, strings.NewReader(body))

			if httpRequestError != nil {
				t.Errorf("goweb: Could not build request: %s", httpRequestError)
			}

			return httpRequest

		}

	case RequestBuilderFunc:

		// just use their method
		requestBuilder = options[0].(RequestBuilderFunc)

	default:
		t.Errorf("goweb: First options argument of goweb.Test must be either a string, or a RequestBuilderFunc, not %v.", reflect.TypeOf(options[0]))
		return
	}

	/*
	   Get the response assertion function
	*/
	var testAssertionFunc func(*testing.T, *testifyhttp.TestResponseWriter)

	switch options[len(options)-1].(type) {
	case func(*testing.T, *testifyhttp.TestResponseWriter):

		testAssertionFunc = options[len(options)-1].(func(*testing.T, *testifyhttp.TestResponseWriter))

	default:
		t.Errorf("goweb: Last options argument of goweb.Test must be a func(*testing.T, *testifyhttp.TestResponseWriter), not %v.", reflect.TypeOf(options[len(options)-1]))
		return
	}

	/*
	   Get the TestRequest using the builder function
	*/
	TestHttpRequest = requestBuilder()

	// make sure it's not nil
	if !assert.NotNil(t, TestHttpRequest, "goweb: RequestBuilderFunc must return a *TestRequest object.") {
		return
	}

	/*
	   Build a context to use
	*/
	TestResponseWriter = new(testifyhttp.TestResponseWriter)

	/*
	   Ask Goweb to handle the context
	*/
	handler.ServeHTTP(TestResponseWriter, TestHttpRequest)

	/*
	   Over to the func(*testing.T, *testifyhttp.TestResponseWriter) to do its magic
	*/
	testAssertionFunc(t, TestResponseWriter)

}
