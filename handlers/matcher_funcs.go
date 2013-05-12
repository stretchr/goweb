package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

// MatcherFuncDecision is the return type of MatcherFunc and should be
// one of, DontCare, NoMatch or Match.
type MatcherFuncDecision int8

const (
	DontCare MatcherFuncDecision = -1
	NoMatch  MatcherFuncDecision = 0
	Match    MatcherFuncDecision = 1
)

// MatcherFunc is a function capable of helping PathMatchHandlers decide whether
// it should handle a specific context.  While the path is matched automatically, MatcherFuncs
// allow you to specify additional checks or constraints.
type MatcherFunc func(context.Context) (MatcherFuncDecision, error)
