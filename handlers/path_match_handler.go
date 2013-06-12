package handlers

import (
	"fmt"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/paths"
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

	// Description is an optional string that describes the mapping.  If present, it will
	// be returned instead of the default when String() is called.
	Description string

	// BreakCurrentPipeline indicates whether the rest of the handlers in the Pipe
	// should be skipped once this handler has done its work.
	//
	// Usually for the main pipe, these handlers will skip as they are responsible
	// for actually returning a response to the clients.
	//
	// In Before and After handlers however, the likelihood is that all handlers
	// in the pipe will want to do something without being skipped.
	//
	// goweb.Map will set this value to true, whereas goweb.MapBefore and goweb.MapAfter
	// will set it to false.
	BreakCurrentPipeline bool
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

	// check HTTP methods
	var httpMethodMatch bool = false

	if len(p.HttpMethods) == 0 {

		// no specific HTTP methods
		httpMethodMatch = true

	} else {

		for _, httpMethod := range p.HttpMethods {
			if httpMethod == c.MethodString() {
				httpMethodMatch = true
				break
			}
		}

	}

	// cancel early if we didn't get an HTTP Method match
	if !httpMethodMatch {
		return false, nil
	}

	// check path match

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
	return p.BreakCurrentPipeline, err
}

// String gets a human readable string describing this PathMatchHandler.
func (p *PathMatchHandler) String() string {

	var desc string

	if len(p.Description) > 0 {
		desc = p.Description
	} else {
		desc = fmt.Sprintf("%v", p.ExecutionFunc)
	}

	var methods string
	if len(p.HttpMethods) > 0 {
		methods = fmt.Sprintf("%s ", strings.Join(p.HttpMethods, "|"))
	}

	var matcherFuncsDesc string

	if len(p.MatcherFuncs) > 0 {
		matcherFuncsDesc = fmt.Sprintf(" - %d matcher func(s)", len(p.MatcherFuncs))
	}

	return fmt.Sprintf("%s%v - %v %s\n", methods, p.PathPattern.RawPath, desc, matcherFuncsDesc)
}
