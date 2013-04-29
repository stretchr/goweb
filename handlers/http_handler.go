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
	h.Handlers = make(Pipe, 0)

	return h
}

func (handler *HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	// make the context
	ctx := context.NewContext(responseWriter, request)

	// run it through the handlers
	err := handler.Handlers.Handle(ctx)

	if err != nil {

		// TODO: handle errors

	}

}
