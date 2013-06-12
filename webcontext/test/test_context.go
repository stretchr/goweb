package webcontext_test

import (
	"fmt"
	codecservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/webcontext"
	http_test "github.com/stretchr/testify/http"
	"net/http"
	"strings"
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
	return MakeTestContextWithFullDetails(fmt.Sprintf("http://stretchr.org/%s", path), method, "")
}

func MakeTestContextWithFullDetails(path, method, body string) *webcontext.WebContext {
	TestCodecService = new(codecservices.WebCodecService)
	TestResponseWriter = new(http_test.TestResponseWriter)

	if len(body) == 0 {
		TestRequest, _ = http.NewRequest(method, path, nil)
	} else {
		TestRequest, _ = http.NewRequest(method, path, strings.NewReader(body))
	}

	return webcontext.NewWebContext(TestResponseWriter, TestRequest, TestCodecService)

}
