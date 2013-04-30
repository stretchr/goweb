package api

import (
	context_test "github.com/stretchrcom/goweb/context/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*APIResponder)(nil), new(GowebAPIResponder))

}

func TestCodecService(t *testing.T) {

	// TODO: this

}

func TestRespond(t *testing.T) {

	var API *GowebAPIResponder = new(GowebAPIResponder)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestWriteResponseObject(t *testing.T) {

	var API *GowebAPIResponder = new(GowebAPIResponder)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

}
