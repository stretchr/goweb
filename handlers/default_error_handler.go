package handlers

import (
	"fmt"
	"github.com/stretchr/goweb/context"
	"net/http"
	"os"
	"reflect"
)

// DefaultErrorHandler is a Handler that writes an error message out
// to the client.
//
// The error will be stored in the context.Data with the DataKeyForError key.
type DefaultErrorHandler struct{}

// WillHandle is ignored on ErrorHandlers.
func (h *DefaultErrorHandler) WillHandle(context.Context) (bool, error) {
	return true, nil
}

// Handle writes the error from the context into the HttpResponseWriter with a
// 500 http.StatusInternalServerError status code.
func (h *DefaultErrorHandler) Handle(ctx context.Context) (stop bool, err error) {

	var handlerError HandlerError = ctx.Data().Get(DataKeyForError).Data().(HandlerError)
	hostname, _ := os.Hostname()

	w := ctx.HttpResponseWriter()

	// write the error out
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("<!DOCTYPE html><html><head>"))
	w.Write([]byte("<style>"))
	w.Write([]byte("h1 { font-size: 17px }"))
	w.Write([]byte("h1 strong {text-decoration:underline}"))
	w.Write([]byte("h2 { background-color: #ffd; padding: 20px }"))
	w.Write([]byte("footer { margin-top: 20px; border-top:1px solid black; padding:10px; font-size:0.9em }"))
	w.Write([]byte("</style>"))
	w.Write([]byte("</head><body>"))
	w.Write([]byte(fmt.Sprintf("<h1>Error in <code>%s</code></h1><h2>%s</h2>", handlerError.Handler, handlerError)))
	w.Write([]byte(fmt.Sprintf("<h3><code>%s</code> error in Handler <code>%v</code></h3> <code><pre>%s</pre></code>", reflect.TypeOf(handlerError.OriginalError), &handlerError.Handler, handlerError.Handler)))
	w.Write([]byte(fmt.Sprintf("on %s", hostname)))
	w.Write([]byte("<footer>Learn more about <a href='http://github.com/stretchr/goweb' target='_blank'>Goweb</a></footer>"))
	w.Write([]byte("</body></html>"))

	// responses are actually ignored
	return false, nil
}
