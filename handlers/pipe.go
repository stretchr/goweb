package handlers

import (
	"github.com/stretchr/goweb/context"
)

/*
  Pipe represents a collection of handlers.
*/
type Pipe []Handler

/*
  AppendHandler adds a handler to the end of this pipe.  Returns a copy of the Pipe with the
  specified handler appended.
*/
func (p Pipe) AppendHandler(handler Handler) Pipe {
	return append(p, handler)
}

/*
  PrependHandler adds a handler to the start of this pipe.  Returns a copy of the Pipe with the
  specified handler prepended.
*/
func (p Pipe) PrependHandler(handler Handler) Pipe {

	handlers := make([]Handler, 0, len(p)+1)
	handlers = append(handlers, handler)
	handlers = append(handlers, p...)

	return handlers
}

/*
  WillHandle always return true for Pipes.
*/
func (p Pipe) WillHandle(context.Context) (bool, error) {
	return true, nil
}

/*
  Handle gives each sub handle the opportinuty to handle the context.
*/
func (p Pipe) Handle(c context.Context) (bool, error) {

	var willHandle bool
	var willHandleErr error
	var handleErr error
	var stop bool

	for _, handler := range p {

		willHandle, willHandleErr = handler.WillHandle(c)

		if willHandleErr != nil {
			return true, willHandleErr
		}

		if willHandle {

			// call the handler
			stop, handleErr = handler.Handle(c)

			if handleErr != nil {

				// already a HandlerError?
				if handlerError, ok := handleErr.(HandlerError); ok {

					// just return it plain
					return true, handlerError

				} else {

					// wrap it and record the handler that caused the error
					return true, HandlerError{handler, handleErr}

				}

			}

			if stop {
				break
			}

		}
	}

	// everything went well
	return false, nil
}
