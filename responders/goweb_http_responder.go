package responders

import (
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/paths"
	"net/http"
)

const (
	// DefaultAlways200ParamName is the default parameter name that tells the GowebHTTPResponder to always
	// return with a 200 status code.  By default, this will be "always200".
	DefaultAlways200ParamName string = "always200"
)

var (
	// Always200ParamName is the parameter name that tells the GowebHTTPResponder to always
	// return with a 200 status code.  By default, this will be "always200" or DefaultAlways200ParamName.
	Always200ParamName string = DefaultAlways200ParamName
)

// GowebHTTPResponder is the default HTTPResponder used to make responses.
type GowebHTTPResponder struct {
}

// With writes a response to the request in the specified context.
func (r *GowebHTTPResponder) With(ctx context.Context, httpStatus int, body []byte) error {

	r.WithStatus(ctx, httpStatus)

	_, writeErr := ctx.HttpResponseWriter().Write(body)
	return writeErr

}

// WithStatus writes the specified HTTP Status Code to the Context's ResponseWriter.
//
// If the Always200ParamName parameter is present, it will ignore the httpStatus argument,
// and always write net/http.StatusOK (200).
func (r *GowebHTTPResponder) WithStatus(ctx context.Context, httpStatus int) error {

	// check for always200
	if len(ctx.FormValue(Always200ParamName)) > 0 {
		// always return OK
		httpStatus = http.StatusOK
	}

	ctx.HttpResponseWriter().WriteHeader(httpStatus)
	return nil
}

// WithStatusText writes the specified HTTP Status Code to the Context's ResponseWriter and
// includes a body with the default status text.
func (r *GowebHTTPResponder) WithStatusText(ctx context.Context, httpStatus int) error {

	writeStatusErr := r.WithStatus(ctx, httpStatus)

	if writeStatusErr != nil {
		return writeStatusErr
	}

	// write the body header
	_, writeErr := ctx.HttpResponseWriter().Write([]byte(http.StatusText(httpStatus)))

	return writeErr
}

// WithOK responds with a 200 OK status code, and no body.
func (r *GowebHTTPResponder) WithOK(ctx context.Context) error {
	return r.WithStatus(ctx, http.StatusOK)
}

// WithRedirect responds with a Found redirection to the specific path or URL.
func (r *GowebHTTPResponder) WithRedirect(ctx context.Context, pathOrURLSegments ...interface{}) error {

	ctx.HttpResponseWriter().Header().Set("Location", paths.PathFromSegments(pathOrURLSegments...))
	return r.WithStatus(ctx, http.StatusFound)

}

// WithTemporaryRedirect responds with a TemporaryRedirect redirection to the specific path or URL.
func (r *GowebHTTPResponder) WithTemporaryRedirect(ctx context.Context, pathOrURLSegments ...interface{}) error {

	ctx.HttpResponseWriter().Header().Set("Location", paths.PathFromSegments(pathOrURLSegments...))
	return r.WithStatus(ctx, http.StatusTemporaryRedirect)

}

// WithPermanentRedirect responds with a redirection to the specific path or URL with the
// http.StatusMovedPermanently status.
func (r *GowebHTTPResponder) WithPermanentRedirect(ctx context.Context, pathOrURLSegments ...interface{}) error {
	ctx.HttpResponseWriter().Header().Set("Location", paths.PathFromSegments(pathOrURLSegments...))
	return r.WithStatus(ctx, http.StatusMovedPermanently)
}
