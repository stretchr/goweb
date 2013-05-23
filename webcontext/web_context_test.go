package webcontext

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	http_test "github.com/stretchrcom/testify/http"
	"net/http"
	"strings"
	"testing"
)

func TestNewContext(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)
	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)

	if assert.NotNil(t, c) {

		assert.Equal(t, "people/123", c.Path().RawPath)
		assert.Equal(t, testRequest, c.httpRequest)
		assert.Equal(t, responseWriter, c.httpResponseWriter)
		assert.Equal(t, codecService, c.codecService)
		assert.Equal(t, codecService, c.CodecService())

	}

}

func TestSetHttpResponseWriter(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)
	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)

	responseWriter2 := new(http_test.TestResponseWriter)
	responseWriter2.Header().Set("Something", "true")
	testRequest2, _ := http.NewRequest("PUT", "http://goweb.org/people/123", nil)

	c.SetHttpRequest(testRequest2)
	c.SetHttpResponseWriter(responseWriter2)

	req := c.HttpRequest()
	res := c.HttpResponseWriter()

	assert.Equal(t, &testRequest2, &req)
	assert.Equal(t, responseWriter2, res)

}

func TestMethodString(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("get", "http://goweb.org/people/123", nil)

	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "GET", c.MethodString())

	responseWriter = new(http_test.TestResponseWriter)
	testRequest, _ = http.NewRequest("put", "http://goweb.org/people/123", nil)

	c = NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "PUT", c.MethodString())

	responseWriter = new(http_test.TestResponseWriter)
	testRequest, _ = http.NewRequest("DELETE", "http://goweb.org/people/123", nil)

	c = NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "DELETE", c.MethodString())

	responseWriter = new(http_test.TestResponseWriter)
	testRequest, _ = http.NewRequest("anything", "http://goweb.org/people/123", nil)

	c = NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "ANYTHING", c.MethodString())

}

func TestData(t *testing.T) {

	c := new(WebContext)

	c.data = nil

	assert.NotNil(t, c.Data())
	assert.NotNil(t, c.data)

}

func TestPathParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)

	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)
	c.Data().Set(context.DataKeyPathParameters, objects.Map{"animal": "monkey"})

	assert.Equal(t, "monkey", c.PathParams().Get("animal"))
	assert.Equal(t, "monkey", c.PathParam("animal"))
	assert.Equal(t, "", c.PathParam("doesn't exist"))

}

func TestRequestData(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", strings.NewReader("{\"something\":true}"))

	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)

	bod, _ := c.RequestBody()
	assert.Equal(t, "{\"something\":true}", string(bod))
	dat, datErr := c.RequestData()

	if assert.NoError(t, datErr) {
		assert.Equal(t, true, dat.(map[string]interface{})["something"])
	}

}

func TestRequestData_ArrayOfData(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", strings.NewReader("[{\"something\":true},{\"something\":false}]"))

	codecService := new(codecservices.WebCodecService)

	c := NewWebContext(responseWriter, testRequest, codecService)

	bod, _ := c.RequestBody()
	assert.Equal(t, "[{\"something\":true},{\"something\":false}]", string(bod))
	dat, datErr := c.RequestData()

	if assert.NoError(t, datErr) {
		assert.NotNil(t, dat.([]interface{}))
		responseDataArray, _ := c.RequestDataArray()
		assert.Equal(t, dat.([]interface{}), responseDataArray)
	}

}
