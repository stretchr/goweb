package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
)

const (
	contextURLParametersDataKey string = "path-params"
)

/*
  PathMatchHandler is a Handler that maps a path to handler code.
*/
type PathMatchHandler struct {
	PathPattern   *paths.PathPattern
	ExecutionFunc HandlerExecutionFunc
}

/*
  WillHandle checks whether this handler will be used to handle the specified
  request or not.
*/
func (p *PathMatchHandler) WillHandle(c *context.Context) (bool, error) {
	match := p.PathPattern.GetPathMatch(c.Path())

	if match.Matches {

		// save the match parameters for later
		c.Data().Set(contextURLParametersDataKey, match.Parameters)

	}

	return match.Matches, nil
}

/*
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p *PathMatchHandler) Handle(c *context.Context) error {
	return p.ExecutionFunc(c)
}
