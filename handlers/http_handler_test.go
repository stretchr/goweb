package handlers

import (
	"github.com/stretchrcom/goweb/context"
	handlers_test "github.com/stretchrcom/goweb/handlers/test"
	"github.com/stretchrcom/testify/assert"
	http_test "github.com/stretchrcom/testify/http"
	"github.com/stretchrcom/testify/mock"
	"net/http"
	"testing"
)

func TestNewHttpHandler(t *testing.T) {

	h := NewHttpHandler()

	if assert.Equal(t, 3, len(h.Handlers)) {
	}

}

func TestAppendHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	h := NewHttpHandler()

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendHandler(handler1)
	h.Handlers.Handle(nil)
	assert.Equal(t, 1, len(h.HandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestAppendPreHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	h := NewHttpHandler()

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(false, nil)

	h.AppendPreHandler(handler1)
	h.Handlers.Handle(nil)
	assert.Equal(t, 1, len(h.PreHandlersPipe()))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestAppendPostHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	h := NewHttpHandler()

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
	h := NewHttpHandler()

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

	assert.Equal(t, 2, h.PreHandlersPipe()[0].(*handlers_test.TestHandler).TestData().Get("id"))
	assert.Equal(t, 1, h.PreHandlersPipe()[1].(*handlers_test.TestHandler).TestData().Get("id"))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestPrependPostHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	h := NewHttpHandler()

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

	assert.Equal(t, 2, h.PostHandlersPipe()[0].(*handlers_test.TestHandler).TestData().Get("id"))
	assert.Equal(t, 1, h.PostHandlersPipe()[1].(*handlers_test.TestHandler).TestData().Get("id"))

	mock.AssertExpectationsForObjects(t, handler1.Mock)

}

func TestServeHTTP(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://stretchr.org/goweb", nil)
	handler := NewHttpHandler()

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

}
