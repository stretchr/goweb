package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	"net/http"
	"os"
	"reflect"
)

type DefaultErrorHandler struct{}

// WillHandle is ignored on ErrorHandlers.
func (h *DefaultErrorHandler) WillHandle(context.Context) (bool, error) {
	return true, nil
}

// Handle writes the error from the context into the HttpResponseWriter with a
// 500 http.StatusInternalServerError status code. 
func (h *DefaultErrorHandler) Handle(ctx context.Context) (stop bool, err error) {

	var handlerError HandlerError = ctx.Data().Get("error").(HandlerError)
	hostname, _ := os.Hostname()

	w := ctx.HttpResponseWriter()

	// write the error out
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("<!DOCTYPE html><html><head>"))
	w.Write([]byte("<style>"))
	w.Write([]byte("h1 strong {text-decoration:underline}"))
	w.Write([]byte("h2 { background-color: #ddd; padding: 20px }"))
	w.Write([]byte("footer { margin-top: 20px; border-top:1px solid black; padding:10px; font-size:0.9em }"))
	w.Write([]byte("</style>"))
	w.Write([]byte("</head><body>"))
	w.Write([]byte(fmt.Sprintf("<h1>Error in <strong>%s</strong></h1><h2>%s</h2>", handlerError.Handler, handlerError)))
	w.Write([]byte(fmt.Sprintf("<code>%s</code> error in Handler <code>%v</code> <code><pre>%s</pre></code>", reflect.TypeOf(handlerError.OriginalError), &handlerError.Handler, handlerError.Handler)))
	w.Write([]byte(fmt.Sprintf("on %s", hostname)))
	w.Write([]byte("<footer>Learn more about <a href='http://github.com/stretchrcom/goweb' target='_blank'>Goweb</a></footer>"))
	w.Write([]byte("</body></html>"))

	// responses are actually ignored
	return false, nil
}
