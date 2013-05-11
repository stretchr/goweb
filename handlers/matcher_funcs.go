package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

type MatcherFuncDecision int8

const (
	DontCare MatcherFuncDecision = -1
	NoMatch  MatcherFuncDecision = 0
	Match    MatcherFuncDecision = 1
)

type MatcherFunc func(context.Context) (MatcherFuncDecision, error)
