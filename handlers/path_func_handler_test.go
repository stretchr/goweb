package handlers

import (
	"github.com/stretchrcom/goweb/context"
	context_test "github.com/stretchrcom/goweb/context/test"
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPathMatchHandler(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("collection/{id}/name")
	var called bool = false
	h := PathFuncHandler{pathPattern, HandlerFunc(func(c *context.Context) error {
		called = true
		return nil
	})}

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := h.WillHandle(ctx1)
	assert.True(t, will)
	h.Handle(ctx1)
	assert.True(t, called, "Method should be called")
	assert.Equal(t, "123", ctx1.Data().Get(contextURLParametersDataKey).(objects.Map).Get("id"))

	ctx2 := context_test.MakeTestContextWithPath("/collection")
	will, _ = h.WillHandle(ctx2)
	assert.False(t, will)
	assert.Nil(t, ctx2.Data().Get(contextURLParametersDataKey))

	assert.Nil(t, h.Handle(ctx2))
	assert.True(t, called, "Handler func should get called")

}
