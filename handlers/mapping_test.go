package handlers

import (
	"github.com/stretchrcom/goweb/context"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	handler := NewHttpHandler()

	called := false
	handler.Map("/people/{id}", func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)

	assert.True(t, called)

}

func TestMap_CatchAllAssumption(t *testing.T) {

	handler := NewHttpHandler()

	called := false
	handler.Map(func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)
	assert.True(t, called)

	called = false
	ctx = context_test.MakeTestContextWithPath("something-else")
	handler.Handlers.Handle(ctx)
	assert.True(t, called)

}
