package goweb

import (
	"github.com/stretchr/goweb/handlers"
)

// defaultHttpHandler is the internal placeholder for the DefaultHttpHandler.
var defaultHttpHandler *handlers.HttpHandler

// DefaultHttpHandler gets the HttpHandler that can be used to handle
// requests.
//
// If nothing has been set using SetDefaultHttpHandler(), a fresh one
// will be created and served each time this function is called.
func DefaultHttpHandler() *handlers.HttpHandler {

	if defaultHttpHandler == nil {
		defaultHttpHandler = handlers.NewHttpHandler(CodecService)
	}

	return defaultHttpHandler

}

// SetDefaultHttpHandler sets the HttpHandler that will be used to
// handle requests.
//
// You do not need to call this, as calling DefaultHttpHandler() will
// create one if none has been set.
//
// Calling `SetDefaultHttpHandler(nil)` will cause a new HttpHandler to be
// created next time one is needed.
//
// Writing tests
//
// When writing tests for Goweb, you should:
//
//     # Create your own HttpHandler,
//     # use SetDefaultHttpHandler() to tell goweb to use it,
//     # call the code being tested (code that presumably calls things like `goweb.Map` etc.)
//     # make assertions against your own HttpHandler
//     # call SetDefaultHttpHandler(nil) to clean things up (not necessary)
func SetDefaultHttpHandler(handler *handlers.HttpHandler) {
	defaultHttpHandler = handler
}
