package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
)

/*
  PathMatchHandler is a Handler that maps a path to handler code.
*/
type PathMatchHandler struct {
	PathPattern   *paths.PathPattern
	ExecutionFunc HandlerExecutionFunc
	MatcherFuncs  []MatcherFunc
}

func NewPathMatchHandler(pathPattern *paths.PathPattern, executionFunc HandlerExecutionFunc) *PathMatchHandler {
	handler := new(PathMatchHandler)
	handler.PathPattern = pathPattern
	handler.ExecutionFunc = executionFunc
	return handler
}

/*
  WillHandle checks whether this handler will be used to handle the specified
  request or not.
*/
func (p *PathMatchHandler) WillHandle(c context.Context) (bool, error) {

	// check each matcher func
	matcherFuncMatches := true
	matcherFuncDecisionMade := false
	for _, matcherFunc := range p.MatcherFuncs {
		decision, matcherFuncErr := matcherFunc(c)

		if matcherFuncErr != nil {
			return false, matcherFuncErr
		}

		switch decision {
		case NoMatch:
			matcherFuncMatches = false
			matcherFuncDecisionMade = true
			break
		case Match:
			matcherFuncMatches = true
			matcherFuncDecisionMade = true
			break
		}

		if matcherFuncDecisionMade {
			break
		}

	}

	pathMatch := p.PathPattern.GetPathMatch(c.Path())

	var allMatch bool

	if matcherFuncDecisionMade {
		allMatch = matcherFuncMatches
	} else {
		allMatch = pathMatch.Matches
	}

	if allMatch {

		// save the match parameters for later
		c.Data().Set(context.DataKeyPathParameters, pathMatch.Parameters)

	}

	return allMatch, nil
}

/*
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p *PathMatchHandler) Handle(c context.Context) (bool, error) {
	err := p.ExecutionFunc(c)
	return true, err
}
