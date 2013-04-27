package handlers

import (
	context_test "github.com/stretchrcom/goweb/context/test"
	handlers_test "github.com/stretchrcom/goweb/handlers/test"
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
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

func TestPipe_AddHandler(t *testing.T) {

	handler1 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := new(Pipe)
	assert.Equal(t, p, p.AppendHandler(handler1))

	if assert.Equal(t, 1, len(p.handlers)) {
		assert.Equal(t, handler1, p.handlers[0])
	}

}

func TestPipe_Handle(t *testing.T) {

	ctx := context_test.MakeTestContext()

	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	// add the handlers to the pipe
	p := new(Pipe)
	p.AppendHandler(handler1).AppendHandler(handler2).AppendHandler(handler3)

	// setup expectations
	handler1.On("WillHandle", ctx).Return(true, nil)
	handler1.On("Handle", ctx).Return(nil)
	handler2.On("WillHandle", ctx).Return(false, nil)
	handler3.On("WillHandle", ctx).Return(true, nil)
	handler3.On("Handle", ctx).Return(nil)

	// call handle
	p.Handle(ctx)

	// assert expectations
	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

}
