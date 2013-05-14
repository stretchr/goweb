package webcontext_test

import (
	"fmt"
	"github.com/stretchrcom/goweb/webcontext"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
)

var TestRequest *http.Request
var TestResponseWriter *http_test.TestResponseWriter

func MakeTestContext() *webcontext.WebContext {
	return MakeTestContextWithPath("/")
}

func MakeTestContextWithPath(path string) *webcontext.WebContext {
	return MakeTestContextWithDetails(path, "GET")
}

func MakeTestContextWithDetails(path, method string) *webcontext.WebContext {
	TestResponseWriter = new(http_test.TestResponseWriter)
	TestRequest, _ = http.NewRequest(method, fmt.Sprintf("http://stretchr.org/%s", path), nil)

	return webcontext.NewWebContext(TestResponseWriter, TestRequest)
}
