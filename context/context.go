package context

import (
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/paths"
	"github.com/stretchr/objx"
	"net/http"
)

// Context represents an object that represents a single HTTP request.
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

	// FileExtension gets the extension of the file from the HttpRequest().
	FileExtension() string

	/*
		Data
		----------------------------------------
	*/

	// Path gets the paths.Path of the request.
	Path() *paths.Path

	// Data gets a map of data about the context.
	Data() objx.Map

	// CodecOptions gets a map of options that are passed to codecs
	// upon responding.  If you need to pass additional options to a
	// codec, you can add data to the value returned by this method.
	// Responders may add data to the value returned by this method
	// before passing these options to the chosen codec, but they
	// should never overwrite or remove options.
	CodecOptions() objx.Map

	// CodecService gets the codecsservices.CodecService that this Context will use to marshal
	// and unmarshal data to and from objects.
	CodecService() codecsservices.CodecService

	// RequestData gets the data out of the body of the request as a usable object.
	RequestData() (interface{}, error)

	// RequestDataArray gets the RequestData as an []interface{} for ease.
	RequestDataArray() ([]interface{}, error)

	// RequestBody gets the byte data out of the body of the request.
	RequestBody() ([]byte, error)

	// PathParams gets the parameters that were pulled from the URL path.
	//
	// Goweb gives you access to different types of parameters:
	//
	//    QueryParams - Parameters only from the URL query string
	//    PostParams  - Parameters only from the body
	//    FormParams  - Parameters from both the body AND the URL query string
	//    PathParams  - Parameters from the path itself (i.e. /people/123)
	PathParams() objx.Map

	// DEPRECATED: Use PathValue instead.
	//
	// PathParam gets the parameter from PathParams() with the specified keypath.
	PathParam(keypath string) string

	// PathValue gets the parameter from PathParams() with the specified keypath.
	PathValue(keypath string) string

	// AllQueryParams gets the parameters that were present after the ? in the URL.
	//
	// Goweb gives you access to different types of parameters:
	//
	//    QueryParams - Parameters only from the URL query string
	//    PostParams  - Parameters only from the body
	//    FormParams  - Parameters from both the body AND the URL query string
	//    PathParams  - Parameters from the path itself (i.e. /people/123)
	QueryParams() objx.Map

	// QueryValues gets an array of the values for the specified keypath from the QueryParams.
	QueryValues(keypath string) []string

	// QueryValue gets a single value for the specified keypath from the QueryParams.  If there
	// are multiple values (i.e. `?name=Mat&name=Laurie`), the first value is returned.
	QueryValue(keypath string) string

	// PostParams gets the parameters that were posted in the request body.
	//
	// Goweb gives you access to different types of parameters:
	//
	//    QueryParams - Parameters only from the URL query string
	//    PostParams  - Parameters only from the body
	//    FormParams  - Parameters from both the body AND the URL query string
	//    PathParams  - Parameters from the path itself (i.e. /people/123)
	PostParams() objx.Map

	// FormValues gets an array of the values for the specified keypath from the
	// form body in the request.
	PostValues(keypath string) []string

	// PostValue gets a single value for the specified keypath from the form body.
	// If there are multiple values the first value is returned.
	PostValue(keypath string) string

	// FormParams gets the parameters that were posted in the request body and were present
	// in the URL query.
	//
	// Goweb gives you access to different types of parameters:
	//
	//    QueryParams - Parameters only from the URL query string
	//    PostParams  - Parameters only from the body
	//    FormParams  - Parameters from both the body AND the URL query string
	//    PathParams  - Parameters from the path itself (i.e. /people/123)
	FormParams() objx.Map

	// FormValues gets an array of the values for the specified keypath from the
	// form body in the request and the URL query.
	FormValues(keypath string) []string

	// FormValue gets a single value for the specified keypath from the form body and
	// URL query.  If there are multiple values the first value is returned.
	FormValue(keypath string) string
}
