package handlers

import (
	"fmt"
	"github.com/stretchr/goweb/context"
	//"github.com/stretchr/goweb/webcontext"
	"github.com/stretchr/goweb/paths"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPathMatchHandler(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("collection/{id}/name")
	var called bool = false
	h := NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error {
		called = true
		return nil
	}))

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := h.WillHandle(ctx1)
	assert.True(t, will)
	h.Handle(ctx1)
	assert.True(t, called, "Method should be called")
	assert.Equal(t, "123", ctx1.Data().Get(context.DataKeyPathParameters).ObjxMap().Get("id").Data())

}

func TestPathMatchHandler(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("collection/{id}/name")
	var called bool = false
	h := NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error {
		called = true
		return nil
	}))

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := h.WillHandle(ctx1)
	assert.True(t, will)
	h.Handle(ctx1)
	assert.True(t, called, "Method should be called")
	assert.Equal(t, "123", ctx1.Data().Get(context.DataKeyPathParameters).ObjxMap().Get("id").Data())

	ctx2 := context_test.MakeTestContextWithPath("/collection")
	will, _ = h.WillHandle(ctx2)
	assert.False(t, will)
	assert.Nil(t, ctx2.Data().Get(context.DataKeyPathParameters).Data())

	h.BreakCurrentPipeline = true
	shouldStop, handleErr := h.Handle(ctx2)
	assert.Nil(t, handleErr)
	assert.True(t, shouldStop)
	assert.True(t, called, "Handler func should get called")

}

func TestPathMatchHandler_BreakCurrentPipeline(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("collection/{id}/name")
	h := NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error {
		return nil
	}))
	h.BreakCurrentPipeline = true

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")

	breakCurrentPipeline, _ := h.Handle(ctx1)

	assert.True(t, breakCurrentPipeline)

	h = NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error {
		return nil
	}))
	h.BreakCurrentPipeline = false

	ctx1 = context_test.MakeTestContextWithPath("/collection/123/name")

	breakCurrentPipeline, _ = h.Handle(ctx1)

	assert.False(t, breakCurrentPipeline)

}

func assertPathMatches(t *testing.T, pattern, path string, shouldMatch bool) bool {

	pathPattern, _ := paths.NewPathPattern(pattern)
	ctx := context_test.MakeTestContextWithPath(path)

	h := NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error { return nil }))

	willHandle, _ := h.WillHandle(ctx)
	return assert.Equal(t, shouldMatch, willHandle, fmt.Sprintf("WillHandle should be %v for '%s' with path '%s'.", shouldMatch, pattern, path))

}

func TestPathMatches(t *testing.T) {

	// people/{id}/books
	assertPathMatches(t, "/people/{id}/books", "people/123/books", true)
	assertPathMatches(t, "/people/{id}/books", "/people/123/books", true)
	assertPathMatches(t, "/people/{id}/books", "people/123", false)
	assertPathMatches(t, "people/{id}/books", "people/123/books", true)
	assertPathMatches(t, "people/{id}/books", "/people/123/books", true)
	assertPathMatches(t, "people/{id}/books", "people/123", false)

	// /
	assertPathMatches(t, "/", "", true)
	assertPathMatches(t, "/", "/", true)
	assertPathMatches(t, "/", "/people", false)
	assertPathMatches(t, "/", "/123", false)

	// people/{id}/books/***
	assertPathMatches(t, "/people/{id}/books/***", "people/123/books/", true)
	assertPathMatches(t, "/people/{id}/books/***", "people/123/books", true)
	assertPathMatches(t, "/people/{id}/books/***", "people/123/books/123", true)
	assertPathMatches(t, "/people/{id}/books/***", "people/123/books/123/something", true)
	assertPathMatches(t, "/people/{id}/books/***", "people/123/books/123/something/else", true)

}

func TestPathMatches_WithHttpMethod(t *testing.T) {

	pathPattern, _ := paths.NewPathPattern("people/123")

	h := NewPathMatchHandler(pathPattern, HandlerExecutionFunc(func(c context.Context) error { return nil }))
	h.HttpMethods = []string{"POST"}

	ctx := context_test.MakeTestContextWithDetails("/people/123", "GET")

	willHandle, _ := h.WillHandle(ctx)
	assert.Equal(t, false, willHandle)

	ctx2 := context_test.MakeTestContextWithDetails("/people/123", "POST")

	willHandle, _ = h.WillHandle(ctx2)
	assert.Equal(t, true, willHandle)

}

func TestPathMatchHandler_WithMatcherFuncs_NoMatch(t *testing.T) {

	matcherFuncCalled := false

	handler := new(PathMatchHandler)
	handler.PathPattern, _ = paths.NewPathPattern("***")
	handler.ExecutionFunc = HandlerExecutionFunc(func(c context.Context) error {
		return nil
	})

	handler.MatcherFuncs = []MatcherFunc{func(c context.Context) (MatcherFuncDecision, error) {
		matcherFuncCalled = true
		return NoMatch, nil
	}}

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := handler.WillHandle(ctx1)

	assert.False(t, will, "Should not want to handle even though the path matches")

}

func TestPathMatchHandler_WithMatcherFuncs_Match(t *testing.T) {

	matcherFuncCalled := false

	handler := new(PathMatchHandler)
	handler.PathPattern, _ = paths.NewPathPattern("/specific/things")
	handler.ExecutionFunc = HandlerExecutionFunc(func(c context.Context) error {
		return nil
	})

	handler.MatcherFuncs = []MatcherFunc{func(c context.Context) (MatcherFuncDecision, error) {
		matcherFuncCalled = true
		return Match, nil
	}}

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := handler.WillHandle(ctx1)

	assert.True(t, will, "Should want to handle even though the path DOESNT match")

}

func TestPathMatchHandler_WithMatcherFuncs_NoMatch_Then_Match(t *testing.T) {

	matcherFuncCalled := false

	handler := new(PathMatchHandler)
	handler.PathPattern, _ = paths.NewPathPattern("/specific/things")
	handler.ExecutionFunc = HandlerExecutionFunc(func(c context.Context) error {
		return nil
	})

	handler.MatcherFuncs = []MatcherFunc{func(c context.Context) (MatcherFuncDecision, error) {
		matcherFuncCalled = true
		return NoMatch, nil
	}, func(c context.Context) (MatcherFuncDecision, error) {
		matcherFuncCalled = true
		return Match, nil
	}}

	ctx1 := context_test.MakeTestContextWithPath("/collection/123/name")
	will, _ := handler.WillHandle(ctx1)

	assert.False(t, will, "Should NOT want to handle even though the path does match because a decision is made early")

}
