package handlers

import (
	"fmt"
	codecsservices "github.com/stretchr/codecs/services"
	gowebhttp "github.com/stretchr/goweb/http"
	"github.com/stretchr/goweb/webcontext"
	"github.com/stretchr/objx"
	"net/http"
	"strings"
)

const (
	// DataKeyForError is the data key (that goes into the context.Data map) for
	// an error that has occurred.  Then, the error Handler can Get(DataKeyForError)
	// to do work on the error.
	DataKeyForError string = "error"
)

type HttpHandler struct {

	// codecServices is the codec service object to use to go from bytes to objects
	// and vice versa.
	codecService codecsservices.CodecService

	// Handlers represent a pipe of handlers that will be used
	// to handle requests.
	Handlers Pipe

	// errorHandler represents the Handler that will be used to handle errors.
	errorHandler Handler

	// Data contains the initial data object that gets copied to each
	// context object.
	Data objx.Map

	// HttpMethodForCreate is the HTTP method to use for this action when mapping controllers.
	HttpMethodForCreate string
	// HttpMethodForReadOne is the HTTP method to use for this action when mapping controllers.
	HttpMethodForReadOne string
	// HttpMethodForReadMany is the HTTP method to use for this action when mapping controllers.
	HttpMethodForReadMany string
	// HttpMethodForDeleteOne is the HTTP method to use for this action when mapping controllers.
	HttpMethodForDeleteOne string
	// HttpMethodForDeleteMany is the HTTP method to use for this action when mapping controllers.
	HttpMethodForDeleteMany string
	// HttpMethodForUpdateOne is the HTTP method to use for this action when mapping controllers.
	HttpMethodForUpdateOne string
	// HttpMethodForUpdateMany is the HTTP method to use for this action when mapping controllers.
	HttpMethodForUpdateMany string
	// HttpMethodForReplace is the HTTP method to use for this action when mapping controllers.
	HttpMethodForReplace string
	// HttpMethodForHead is the HTTP method to use for this action when mapping controllers.
	HttpMethodForHead string
	// HttpMethodForOptions is the HTTP method to use for this action when mapping controllers.
	HttpMethodForOptions string
}

// NewHttpHandler creates a new HttpHandler obejct with the specified CodecService.
//
// New HttpHandlers will be initialised with three handler Pipes:
//
//     0 - Pre handlers
//     1 - Main handlers
//     2 - Post handlers
func NewHttpHandler(codecService codecsservices.CodecService) *HttpHandler {
	h := new(HttpHandler)

	// make empty data
	h.Data = make(objx.Map)

	// make pre, process and post pipes
	h.Handlers = make(Pipe, 3)
	h.Handlers[0] = make(Pipe, 0) // pre
	h.Handlers[1] = make(Pipe, 0) // process
	h.Handlers[2] = make(Pipe, 0) // post

	// save the codec service
	h.codecService = codecService

	// assign default HTTP methods
	h.HttpMethodForCreate = gowebhttp.MethodPost
	h.HttpMethodForReadOne = gowebhttp.MethodGet
	h.HttpMethodForReadMany = gowebhttp.MethodGet
	h.HttpMethodForDeleteOne = gowebhttp.MethodDelete
	h.HttpMethodForDeleteMany = gowebhttp.MethodDelete
	h.HttpMethodForUpdateOne = gowebhttp.MethodPatch
	h.HttpMethodForUpdateMany = gowebhttp.MethodPatch
	h.HttpMethodForReplace = gowebhttp.MethodPut
	h.HttpMethodForHead = gowebhttp.MethodHead
	h.HttpMethodForOptions = gowebhttp.MethodOptions

	return h
}

// CodecService gets the codec service that this HttpHandler will use to
// marshal and unmarshal objects to and from data.
func (handler *HttpHandler) CodecService() codecsservices.CodecService {
	return handler.codecService
}

// ServeHTTP servers the actual HTTP request by buidling a context and running
// it through all the handlers.
func (handler *HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// override the method if needed
	method := request.Header.Get("X-HTTP-Method-Override")
	if method != "" {
		request.Method = method
	}

	// make the context
	ctx := webcontext.NewWebContext(responseWriter, request, handler.codecService)

	// copy the data
	for k, v := range handler.Data {
		ctx.Data()[k] = v
	}

	// run it through the handlers
	_, err := handler.Handlers.Handle(ctx)

	// do we need to handle an error?
	if err != nil {

		// set the error
		ctx.Data().Set(DataKeyForError, err)

		// tell the handler to handle it
		handler.ErrorHandler().Handle(ctx)

	}

}

// ErrorHandler gets the Handler that will be used to handle errors.
//
// If no error handler has been set, a default error handler will be returned
// which will just write the error out in plain text.  If you are building an API,
// it is recommended that you roll your own ErrorHandler.
//
// For more information on rolling your own ErrorHandler, see the SetErrorHandler
// method.
func (h *HttpHandler) ErrorHandler() Handler {

	if h.errorHandler == nil {

		h.errorHandler = &DefaultErrorHandler{}

	}

	return h.errorHandler
}

// SetErrorHandler sets the Handler that will be used to handle errors.
//
// The error handler is like a normal Handler, except with a few oddities.
// The WillHandle method never gets called on the ErrorHandler, and any errors
// returned from the Handle method are ignored (as is the stop argument).
// If you want to log errors, you should do so from within the ErrorHandler.
//
// Goweb will place the error object into the context.Data() map with the
// DataKeyForError key.
func (h *HttpHandler) SetErrorHandler(errorHandler Handler) {
	h.errorHandler = errorHandler
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

// PrepentPreHandler prepends a handler to be executed before processing begins.
func (h *HttpHandler) PrependPreHandler(handler Handler) {
	h.Handlers[0] = h.PreHandlersPipe().PrependHandler(handler)
}

// AppendPostHandler appends a handler to be executed after processing completes.
func (h *HttpHandler) AppendPostHandler(handler Handler) {
	h.Handlers[2] = h.PostHandlersPipe().AppendHandler(handler)
}

// PrependPostHandler prepends a handler to be executed after processing completes.
func (h *HttpHandler) PrependPostHandler(handler Handler) {
	h.Handlers[2] = h.PostHandlersPipe().PrependHandler(handler)
}

/*
	Debug and information
*/

// String generates a list of the handlers registered inside this HttpHandler.
func (h *HttpHandler) String() string {
	return stringForHandlers(h.Handlers, 0)
}

// stringForHandlers generates the string for the handlers array indented to the
// appropriate level.
func stringForHandlers(handlers []Handler, level int) string {

	lines := []string{}
	var levelStr string = strings.Repeat("  ", level)

	for handlerIndex, handler := range handlers {
		if pipe, ok := handler.(Pipe); ok {
			lines = append(lines, fmt.Sprintf("\n%sPipe %d:\n", levelStr, handlerIndex))
			lines = append(lines, stringForHandlers(pipe, level+1))
		} else {
			lines = append(lines, fmt.Sprintf("%s%s", levelStr, handler))
		}
	}

	return strings.Join(lines, "")

}
