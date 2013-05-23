package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
)

type DefaultErrorHandler struct{}

// WillHandle is ignored on ErrorHandlers.
func (h *DefaultErrorHandler) WillHandle(context.Context) (bool, error) {
	return true, nil
}

// Handle writes the error from the context into the HttpResponseWriter.
func (h *DefaultErrorHandler) Handle(ctx context.Context) (stop bool, err error) {

	// write the error out
	ctx.HttpResponseWriter().Write([]byte(fmt.Sprintf("Oops, something went wrong: %s", ctx.Data().Get("error"))))

	// responses are actually ignored
	return false, nil
}
