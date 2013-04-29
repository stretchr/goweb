package handlers

import (
	//"github.com/stretchrcom/goweb/context"
	"net/http"
)

type HttpHandler struct {
	handlers Pipe
}

func NewHttpHandler() *HttpHandler {
	h := new(HttpHandler)
	h.handlers = make(Pipe, 3)

	// make the default pipes
	h.handlers[0] = make(Pipe, 0)
	h.handlers[1] = make(Pipe, 0)
	h.handlers[2] = make(Pipe, 0)

	return h
}

func (handler *HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	// make the context
	//ctx := context.NewContext(request.URL.Path)

	// run it through the handlers

}

/*
  Handlers
*/

func (h *HttpHandler) Handlers() Pipe {
	return h.handlers[1].(Pipe)
}

func (h *HttpHandler) BeforeHandlers() Pipe {
	return h.handlers[0].(Pipe)
}

func (h *HttpHandler) AfterHandlers() Pipe {
	return h.handlers[2].(Pipe)
}
