package webcontext

import (
	codecservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/paths"
	"github.com/stretchr/stew/objects"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// WebContext is a real context.Context that represents a single request.
type WebContext struct {
	path               *paths.Path
	data               objects.Map
	httpRequest        *http.Request
	httpResponseWriter http.ResponseWriter
	requestBody        []byte
	codecService       codecservices.CodecService
	queryParams        objects.Map
	formParams         objects.Map
	postParams         objects.Map
}

// NewWebContext creates a new WebContext with the given request and response objects.
func NewWebContext(responseWriter http.ResponseWriter, request *http.Request, codecService codecservices.CodecService) *WebContext {

	c := new(WebContext)

	c.httpRequest = request
	c.httpResponseWriter = responseWriter
	c.codecService = codecService

	c.path = paths.NewPath(request.URL.Path)

	return c

}

// CodecService gets the codecservices.CodecService that this Context will use to marshal
// and unmarshal data to and from objects.
func (c *WebContext) CodecService() codecservices.CodecService {
	return c.codecService
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

// FileExtension gets the extension of the file from the HttpRequest().
func (c *WebContext) FileExtension() string {
	return strings.ToLower(path.Ext(c.HttpRequest().URL.RequestURI()))
}

// RequestData gets the data out of the body of the request as a usable object.
func (c *WebContext) RequestData() (interface{}, error) {

	// get the bytes
	bodyBytes, bodyErr := c.RequestBody()

	if bodyErr != nil {
		return nil, bodyErr
	}

	// get the right codec for the job
	codec, codecErr := c.CodecService().GetCodec(c.HttpRequest().Header.Get("Content-Type"))

	if codecErr != nil {
		return nil, codecErr
	}

	// create the object
	var obj interface{}
	unmarhsalErr := c.CodecService().UnmarshalWithCodec(codec, bodyBytes, &obj)

	return obj, unmarhsalErr
}

// RequestDataArray gets the RequestData as an []interface{} for ease.
func (c *WebContext) RequestDataArray() ([]interface{}, error) {

	obj, err := c.RequestData()
	if err != nil {
		return nil, err
	}

	return obj.([]interface{}), nil

}

// RequestBody gets the byte data out of the body of the request.
func (c *WebContext) RequestBody() ([]byte, error) {

	if len(c.requestBody) > 0 {
		return c.requestBody, nil
	}

	body, bodyErr := ioutil.ReadAll(c.HttpRequest().Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	c.requestBody = body

	return c.requestBody, nil
}

// MethodString gets the HTTP method of this request as an uppercase string.
func (c *WebContext) MethodString() string {
	return strings.ToUpper(c.HttpRequest().Method)
}

// HttpRequest gets the underlying http.Request that this Context represents.
func (c *WebContext) HttpRequest() *http.Request {
	return c.httpRequest
}

// HttpResponseWriter gets the underlying http.ResponseWriter that will be used
// to respond to this request.
func (c *WebContext) HttpResponseWriter() http.ResponseWriter {
	return c.httpResponseWriter
}

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
func (c *WebContext) SetHttpResponseWriter(responseWriter http.ResponseWriter) {
	c.httpResponseWriter = responseWriter
}

// SetHttpRequest sets the HttpRequest that represents the original request that
// issued the interaction.  This is set automatically by Goweb, but can be overridden for
// advanced cases.
func (c *WebContext) SetHttpRequest(httpRequest *http.Request) {
	c.httpRequest = httpRequest
}

// PathParams gets any parameters that were pulled from the URL path.	//
// Goweb gives you access to different types of parameters:
//
//    QueryParams - Parameters only from the URL query string
//    PostParams  - Parameters only from the body
//    FormParams  - Parameters from both the body AND the URL query string
//    PathParams  - Parameters from the path itself (i.e. /people/123)
func (c *WebContext) PathParams() objects.Map {
	return c.data.GetMap(context.DataKeyPathParameters)
}

// DEPRECATED: Use PathValue instead.
//
// PathParam gets the parameter from PathParams() with the specified keypath.
func (c *WebContext) PathParam(keypath string) string {
	panic("goweb: DEPRECATED - Use PathValue instead of PathParam.")
	return ""
}

// PathValue gets the parameter from PathParams() with the specified key.
func (c *WebContext) PathValue(keypath string) string {
	val := c.PathParams().Get(keypath)
	if valString, ok := val.(string); ok {
		return valString
	}
	return ""
}

// urlValuesToObjectsMap turns a url.Values into an objects.Map object.
//
// Will always return a real objects.Map, even if there are no values.
func (c *WebContext) urlValuesToObjectsMap(values url.Values) objects.Map {
	m := make(objects.Map)
	for k, vs := range values {
		m.Set(k, vs)
	}
	return m
}

// FormParams gets the parameters that were posted in the request body and were present
// in the URL query.
//
// Goweb gives you access to different types of parameters:
//
//    QueryParams - Parameters only from the URL query string
//    PostParams  - Parameters only from the body
//    FormParams  - Parameters from both the body AND the URL query string
//    PathParams  - Parameters from the path itself (i.e. /people/123)
func (c *WebContext) FormParams() objects.Map {

	if c.formParams == nil {

		req := c.HttpRequest()

		if req.Form == nil {
			req.ParseForm()
		}

		c.formParams = c.urlValuesToObjectsMap(req.Form)

	}

	return c.formParams
}

// FormValues gets an array of the values for the specified keypath from the
// form body in the request and the URL query.
func (c *WebContext) FormValues(keypath string) []string {

	values := c.FormParams().Get(keypath)

	if values == nil {
		return nil
	}

	return values.([]string)
}

// FormValue gets a single value for the specified keypath from the form body and
// URL query.  If there are multiple values the first value is returned.
func (c *WebContext) FormValue(keypath string) string {

	values := c.FormValues(keypath)

	if values == nil || len(values) == 0 {
		return ""
	}

	return values[0]
}

// QueryParams gets the query parameters that are present after the ? in the URL.
//
// Goweb gives you access to different types of parameters:
//
//    QueryParams - Parameters only from the URL query string
//    PostParams  - Parameters only from the body
//    FormParams  - Parameters from both the body AND the URL query string
//    PathParams  - Parameters from the path itself (i.e. /people/123)
func (c *WebContext) QueryParams() objects.Map {

	if c.queryParams == nil {
		c.queryParams = c.urlValuesToObjectsMap(c.HttpRequest().URL.Query())
	}

	return c.queryParams
}

// QueryValues gets an array of the values for the specified key from the QueryParams.
//
// Returns []string because in URLs it's possible to have multiple values for the same key,
// for example; ?name=Mat&name=Laurie&name=Tyler.
func (c *WebContext) QueryValues(keypath string) []string {

	values := c.QueryParams().Get(keypath)

	if values == nil {
		return nil
	}

	return values.([]string)
}

// QueryValue gets a single value for the specified key from the QueryParams.  If there
// are multiple values (i.e. `?name=Mat&name=Laurie`), the first value is returned.
func (c *WebContext) QueryValue(keypath string) string {

	values := c.QueryValues(keypath)

	if values == nil || len(values) == 0 {
		return ""
	}

	return values[0]
}

// PostParams gets the parameters that were posted in the request body.
//
// Goweb gives you access to different types of parameters:
//
//    QueryParams - Parameters only from the URL query string
//    PostParams  - Parameters only from the body
//    FormParams  - Parameters from both the body AND the URL query string
//    PathParams  - Parameters from the path itself (i.e. /people/123)
func (c *WebContext) PostParams() objects.Map {

	if c.postParams == nil {

		req := c.HttpRequest()

		if req.Form == nil {
			req.ParseForm()
		}

		c.postParams = c.urlValuesToObjectsMap(req.PostForm)

	}

	return c.postParams

}

// FormValues gets an array of the values for the specified keypath from the
// form body in the request.
func (c *WebContext) PostValues(keypath string) []string {

	values := c.PostParams().Get(keypath)

	if values == nil {
		return nil
	}

	return values.([]string)

}

// PostValue gets a single value for the specified keypath from the form body.
// If there are multiple values the first value is returned.
func (c *WebContext) PostValue(keypath string) string {

	values := c.PostValues(keypath)

	if values == nil || len(values) == 0 {
		return ""
	}

	return values[0]

}
