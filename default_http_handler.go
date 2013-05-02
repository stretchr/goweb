package goweb

import (
	"github.com/stretchrcom/goweb/handlers"
)

// defaultHttpHandler is the internal placeholder for the DefaultHttpHandler.
var defaultHttpHandler *handlers.HttpHandler

// DefaultHttpHandler gets the HttpHandler that can be used to handle
// requests.
func DefaultHttpHandler() *handlers.HttpHandler {

	if defaultHttpHandler == nil {
		defaultHttpHandler = handlers.NewHttpHandler()
	}

	return defaultHttpHandler

}
