package webcontext

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
  WebContext is a real context.Context that represents a single request.
*/
type WebContext struct {
	path           *paths.Path
	data           objects.Map
	request        *http.Request
	responseWriter http.ResponseWriter
	requestBody    []byte
	codecService   codecservices.CodecService
}

// NewWebContext creates a new WebContext with the given request and response objects.
func NewWebContext(responseWriter http.ResponseWriter, request *http.Request, codecService codecservices.CodecService) *WebContext {

	c := new(WebContext)

	c.request = request
	c.responseWriter = responseWriter
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
	unmarhsalErr := codec.Unmarshal(bodyBytes, &obj)

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
	return c.request
}

// HttpResponseWriter gets the underlying http.ResponseWriter that will be used
// to respond to this request.
func (c *WebContext) HttpResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

// PathParams gets any parameters that were pulled from the URL path.
func (c *WebContext) PathParams() objects.Map {
	return c.data.GetMap(context.DataKeyPathParameters)
}

// PathParam the parameter from PathParams() with the specified key.
func (c *WebContext) PathParam(key string) string {
	val := c.PathParams().Get(key)
	if valString, ok := val.(string); ok {
		return valString
	}
	return ""
}
