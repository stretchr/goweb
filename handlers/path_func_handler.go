package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
)

const (
	contextURLParametersDataKey string = "urlparams"
)

/**
  PathFuncHandler is a Handler that maps a path to handler code.
*/
type PathFuncHandler struct {
	PathPattern *paths.PathPattern

	HandlerFunc HandlerFunc
}

/**
  WillHandle always return true for Pipes.
*/
func (p *PathFuncHandler) WillHandle(c *context.Context) (bool, error) {
	match := p.PathPattern.GetPathMatch(c.Path())

	// save the match parameters for later
	c.Data().Set(contextURLParametersDataKey, match.Parameters)

	return match.Matches, nil
}

/**
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p *PathFuncHandler) Handle(c *context.Context) error {
	return p.HandlerFunc(c)
}
