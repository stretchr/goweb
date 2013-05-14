package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

/*
	Handler represents an object capable of handling a request.
*/
type Handler interface {

	/*
		WillHandle gets whether this handler will have its Handle method
		called for the specified Context or not.
	*/
	WillHandle(context.Context) (bool, error)

	/*
	   Handle tells the handler to do its work.
	*/
	Handle(context.Context) (stop bool, err error)
}

/*
	HandlerExecutionFunc represents a function that can handle requests.
*/
type HandlerExecutionFunc func(context.Context) error
