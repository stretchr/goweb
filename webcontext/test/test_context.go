package webcontext_test

import (
	"fmt"
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/webcontext"
	http_test "github.com/stretchr/testify/http"
	"net/http"
	"strings"
)

var TestRequest *http.Request
var TestResponseWriter *http_test.TestResponseWriter
var TestCodecService codecsservices.CodecService

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
	TestCodecService = codecsservices.NewWebCodecService()
	TestResponseWriter = new(http_test.TestResponseWriter)

	if len(body) == 0 {
		TestRequest, _ = http.NewRequest(method, path, nil)
	} else {
		TestRequest, _ = http.NewRequest(method, path, strings.NewReader(body))
	}

	return webcontext.NewWebContext(TestResponseWriter, TestRequest, TestCodecService)

}
