package handlers

import (
	"github.com/stretchrcom/goweb/webcontext"
	"net/http"
)

type HttpHandler struct {

	// Handlers represent a pipe of handlers that will be used
	// to handle requests.
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
	ctx := webcontext.NewWebContext(responseWriter, request)

	// run it through the handlers
	_, err := handler.Handlers.Handle(ctx)

	if err != nil {

		// TODO: handle errors

	}

}

// HandlersPipe gets the pipe for handlers.
func (h *HttpHandler) HandlersPipe() Pipe {
	return h.Handlers[1].(Pipe)
}

// PreHandlersPipe gets the handlers that are executed before processing begins.
func (h *HttpHandler) PreHandlersPipe() Pipe {
	return h.Handlers[0].(Pipe)
}

// PostHandlersPipe gets the handlers that are executed after processing completes.
func (h *HttpHandler) PostHandlersPipe() Pipe {
	return h.Handlers[2].(Pipe)
}

// AppendHandler appends a handler to the processing pipe.
func (h *HttpHandler) AppendHandler(handler Handler) {
	h.Handlers[1] = h.HandlersPipe().AppendHandler(handler)
}

// AppendPreHandler appends a handler to be executed before processing begins.
func (h *HttpHandler) AppendPreHandler(handler Handler) {
	h.Handlers[0] = h.PreHandlersPipe().AppendHandler(handler)
}

// AppendPostHandler appends a handler to be executed after processing completes.
func (h *HttpHandler) AppendPostHandler(handler Handler) {
	h.Handlers[2] = h.PostHandlersPipe().AppendHandler(handler)
}
