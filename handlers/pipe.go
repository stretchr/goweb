package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

/**
  Pipe represents a collection of handlers.
*/
type Pipe struct {
	handlers []Handler
}

/**
  AddHandler adds a handler to this pipe.
*/
func (p *Pipe) AddHandler(handler Handler) *Pipe {
	p.handlers = append(p.handlers, handler)
	return p
}

/**
  WillHandle always return true for Pipes.
*/
func (p *Pipe) WillHandle(*context.Context) (bool, error) {
	return true, nil
}

/**
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p *Pipe) Handle(c *context.Context) error {

	var willHandle bool
	var willHandleErr error
	var handleErr error

	for _, handler := range p.handlers {

		willHandle, willHandleErr = handler.WillHandle(c)

		if willHandleErr != nil {
			return willHandleErr
		}

		if willHandle {

			// call the handler
			handleErr = handler.Handle(c)

			if handleErr != nil {
				return handleErr
			}

		}
	}

	// everything went well
	return nil
}
