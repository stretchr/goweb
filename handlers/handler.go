package handlers

import (
	"github.com/stretchrcom/goweb/context"
)

/*
	Handler represents an object capable of handling a request.
*/
type Handler interface {

	/*
		WillHandle gets whether this 
	*/
	WillHandle(* context.Context) (bool, error)

}