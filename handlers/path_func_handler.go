package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

/**
  PathFuncHandler is a Handler that maps a path to handler code.
*/
type PathFuncHandler struct {
	Path string

	HandlerFunc HandlerFunc
}

/**
  WillHandle always return true for Pipes.
*/
func (p *PathFuncHandler) WillHandle(c *context.Context) (bool, error) {
	return false, nil
}

/**
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p *PathFuncHandler) Handle(c *context.Context) error {
	return p.HandlerFunc(c)
}
