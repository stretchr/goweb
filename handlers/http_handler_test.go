package handlers

import (
	handlers_test "github.com/stretchrcom/goweb/handlers/test"
	"github.com/stretchrcom/testify/mock"
	"net/http"
	"testing"
)

func TestServeHTTP(t *testing.T) {

	testRequest, _ := http.NewRequest("GET", "http://github.com/strecthrcom/goweb", nil)
	handler := NewHttpHandler()

	// setup some test handlers
	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	handler.Handlers = append(handler.Handlers, handler1)
	handler.Handlers = append(handler.Handlers, handler2)
	handler.Handlers = append(handler.Handlers, handler3)

	handler1.On("WillHandle", mock.Anything).Return(true, nil)
	handler1.On("Handle", mock.Anything).Return(nil)
	handler2.On("WillHandle", mock.Anything).Return(true, nil)
	handler2.On("Handle", mock.Anything).Return(nil)
	handler3.On("WillHandle", mock.Anything).Return(true, nil)
	handler3.On("Handle", mock.Anything).Return(nil)

	handler.ServeHTTP(nil, testRequest)

	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

}
