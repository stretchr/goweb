package api

import (
	"github.com/stretchrcom/goweb/context"
)

/*
  APIResponder represents objects capable of provide API responses.
*/
type APIResponder interface {

	/*
	   Responding
	*/

	// Responds to the Context with the specified status, data and errors.
	Respond(ctx *context.Context, status int, data interface{}, errors []string)
}
