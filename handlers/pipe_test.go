package handlers

import (
	handlers_test "github.com/stretchr/goweb/handlers/test"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPipe(t *testing.T) {

	assert.Implements(t, (*Handler)(nil), new(Pipe))

}

func TestPipe_WillHandle(t *testing.T) {

	p := new(Pipe)

	handle, _ := p.WillHandle(nil)
	assert.True(t, handle, "Pipes always will handle")

}

func TestPipe_AppendHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := make(Pipe, 0)
	p = p.AppendHandler(handler1)

	if assert.Equal(t, 1, len(p)) {
		assert.Equal(t, handler1, p[0])
	}

}

func TestPipe_PrependHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := make(Pipe, 0)
	p = p.PrependHandler(handler1)
	p = p.PrependHandler(handler2)
	p = p.PrependHandler(handler3)

	if assert.Equal(t, 3, len(p)) {
		assert.Equal(t, handler3, p[0])
		assert.Equal(t, handler2, p[1])
		assert.Equal(t, handler1, p[2])
	}

}

func TestPipe_Handle(t *testing.T) {

	ctx := context_test.MakeTestContext()

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := Pipe{handler1, handler2, handler3}

	// setup expectations
	handler1.On("WillHandle", ctx).Return(true, nil)
	handler1.On("Handle", ctx).Return(false, nil)
	handler2.On("WillHandle", ctx).Return(false, nil)
	handler3.On("WillHandle", ctx).Return(true, nil)
	handler3.On("Handle", ctx).Return(false, nil)

	// call handle
	p.Handle(ctx)

	// assert expectations
	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

}

func TestPipe_Handle_Stopping(t *testing.T) {

	ctx := context_test.MakeTestContext()

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := Pipe{handler1, handler2, handler3}

	// setup expectations
	handler1.On("WillHandle", ctx).Return(true, nil)
	handler1.On("Handle", ctx).Return(true, nil)

	// call handle
	p.Handle(ctx)

	// assert expectations
	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

}
