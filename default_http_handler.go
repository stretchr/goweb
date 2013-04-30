package goweb

import (
	"github.com/stretchrcom/goweb/handlers"
)

var defaultHttpHandler *handlers.HttpHandler

func DefaultHttpHandler() *handlers.HttpHandler {

	if defaultHttpHandler == nil {
		defaultHttpHandler = handlers.NewHttpHandler()
	}

	return defaultHttpHandler

}
