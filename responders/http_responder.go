package responders

import (
	"github.com/stretchr/goweb/context"
)

type HTTPResponder interface {

	// WithStatus writes the specified HTTP Status Code to the Context's ResponseWriter.
	WithStatus(ctx context.Context, httpStatus int) error

	// WithStatusText writes the specified HTTP Status Code to the Context's ResponseWriter and
	// includes a body with the default status text.
	WithStatusText(ctx context.Context, httpStatus int) error

	// WithOK responds with a 200 OK status code, and no body.
	WithOK(ctx context.Context) error

	// With writes a response to the request in the specified context.
	With(ctx context.Context, httpStatus int, body []byte) error

	// WithRedirect responds with a redirection to the specific path or URL with the
	// http.StatusTemporaryRedirect status.
	WithRedirect(ctx context.Context, pathOrURLSegments ...interface{}) error

	// WithPermanentRedirect responds with a redirection to the specific path or URL with the
	// http.StatusMovedPermanently status.
	WithPermanentRedirect(ctx context.Context, pathOrURLSegments ...interface{}) error
}
