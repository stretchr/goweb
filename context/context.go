package context

import (
	codecservices "github.com/stretchrcom/codecs/services"
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

	// CodecService gets the codecservices.CodecService that this Context will use to marshal
	// and unmarshal data to and from objects.
	CodecService() codecservices.CodecService

	// RequestData gets the data out of the body of the request as a usable object.
	RequestData() (interface{}, error)

	// RequestDataArray gets the RequestData as an []interface{} for ease.
	RequestDataArray() ([]interface{}, error)

	// RequestBody gets the byte data out of the body of the request.
	RequestBody() ([]byte, error)

	// PathParams gets any parameters that were pulled from the URL path.
	PathParams() objects.Map

	// PathParam the parameter from PathParams() with the specified key.
	PathParam(key string) string
}
