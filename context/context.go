package context

import (
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"net/http"
)

/*
  Context represents an object that represents a single HTTP request.
*/
type Context interface {

	/*
		HTTP
		----------------------------------------
	*/

	// HttpResponseWriter gets the underlying http.ResponseWriter that will be used
	// to respond to this request.
	HttpResponseWriter() http.ResponseWriter

	// HttpRequest gets the underlying http.Request that this Context represents.
	HttpRequest() *http.Request

	/*
		Data
		----------------------------------------
	*/

	// Path gets the paths.Path of the request.
	Path() *paths.Path

	// Data gets a map of data about the context.
	Data() objects.Map

	// PathParams gets any parameters that were pulled from the URL path.
	PathParams() objects.Map
}
