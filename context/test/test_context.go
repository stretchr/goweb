package context_test

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
)

func MakeTestContext() *context.Context {
	return MakeTestContextWithPath("/")
}

func MakeTestContextWithPath(path string) *context.Context {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", fmt.Sprintf("http://stretchr.org/%s", path), nil)

	return context.NewContext(responseWriter, testRequest)
}
