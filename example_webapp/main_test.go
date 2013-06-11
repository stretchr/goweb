package main

import (
	"github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb"
	"github.com/stretchrcom/goweb/handlers"
	//"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestRoutes(t *testing.T) {

	// make a test HttpHandler and use it
	codecService := new(services.WebCodecService)
	handler := handlers.NewHttpHandler(codecService)
	goweb.SetDefaultHttpHandler(handler)

	// call the target code
	mapRoutes()

	// TODO: make assertions

	/*

	   we want to be able to do things like this:

	     handlers := handler.GetHandlersFor("GET", "people/123")
	     h := handler.HandlerForRequest(request)

	*/

}
