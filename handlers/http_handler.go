package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"net/http"
)

type HttpHandler struct {
	Handlers Pipe
}

func NewHttpHandler() *HttpHandler {
	h := new(HttpHandler)

	// make pre, process and post pipes
	h.Handlers = make(Pipe, 3)
	h.Handlers[0] = make(Pipe, 0) // pre
	h.Handlers[1] = make(Pipe, 0) // process
	h.Handlers[2] = make(Pipe, 0) // post

	return h
}

func (handler *HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	// make the context
	ctx := context.NewContext(handler, responseWriter, request)

	// run it through the handlers
	_, err := handler.Handlers.Handle(ctx)

	if err != nil {

		// TODO: handle errors

	}

}
