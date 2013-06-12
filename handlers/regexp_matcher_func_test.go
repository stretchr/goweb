package handlers

import (
	"github.com/stretchr/goweb/context"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexPath(t *testing.T) {

	pattern := `^[a-z]+\[[0-9]+\]$`
	matcherFunc := RegexPath(pattern)

	var ctx context.Context
	var decision MatcherFuncDecision

	ctx = context_test.MakeTestContextWithPath("adam[23]")
	decision, _ = matcherFunc(ctx)
	assert.Equal(t, Match, decision, "adam[23] should match")

	ctx = context_test.MakeTestContextWithPath("eve[7]")
	decision, _ = matcherFunc(ctx)
	assert.Equal(t, Match, decision, "eve[7] should match")

	ctx = context_test.MakeTestContextWithPath("Job[23]")
	decision, _ = matcherFunc(ctx)
	assert.Equal(t, NoMatch, decision, "Job[23] should NOT match")

	ctx = context_test.MakeTestContextWithPath("snakey")
	decision, _ = matcherFunc(ctx)
	assert.Equal(t, NoMatch, decision, "snakey should NOT match")

}

func TestRegexPath_Error(t *testing.T) {

	pattern := `[0-9]++`
	matcherFunc := RegexPath(pattern)

	var ctx context.Context

	ctx = context_test.MakeTestContextWithPath("adam[23]")
	_, err := matcherFunc(ctx)
	assert.Error(t, err)

}
