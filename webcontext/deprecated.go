package webcontext

import (
	"github.com/stretchrcom/goweb/context"
)

const (
	panicTextForDeprecatedResponses string = "goweb: context.Respond is deprecated, instead use the goweb.API.Respond methods."
)

// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
func (c *WebContext) Respond(data interface{}, statusCode int, errors []string, context context.Context) error {
	panic(panicTextForDeprecatedResponses)
}

// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
func (c *WebContext) RespondWithData(data interface{}) error {
	panic(panicTextForDeprecatedResponses)
}

// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
func (c *WebContext) RespondWithError(statusCode int) error {
	panic(panicTextForDeprecatedResponses)
}

// Deprecated: Code should be tweaked to use goweb.API.Respond methods instead.
func (c *WebContext) RespondWithErrorMessage(message string, statusCode int) error {
	panic(panicTextForDeprecatedResponses)
}
