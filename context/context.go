package context

import (
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
	"net/http"
)

type Context interface {
	Path() *paths.Path
	Data() objects.Map
	HttpResponseWriter() http.ResponseWriter
	HttpRequest() *http.Request
}
