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
		Request helpers
		----------------------------------------
	*/

	// MethodString gets the HTTP Method of the request as an uppercase string.
	MethodString() string

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

	/*
		Deprecated (these functions should panic)
		-----------------------------------------
	*/

	// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
	Respond(data interface{}, statusCode int, errors []string, context Context) error

	// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
	RespondWithData(data interface{}) error

	// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
	RespondWithError(statusCode int) error

	// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
	RespondWithErrorMessage(message string, statusCode int) error
}
