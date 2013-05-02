package responders

import (
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"net/http"
	"testing"
)

func TestHTTP_Interface(t *testing.T) {

	assert.Implements(t, (*HTTPResponder)(nil), new(GowebHTTPResponder))

}

func TestHTTP_With(t *testing.T) {

	http := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	http.With(ctx, 200, []byte("Hello Goweb"))

	assert.Equal(t, context_test.TestResponseWriter.Output, "Hello Goweb")
	assert.Equal(t, context_test.TestResponseWriter.WrittenHeaderInt, 200)

}

func TestHTTP_WithStatus(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithStatus(ctx, 500)

	assert.Equal(t, context_test.TestResponseWriter.WrittenHeaderInt, 500)

}

func TestHTTP_WithOK(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithOK(ctx)

	assert.Equal(t, context_test.TestResponseWriter.WrittenHeaderInt, http.StatusOK)

}

func TestHTTP_WithLocation(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithLocation(ctx, "people/123")

	assert.Equal(t, context_test.TestResponseWriter.WrittenHeaderInt, http.StatusTemporaryRedirect)
	assert.Equal(t, context_test.TestResponseWriter.Header()["Location"][0], "people/123")

}
