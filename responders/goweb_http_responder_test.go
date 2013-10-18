package responders

import (
	"github.com/stretchr/goweb/context"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, context_test.TestResponseWriter.StatusCode, 200)

}

func TestHTTP_WithStatus(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithStatus(ctx, 500)

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, 500)

}

func TestHTTP_WithStatusAndText(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithStatusText(ctx, 500)

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, 500)
	assert.Equal(t, context_test.TestResponseWriter.Output, http.StatusText(500))

}

// https://github.com/stretchr/goweb/issues/30
func TestHTTP_WithStatus_WithAlways200(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	var ctx context.Context

	ctx = context_test.MakeTestContextWithPath("people/123?always200=true")
	httpResponder.WithStatus(ctx, 500)
	assert.Equal(t, context_test.TestResponseWriter.StatusCode, 200)

	ctx = context_test.MakeTestContextWithPath("people/123?always200=1")
	httpResponder.WithStatus(ctx, 500)
	assert.Equal(t, context_test.TestResponseWriter.StatusCode, 200)

}

func TestHTTP_WithOK(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithOK(ctx)

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, http.StatusOK)

}

func TestHTTP_WithRedirect(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithRedirect(ctx, "people/123")

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, http.StatusFound)
	assert.Equal(t, context_test.TestResponseWriter.Header()["Location"][0], "people/123")

}

func TestHTTP_WithTemporaryRedirect(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithTemporaryRedirect(ctx, "people/123")

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, http.StatusTemporaryRedirect)
	assert.Equal(t, context_test.TestResponseWriter.Header()["Location"][0], "people/123")

}

func TestHTTP_WithPermanentRedirect(t *testing.T) {

	httpResponder := new(GowebHTTPResponder)
	ctx := context_test.MakeTestContext()

	httpResponder.WithPermanentRedirect(ctx, "people/123")

	assert.Equal(t, context_test.TestResponseWriter.StatusCode, http.StatusMovedPermanently)
	assert.Equal(t, context_test.TestResponseWriter.Header()["Location"][0], "people/123")

}
