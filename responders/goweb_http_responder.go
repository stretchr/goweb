package responders

import (
	"github.com/stretchrcom/goweb/context"
	"net/http"
)

type GowebHTTPResponder struct {
}

// With writes a response to the request in the specified context.
func (r *GowebHTTPResponder) With(ctx context.Context, httpStatus int, body []byte) error {

	r.WithStatus(ctx, httpStatus)

	_, writeErr := ctx.HttpResponseWriter().Write(body)
	return writeErr

}

// WithStatus writes the specified HTTP Status Code to the Context's ResponseWriter.
func (r *GowebHTTPResponder) WithStatus(ctx context.Context, httpStatus int) error {
	ctx.HttpResponseWriter().WriteHeader(httpStatus)
	return nil
}

// WithOK responds with a 200 OK status code, and no body.
func (r *GowebHTTPResponder) WithOK(ctx context.Context) error {
	return r.WithStatus(ctx, http.StatusOK)
}

// WithRedirect responds with a redirection to the specific path or URL.
func (r *GowebHTTPResponder) WithRedirect(ctx context.Context, pathOrURL string) error {

	ctx.HttpResponseWriter().Header().Set("Location", pathOrURL)
	return r.WithStatus(ctx, http.StatusTemporaryRedirect)

}
