package handlers

import (
	"github.com/stretchr/goweb/context"
)

// Handler represents an object capable of handling a request.
type Handler interface {

	// WillHandle gets whether this handler will have its Handle method
	// called for the specified Context or not.
	WillHandle(context.Context) (bool, error)

	// Handle tells the handler to do its work.
	//
	// If stop is returned as true, no more handlers in the current pipeline
	// will be executed.
	//
	// If an error is returned, it will be handed to the HttpHandler.ErrorHandler
	// to deal with.  Ideally, you should implement your own custom error handling
	// mechanism instead of returning errors, and use that for system errors.
	Handle(context.Context) (stop bool, err error)
}

// HandlerExecutionFunc represents a function that can handle requests.
type HandlerExecutionFunc func(context.Context) error
