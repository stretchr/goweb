package webcontext

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
	"testing"
)

func TestNewContext(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)

	c := NewWebContext(responseWriter, testRequest)

	if assert.NotNil(t, c) {

		assert.Equal(t, "people/123", c.Path().RawPath)
		assert.Equal(t, testRequest, c.request)
		assert.Equal(t, responseWriter, c.responseWriter)

	}

}

func TestData(t *testing.T) {

	c := new(WebContext)

	c.data = nil

	assert.NotNil(t, c.Data())
	assert.NotNil(t, c.data)

}

func TestPathParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)

	c := NewWebContext(responseWriter, testRequest)
	c.Data().Set(context.DataKeyPathParameters, objects.Map{"animal": "monkey"})

	assert.Equal(t, "monkey", c.PathParams().Get("animal"))

}
