package handlers

import (
	"net/http"
)

type HttpHandler struct {
	handlers []Handler
}

func NewHttpHandler() *HttpHandler {
	h := new(HttpHandler)
	h.handlers = make([]Handler, 3)

	// make the default pipes
	h.handlers[0] = new(Pipe)
	h.handlers[1] = new(Pipe)
	h.handlers[2] = new(Pipe)

	return h
}

func (handler *HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

}

/*
  Handlers
*/

func (h *HttpHandler) Handlers() *Pipe {
	return h.handlers[1].(*Pipe)
}
func (h *HttpHandler) BeforeHandlers() *Pipe {
	return h.handlers[0].(*Pipe)
}
func (h *HttpHandler) AfterHandlers() *Pipe {
	return h.handlers[2].(*Pipe)
}
