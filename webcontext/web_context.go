package webcontext

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"net/http"
)

/*
  WebContext is a real context.Context that represents a single request.
*/
type WebContext struct {
	path           *paths.Path
	data           objects.Map
	request        *http.Request
	responseWriter http.ResponseWriter
}

// NewWebContext creates a new WebContext with the given request and response objects.
func NewWebContext(responseWriter http.ResponseWriter, request *http.Request) *WebContext {

	c := new(WebContext)

	c.request = request
	c.responseWriter = responseWriter

	c.path = paths.NewPath(request.URL.Path)

	return c

}

// Path gets the paths.Path of the request.
func (c *WebContext) Path() *paths.Path {
	return c.path
}

// Data gets a map of data about the context.
func (c *WebContext) Data() objects.Map {
	if c.data == nil {
		c.data = make(objects.Map)
	}
	return c.data
}

// HttpRequest gets the underlying http.Request that this Context represents.
func (c *WebContext) HttpRequest() *http.Request {
	return c.request
}

// HttpResponseWriter gets the underlying http.ResponseWriter that will be used
// to respond to this request.
func (c *WebContext) HttpResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

// PathParams gets any parameters that were pulled from the URL path.
func (c *WebContext) PathParams() objects.Map {
	return c.data.GetMap(context.DataKeyURLParameters)
}
