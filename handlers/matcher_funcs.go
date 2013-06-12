package handlers

import (
	"github.com/stretchr/goweb/context"
)

// MatcherFuncDecision is the return type of MatcherFunc and should be
// one of, DontCare, NoMatch or Match.
type MatcherFuncDecision int8

const (
	// Indicates that the MatcherFunc does not have an opinion about whether this
	// matches or not, and it will leave it to other MatcherFuncs (or the PathPattern)
	// to decide.
	DontCare MatcherFuncDecision = -1

	// NoMatch indicates that the handler does not match.
	NoMatch MatcherFuncDecision = 0

	// Match indicates that the handler does match and should handle the Context.
	Match MatcherFuncDecision = 1
)

// MatcherFunc is a function capable of helping PathMatchHandlers decide whether
// it should handle a specific context.  While the path is matched automatically, MatcherFuncs
// allow you to specify additional checks or constraints.
type MatcherFunc func(context.Context) (MatcherFuncDecision, error)
