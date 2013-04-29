package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

/*
  Pipe represents a collection of handlers.
*/
type Pipe []Handler

/*
  AppendHandler adds a handler to the end of this pipe.
*/
func (p Pipe) AppendHandler(handler Handler) Pipe {
	return append(p, handler)
}

/*
  PrependHandler adds a handler to the start of this pipe.
*/
func (p Pipe) PrependHandler(handler Handler) Pipe {

	// TODO: is there a better way to do prepends?

	handlers := make([]Handler, len(p)+1)
	handlers[0] = handler
	for hIndex, handler := range p {
		handlers[hIndex+1] = handler
	}

	return handlers
}

/*
  WillHandle always return true for Pipes.
*/
func (p Pipe) WillHandle(*context.Context) (bool, error) {
	return true, nil
}

/*
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p Pipe) Handle(c *context.Context) error {

	var willHandle bool
	var willHandleErr error
	var handleErr error

	for _, handler := range p {

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
