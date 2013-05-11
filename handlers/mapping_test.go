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

func TestMap_WithMatcherFuncs(t *testing.T) {

	handler := NewHttpHandler()

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	handler.Map("/people/{id}", func(c context.Context) error {
		return nil
	}, matcherFunc)

	assert.Equal(t, 1, len(handler.HandlersPipe()))
	h := handler.HandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(h.MatcherFuncs))
	assert.Equal(t, matcherFunc, h.MatcherFuncs[0])

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

/*
func TestMapRest(t *testing.T) {

	rest := new(controllers_test.TestController)

	handler := NewHttpHandler()
	handler.MapRest("people", rest)

	assert.Equal(t, 1, len(handler.HandlersPipe()))

}
*/
