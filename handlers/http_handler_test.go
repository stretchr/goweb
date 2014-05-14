package handlers

import (
	"errors"
	"net/http"
	"testing"

	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
	handlers_test "github.com/stretchr/goweb/handlers/test"
	"github.com/stretchr/testify/assert"
	http_test "github.com/stretchr/testify/http"
	"github.com/stretchr/testify/mock"
)

func TestNewHttpHandler(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	assert.Equal(t, 3, len(h.Handlers))
	assert.Equal(t, codecService, h.CodecService())

}

func TestAppendHandler(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler1 := new(handlers_test.TestHandler)
	h := NewHttpHandler(codecService)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendHandler(handler1)
	h.Handlers.Handle(nil)
	assert.Equal(t, 1, len(h.HandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestDataGetsCopiedToEachContext(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler1 := new(handlers_test.TestHandler)
	h := NewHttpHandler(codecService)

	h.Data.Set("name", "Mat")

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendHandler(handler1)
	req, _ := http.NewRequest("GET", "something", nil)
	h.ServeHTTP(nil, req)

	assert.Equal(t, 1, len(h.HandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)
	ctx := handler1.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, ctx.Data().Get("name").Str(), "Mat")

}

func TestAppendPreHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendPreHandler(handler1)
	h.Handlers.Handle(nil)
	assert.Equal(t, 1, len(h.PreHandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestAppendPostHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendPostHandler(handler1)
	h.Handlers.Handle(nil)
	assert.Equal(t, 1, len(h.PostHandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestPrependPreHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	handler1.TestData().Set("id", 1)
	handler2.TestData().Set("id", 2)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)
	handler2.On("WillHandle", mock.Anything).Return(true, nil)
	handler2.On("Handle", mock.Anything).Return(false, nil)

	h.PrependPreHandler(handler1)
	h.PrependPreHandler(handler2)
	h.Handlers.Handle(nil)
	assert.Equal(t, 2, len(h.PreHandlersPipe()))

	assert.Equal(t, 2, h.PreHandlersPipe()[0].(*handlers_test.TestHandler).TestData().Get("id").Data())
	assert.Equal(t, 1, h.PreHandlersPipe()[1].(*handlers_test.TestHandler).TestData().Get("id").Data())

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestPrependPostHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	handler1.TestData().Set("id", 1)
	handler2.TestData().Set("id", 2)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)
	handler2.On("WillHandle", mock.Anything).Return(true, nil)
	handler2.On("Handle", mock.Anything).Return(false, nil)

	h.PrependPostHandler(handler1)
	h.PrependPostHandler(handler2)
	h.Handlers.Handle(nil)
	assert.Equal(t, 2, len(h.PostHandlersPipe()))

	assert.Equal(t, 2, h.PostHandlersPipe()[0].(*handlers_test.TestHandler).TestData().Get("id").Data())
	assert.Equal(t, 1, h.PostHandlersPipe()[1].(*handlers_test.TestHandler).TestData().Get("id").Data())

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestServeHTTP(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://stretchr.org/goweb", nil)
	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	// setup some test handlers
	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	handler.Handlers = append(handler.Handlers, handler1)
	handler.Handlers = append(handler.Handlers, handler2)
	handler.Handlers = append(handler.Handlers, handler3)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)
	handler2.On("WillHandle", mock.Anything).Return(true, nil)
	handler2.On("Handle", mock.Anything).Return(false, nil)
	handler3.On("WillHandle", mock.Anything).Return(true, nil)
	handler3.On("Handle", mock.Anything).Return(false, nil)

	handler.ServeHTTP(responseWriter, testRequest)

	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

	// get the first context
	ctx := handler1.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, responseWriter, ctx.HttpResponseWriter())
	assert.Equal(t, testRequest, ctx.HttpRequest())

	// make sure it's always the same context
	ctx1 := handler1.Calls[0].Arguments[0].(context.Context)
	ctx2 := handler2.Calls[0].Arguments[0].(context.Context)
	ctx3 := handler3.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, ctx1, ctx2, "Contexts should be the same")
	assert.Equal(t, ctx2, ctx3, "Contexts should be the same")

}

func TestServeHTTPMethodOverride(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("POST", "http://stretchr.org/goweb", nil)
	testRequest.Header.Set("X-HTTP-Method-Override", "GET")
	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	// setup some test handlers
	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	handler.Handlers = append(handler.Handlers, handler1)
	handler.Handlers = append(handler.Handlers, handler2)
	handler.Handlers = append(handler.Handlers, handler3)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)
	handler2.On("WillHandle", mock.Anything).Return(true, nil)
	handler2.On("Handle", mock.Anything).Return(false, nil)
	handler3.On("WillHandle", mock.Anything).Return(true, nil)
	handler3.On("Handle", mock.Anything).Return(false, nil)

	handler.ServeHTTP(responseWriter, testRequest)

	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

	// get the first context
	ctx := handler1.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, responseWriter, ctx.HttpResponseWriter())
	assert.Equal(t, testRequest, ctx.HttpRequest())

	// make sure it's always the same context
	ctx1 := handler1.Calls[0].Arguments[0].(context.Context)
	ctx2 := handler2.Calls[0].Arguments[0].(context.Context)
	ctx3 := handler3.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, ctx1, ctx2, "Contexts should be the same")
	assert.Equal(t, ctx2, ctx3, "Contexts should be the same")

}

/*
	Errors
*/

func TestGetAndSetErrorHandler(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	errorHandler := new(handlers_test.TestHandler)

	// default one should be made
	assert.NotNil(t, handler.ErrorHandler())

	//... but if we set one explicitally
	handler.SetErrorHandler(errorHandler)

	//... it should be set!
	assert.Equal(t, errorHandler, handler.ErrorHandler())

}

func TestErrorHandlerGetsUsedOnError(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://stretchr.org/goweb", nil)

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	errorHandler := new(handlers_test.TestHandler)
	handler.SetErrorHandler(errorHandler)

	errorHandler.On("Handle", mock.Anything).Return(false, nil)

	// make a handler throw an error
	var theError error = errors.New("Test error")
	handler.Map(func(c context.Context) error {
		return theError
	})

	handler.ServeHTTP(responseWriter, testRequest)

	if mock.AssertExpectationsForObjects(t, errorHandler.Mock) {

		// get the first context
		ctx := errorHandler.Calls[0].Arguments[0].(context.Context)

		// make sure the error data field was set
		assert.Equal(t, theError.Error(), ctx.Data().Get("error").Data().(HandlerError).Error(), "the error should be set in the data with the 'error' key")

		assert.Equal(t, responseWriter, ctx.HttpResponseWriter())
		assert.Equal(t, testRequest, ctx.HttpRequest())

	}

}
