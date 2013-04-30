package context_test

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
)

var TestRequest *http.Request
var TestResponseWriter *http_test.TestResponseWriter

func MakeTestContext() *context.Context {
	return MakeTestContextWithPath("/")
}

func MakeTestContextWithPath(path string) *context.Context {

	TestResponseWriter = new(http_test.TestResponseWriter)
	TestRequest, _ = http.NewRequest("GET", fmt.Sprintf("http://stretchr.org/%s", path), nil)

	return context.NewContext(nil, TestResponseWriter, TestRequest)
}
