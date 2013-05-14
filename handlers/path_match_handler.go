package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
	"strings"
)

// PathMatchHandler is a Handler that maps a path to handler code.
type PathMatchHandler struct {

	// PathPattern is the pattern which paths must match in order for this
	// object to handle the context.
	PathPattern *paths.PathPattern

	// ExecutionFunc is the function that will be executed if there is a successful
	// match.
	ExecutionFunc HandlerExecutionFunc

	// MatcherFuncs are additional functions that are each consulted until a decision is
	// made as to whether this object will handle the context or not.
	MatcherFuncs []MatcherFunc

	// HttpMethods contains a list of HTTP Methods that will match, or an empty list if all
	// methods match.
	HttpMethods []string
}

// NewPathMatchHandler makes a new PathMatchHandler with the specified PathPattern
// and HandlerExecutionFunc.
func NewPathMatchHandler(pathPattern *paths.PathPattern, executionFunc HandlerExecutionFunc) *PathMatchHandler {
	handler := new(PathMatchHandler)
	handler.PathPattern = pathPattern
	handler.ExecutionFunc = executionFunc
	return handler
}

// WillHandle checks whether this handler will be used to handle the specified
// request or not.
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

// String gets a human readable string describing this PathMatchHandler.
func (p *PathMatchHandler) String() string {

	var methods string
	if len(p.HttpMethods) > 0 {
		methods = fmt.Sprintf("%s ", strings.Join(p.HttpMethods, "|"))
	}

	return fmt.Sprintf("%s%v\t\t\t%v - %d matcher func(s).", methods, p.PathPattern.RawPath, p.ExecutionFunc, len(p.MatcherFuncs))
}
