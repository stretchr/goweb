package webcontext_test

import (
	"fmt"
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/webcontext"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
)

var TestRequest *http.Request
var TestResponseWriter *http_test.TestResponseWriter
var TestCodecService codecservices.CodecService

func MakeTestContext() *webcontext.WebContext {
	return MakeTestContextWithPath("/")
}

func MakeTestContextWithPath(path string) *webcontext.WebContext {
	return MakeTestContextWithDetails(path, "GET")
}

func MakeTestContextWithDetails(path, method string) *webcontext.WebContext {

	TestCodecService = new(codecservices.WebCodecService)
	TestResponseWriter = new(http_test.TestResponseWriter)
	TestRequest, _ = http.NewRequest(method, fmt.Sprintf("http://stretchr.org/%s", path), nil)

	return webcontext.NewWebContext(TestResponseWriter, TestRequest, TestCodecService)
}
