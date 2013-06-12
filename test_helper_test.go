package goweb

import (
	"github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/handlers"
	"github.com/stretchr/testify/assert"
	testifyhttp "github.com/stretchr/testify/http"
	"testing"
)

func TestTestFunc(t *testing.T) {

	// keep track of things that occur
	var actualT *testing.T
	var called bool = false

	// setup test objects
	testingObj := new(testing.T)

  // map something to a handler
  testCodecService := new(services.WebCodecService)
  handler := handlers.NewHttpHandler(testCodecService)

  // map the target method
  handler.Map("GET", "people/{id}", func())

	// call the target method
	TestOn(testingObj, handler, "GET people/123", func(passedT *testing.T, response *testifyhttp.TestResponseWriter) {

		called = true

		// save the passed in T object
		actualT = passedT

	})

	if assert.True(t, called, "") {

		// make assertions about what happened
		assert.Equal(t, actualT, testingObj, "Passed in T wasn't right")

	}

}
