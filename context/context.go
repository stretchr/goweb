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

	// SetHttpResponseWriter sets the HttpResponseWriter that will be used to respond
	// to the request.
	//
	// This is set by Goweb, but can be overridden if you want to intercept the usual
	// writes to do something lower level with them.
	// For example, save the response in memory for testing or
	// logging purposes.
	//
	// For production, if you set your own ResponseWriter, be sure to also write the
	// response to the original ResponseWriter so that clients actually receive it.  You can
	// get the original ResponseWriter by calling the HttpResponseWriter() method on this
	// object.
	SetHttpResponseWriter(responseWriter http.ResponseWriter)

	// SetHttpRequest sets the HttpRequest that represents the original request that
	// issued the interaction.  This is set automatically by Goweb, but can be overridden for
	// advanced cases.
	SetHttpRequest(httpRequest *http.Request)

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
