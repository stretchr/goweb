package webcontext_test

import (
	"fmt"
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/webcontext"
	http_test "github.com/stretchr/testify/http"
	"net/http"
	"strings"
)

// TestRequest is the most recent test *http.Request object that
// was created and used in a WebContext.
var TestRequest *http.Request

// TestResponseWriter is the most recent testify/http/TestResponseWriter object
// that was used in the WebContext.  It will contain the last response.
var TestResponseWriter *http_test.TestResponseWriter

// testCodecService is the most recent CodecService that was used in the
// WebContext.
var testCodecService codecsservices.CodecService

// MakeTestContext makes a *webcontext.WebContext (that can be used as
// a context.Context) that can be used for esting.
func MakeTestContext() *webcontext.WebContext {
	return MakeTestContextWithPath("/")
}

// MakeTestContextWithPath makes a *webcontext.WebContext for testing with
// the specified path.
//
//     webcontext_test.MakeTestContextWithPath("http://mysite.com/path/{var}?queryparam1=1")
func MakeTestContextWithPath(path string) *webcontext.WebContext {
	return MakeTestContextWithDetails(path, "GET")
}

// MakeTestContextWithDetails makes a *webcontext.WebContext with the specified path
// and HTTP Method.
func MakeTestContextWithDetails(path, method string) *webcontext.WebContext {
	return MakeTestContextWithFullDetails(fmt.Sprintf("http://stretchr.org/%s", path), method, "")
}

// MakeTestContextWithFullDetails makes a *webcontext.WebContext with the specified
// path, HTTP Method and body string.
func MakeTestContextWithFullDetails(path, method, body string) *webcontext.WebContext {
	testCodecService = codecsservices.NewWebCodecService()
	TestResponseWriter = new(http_test.TestResponseWriter)

	if len(body) == 0 {
		TestRequest, _ = http.NewRequest(method, path, nil)
	} else {
		TestRequest, _ = http.NewRequest(method, path, strings.NewReader(body))
	}

	return webcontext.NewWebContext(TestResponseWriter, TestRequest, testCodecService)

}
