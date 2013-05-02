package handlers

import (
	"github.com/stretchrcom/goweb/context"
	//"github.com/stretchrcom/goweb/webcontext"
	"github.com/stretchrcom/goweb/paths"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPathMatchHandler(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("collection/{id}/name")
	var called bool = false
	h := PathMatchHandler{pathPattern, HandlerExecutionFunc(func(c context.Context) error {
		called = true
		return nil
	})}

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := h.WillHandle(ctx1)
	assert.True(t, will)
	h.Handle(ctx1)
	assert.True(t, called, "Method should be called")
	assert.Equal(t, "123", ctx1.Data().Get(context.DataKeyURLParameters).(objects.Map).Get("id"))

	ctx2 := context_test.MakeTestContextWithPath("/collection")
	will, _ = h.WillHandle(ctx2)
	assert.False(t, will)
	assert.Nil(t, ctx2.Data().Get(context.DataKeyURLParameters))

	shouldStop, handleErr := h.Handle(ctx2)
	assert.Nil(t, handleErr)
	assert.True(t, shouldStop)
	assert.True(t, called, "Handler func should get called")

}
