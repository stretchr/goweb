package webcontext

import (
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	http_test "github.com/stretchr/testify/http"
	"net/http"
	"strings"
	"testing"
)

func TestNewContext(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)
	codecService := codecsservices.NewWebCodecService()

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
	codecService := codecsservices.NewWebCodecService()

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

func TestFileExtension(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	codecService := codecsservices.NewWebCodecService()

	testRequest, _ := http.NewRequest("get", "http://goweb.org/people/123.json", nil)
	c := NewWebContext(responseWriter, testRequest, codecService)
	assert.Equal(t, ".json", c.FileExtension())

	testRequest, _ = http.NewRequest("get", "http://goweb.org/people/123.bson", nil)
	c = NewWebContext(responseWriter, testRequest, codecService)
	assert.Equal(t, ".bson", c.FileExtension())

	testRequest, _ = http.NewRequest("get", "http://goweb.org/people/123.xml", nil)
	c = NewWebContext(responseWriter, testRequest, codecService)
	assert.Equal(t, ".xml", c.FileExtension())

	testRequest, _ = http.NewRequest("get", "http://goweb.org/people.with.dots/123.xml", nil)
	c = NewWebContext(responseWriter, testRequest, codecService)
	assert.Equal(t, ".xml", c.FileExtension())

	testRequest, _ = http.NewRequest("get", "http://goweb.org/people.with.dots/123.xml?a=b", nil)
	c = NewWebContext(responseWriter, testRequest, codecService)
	assert.Equal(t, ".xml", c.FileExtension())

}

func TestMethodString(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("get", "http://goweb.org/people/123", nil)

	codecService := codecsservices.NewWebCodecService()

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

	responseWriter = new(http_test.TestResponseWriter)
	testRequest, _ = http.NewRequest("GET", "http://goweb.org/people/123?method=PATCH", nil)

	c = NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "PATCH", c.MethodString())

}

func TestData(t *testing.T) {

	c := new(WebContext)

	c.data = nil

	assert.NotNil(t, c.Data())
	assert.NotNil(t, c.data)

}

func TestCodecOptions(t *testing.T) {

	c := new(WebContext)

	c.codecOptions = nil

	assert.NotNil(t, c.CodecOptions())
	assert.NotNil(t, c.codecOptions)
}

func TestPathParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", nil)

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)
	c.Data().Set(context.DataKeyPathParameters, objx.Map{"animal": "monkey"})

	assert.Equal(t, "monkey", c.PathParams().Get("animal").Data())
	assert.Equal(t, "monkey", c.PathValue("animal"))
	assert.Equal(t, "", c.PathValue("doesn't exist"))

}

func TestRequestData(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", strings.NewReader("{\"something\":true}"))

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	bod, _ := c.RequestBody()
	assert.Equal(t, "{\"something\":true}", string(bod))
	dat, datErr := c.RequestData()

	if assert.NoError(t, datErr) {
		assert.Equal(t, true, dat.(map[string]interface{})["something"])
	}

	responseWriter = new(http_test.TestResponseWriter)
	testRequest, _ = http.NewRequest("GET", "http://goweb.org/people/123?body={\"something\":true}", nil)

	codecService = codecsservices.NewWebCodecService()

	c = NewWebContext(responseWriter, testRequest, codecService)

	bod, _ = c.RequestBody()
	assert.Equal(t, "{\"something\":true}", string(bod))
	dat, datErr = c.RequestData()

	if assert.NoError(t, datErr) {
		assert.Equal(t, true, dat.(map[string]interface{})["something"])
	}

}

func TestRequestData_ArrayOfData(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123", strings.NewReader("[{\"something\":true},{\"something\":false}]"))

	codecService := codecsservices.NewWebCodecService()

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

/*
	Post parameters
*/
func TestPostParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	params := c.PostParams()

	if assert.NotNil(t, params) {

		assert.Equal(t, "Mat", params.Get("name").StrSlice()[0])
		assert.Equal(t, "Laurie", params.Get("name").StrSlice()[1])
		assert.Equal(t, "30", params.Get("age").StrSlice()[0])
		assert.Equal(t, "true", params.Get("something").StrSlice()[0])
		assert.Nil(t, params.Get("query").Data())

	}

}

func TestPostValues(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	names := c.PostValues("name")

	if assert.Equal(t, 2, len(names)) {
		assert.Equal(t, "Mat", names[0])
		assert.Equal(t, "Laurie", names[1])
	}

	assert.Nil(t, c.PostValues("no-such-value"))
	assert.Nil(t, c.PostValues("query"))

}

func TestPostValue(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "Mat", c.PostValue("name"), "QueryValue should get first value")
	assert.Equal(t, "30", c.PostValue("age"))
	assert.Equal(t, "", c.PostValue("no-such-value"))
	assert.Equal(t, "", c.PostValue("query"))

}

/*
	Form parameters
*/

func TestFormParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	params := c.FormParams()

	if assert.NotNil(t, params) {

		assert.Equal(t, "Mat", params.Get("name").StrSlice()[0])
		assert.Equal(t, "Laurie", params.Get("name").StrSlice()[1])
		assert.Equal(t, "30", params.Get("age").StrSlice()[0])
		assert.Equal(t, "true", params.Get("something").StrSlice()[0])
		assert.Equal(t, "yes", params.Get("query").StrSlice()[0])

	}

}

func TestFormValues(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	names := c.FormValues("name")

	if assert.Equal(t, 2, len(names)) {
		assert.Equal(t, "Mat", names[0])
		assert.Equal(t, "Laurie", names[1])
	}

	assert.Nil(t, c.FormValues("no-such-value"))
	assert.Equal(t, "yes", c.FormValues("query")[0])

}

func TestFormValue(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://goweb.org/people/123?query=yes", strings.NewReader("name=Mat&name=Laurie&age=30&something=true"))
	testRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "Mat", c.FormValue("name"), "QueryValue should get first value")
	assert.Equal(t, "30", c.FormValue("age"))
	assert.Equal(t, "", c.FormValue("no-such-value"))
	assert.Equal(t, "yes", c.FormValue("query"))

}

/*
	Query parameters
*/

func TestQueryParams(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123?name=Mat&name=Laurie&age=30&something=true", strings.NewReader("[{\"something\":true},{\"something\":false}]"))

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	params := c.QueryParams()

	if assert.NotNil(t, params) {

		assert.Equal(t, "Mat", params.Get("name").StrSlice()[0])
		assert.Equal(t, "Laurie", params.Get("name").StrSlice()[1])
		assert.Equal(t, "30", params.Get("age").StrSlice()[0])
		assert.Equal(t, "true", params.Get("something").StrSlice()[0])

	}

}

func TestQueryValues(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123?name=Mat&name=Laurie&age=30&something=true", strings.NewReader("[{\"something\":true},{\"something\":false}]"))

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	names := c.QueryValues("name")

	if assert.Equal(t, 2, len(names)) {
		assert.Equal(t, "Mat", names[0])
		assert.Equal(t, "Laurie", names[1])
	}

	assert.Nil(t, c.QueryValues("no-such-value"))

}

func TestQueryValue(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://goweb.org/people/123?name=Mat&name=Laurie&age=30&something=true", strings.NewReader("[{\"something\":true},{\"something\":false}]"))

	codecService := codecsservices.NewWebCodecService()

	c := NewWebContext(responseWriter, testRequest, codecService)

	assert.Equal(t, "Mat", c.QueryValue("name"), "QueryValue should get first value")
	assert.Equal(t, "30", c.QueryValue("age"))
	assert.Equal(t, "", c.QueryValue("no-such-value"))

}
