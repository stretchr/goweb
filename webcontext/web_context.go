package webcontext

import (
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"net/http"
)

/*
  WebContext represents a single request and provides methods for responding.
*/
type WebContext struct {
	path           *paths.Path
	data           objects.Map
	request        *http.Request
	responseWriter http.ResponseWriter
}

func NewWebContext(responseWriter http.ResponseWriter, request *http.Request) *WebContext {

	c := new(WebContext)

	c.request = request
	c.responseWriter = responseWriter

	c.data = make(objects.Map)
	c.path = paths.NewPath(request.URL.Path)

	return c

}

func (c *WebContext) Path() *paths.Path {
	return c.path
}

func (c *WebContext) Data() objects.Map {
	return c.data
}

func (c *WebContext) HttpRequest() *http.Request {
	return c.request
}

func (c *WebContext) HttpResponseWriter() http.ResponseWriter {
	return c.responseWriter
}
