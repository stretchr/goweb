package context

import (
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"net/http"
)

/**
  Context represents a single request and provides methods for responding.
*/
type Context struct {
	path           *paths.Path
	data           objects.Map
	request        *http.Request
	responseWriter http.ResponseWriter
}

func NewContext(responseWriter http.ResponseWriter, request *http.Request) *Context {

	c := new(Context)

	c.request = request
	c.responseWriter = responseWriter

	c.data = make(objects.Map)
	c.path = paths.NewPath(request.URL.Path)

	return c

}

func (c *Context) Path() *paths.Path {
	return c.path
}

func (c *Context) Data() objects.Map {
	return c.data
}

func (c *Context) HttpRequest() *http.Request {
	return c.request
}

func (c *Context) HttpResponseWriter() http.ResponseWriter {
	return c.responseWriter
}
