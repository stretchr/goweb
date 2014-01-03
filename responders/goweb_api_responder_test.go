package responders

import (
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"github.com/stretchr/codecs/json"
)

func TestAPI_Interface(t *testing.T) {

	assert.Implements(t, (*APIResponder)(nil), new(GowebAPIResponder))

}

func TestNewGowebAPIResponder(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	api := NewGowebAPIResponder(codecService, http)

	assert.Equal(t, http, api.httpResponder)
	assert.Equal(t, codecService, api.GetCodecService())

	assert.Equal(t, api.StandardFieldStatusKey, "s")
	assert.Equal(t, api.StandardFieldDataKey, "d")
	assert.Equal(t, api.StandardFieldErrorsKey, "e")

}

func TestRespond(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestRespondEnvelopOptions(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContextWithPath("/?envelop=false")
	data := map[string]interface{}{"name": "Mat"}

	// When AlwaysEvenlopResponse = true but ?envelop=false
	API.Respond(ctx, 200, data, nil)
	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

	// When AlwaysEvenlopResponse = false
	ctx = context_test.MakeTestContext()
	API.AlwaysEnvelopResponse = false

	API.Respond(ctx, 200, data, nil)
	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

	// When AlwaysEvenlopResponse = false but ?envelop=true
	ctx = context_test.MakeTestContextWithPath("/?envelop=true")

	API.Respond(ctx, 200, data, nil)
	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

/*
	testing codecs.Facade pattern
*/

type dataObject struct{}

func (d *dataObject) PublicData(options map[string]interface{}) (interface{}, error) {
	return map[string]interface{}{"used-public-data": true}, nil
}

func TestRespondWithPublicDataFacade(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := new(dataObject)

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"used-public-data\":true},\"s\":200}")

}

func TestRespondWithArray(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := []map[string]interface{}{
		map[string]interface{}{"name": "Mat"},
		map[string]interface{}{"name": "Tyler"},
		map[string]interface{}{"name": "Oleksandr"},
	}

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":[{\"name\":\"Mat\"},{\"name\":\"Tyler\"},{\"name\":\"Oleksandr\"}],\"s\":200}")

}

func TestRespondWithCustomFieldnames(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.StandardFieldDataKey = "data"
	API.StandardFieldStatusKey = "status"

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"data\":{\"name\":\"Mat\"},\"status\":200}")

}

func TestWriteResponseObject(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

}

// https://github.com/stretchr/goweb/issues/20
func TestWriteResponseObject_ContentNegotiation_AcceptHeader(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	ctx.HttpRequest().Header.Set("Accept", "application/x-msgpack")
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	// get the expected output
	codec, codecErr := codecService.GetCodec("application/x-msgpack")
	if assert.NoError(t, codecErr) {

		expectedOutput, marshalErr := codec.Marshal(data, nil)
		if assert.NoError(t, marshalErr) {
			assert.Equal(t, []byte(context_test.TestResponseWriter.Output), expectedOutput)
		}

	}

}

// https://github.com/stretchr/goweb/issues/20
func TestWriteResponseObject_ContentNegotiation_HasCallback(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	ctx.HttpRequest().URL, _ = url.Parse("http://stretchr.org/something?callback=doSomething")
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	// get the expected output
	codec, codecErr := codecService.GetCodec("text/javascript")
	if assert.NoError(t, codecErr) {

		expectedOutput, marshalErr := codec.Marshal(data, map[string]interface{}{"options.client.callback": "doSomething"})
		if assert.NoError(t, marshalErr) {
			assert.Equal(t, []byte(context_test.TestResponseWriter.Output), expectedOutput)
		}

	}

}

// https://github.com/stretchr/goweb/issues/20
func TestWriteResponseObject_ContentNegotiation_FileExtension(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	ctx.HttpRequest().URL, _ = url.Parse("http://stretchr.org/something.msgpack")
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	// get the expected output
	codec, codecErr := codecService.GetCodec("application/x-msgpack")
	if assert.NoError(t, codecErr) {

		expectedOutput, marshalErr := codec.Marshal(data, nil)
		if assert.NoError(t, marshalErr) {
			assert.Equal(t, []byte(context_test.TestResponseWriter.Output), expectedOutput)
		}

	}

}

type CodecOptionsTester struct {
	json.JsonCodec
}

func (c *CodecOptionsTester) Marshal(data interface{}, options map[string]interface{}) ([]byte, error) {
	encapsulated := map[string]interface{}{
		"data": data,
		"test_option": options["test_option"],
	}
	return c.JsonCodec.Marshal(encapsulated, nil)
}

func TestAPI_WriteResponseObject_CodecOptions(t *testing.T) {
	http := new(GowebHTTPResponder)
	codecService := new(codecsservices.WebCodecService)
	codecService.AddCodec(new(CodecOptionsTester))
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	test_option := "test"
	ctx.CodecOptions().Set("test_option", test_option)

	testData := "data"

	API.WriteResponseObject(ctx, 200, testData)

	assert.Equal(t, context_test.TestResponseWriter.Output, `{"data":"`+testData+`","test_option":"`+test_option+`"}`)
}

func TestAPI_StandardResponseObjectTransformer(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.SetStandardResponseObjectTransformer(func(ctx context.Context, sro interface{}) (interface{}, error) {

		return map[string]interface{}{
			"sro":       sro,
			"something": true,
		}, nil

	})

	API.RespondWithData(ctx, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"something\":true,\"sro\":{\"d\":{\"name\":\"Mat\"},\"s\":200}}")

}

func TestAPI_RespondWithData(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.RespondWithData(ctx, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestAPI_RespondWithError(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := codecsservices.NewWebCodecService()
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	errObject := "error message"

	API.RespondWithError(ctx, 500, errObject)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"e\":[\"error message\"],\"s\":500}")

}
