package goweb

import (
	"github.com/stretchr/goweb/paths"
	"net/http"
)

const (
	// headerKeyLocation is the header key for Location.
	headerKeyLocation string = "Location"
)

// Status writes the header with the specified status.
func Status(response http.ResponseWriter, status int) {
	response.WriteHeader(status)
}

// RedirectTo writes the Location header.
func RedirectTo(response http.ResponseWriter, pathOrURLSegments ...interface{}) {
	response.Header().Set(headerKeyLocation, paths.PathFromSegments(pathOrURLSegments...))
}

// redirectWithStatus writes the Location header and sets the specified
// status in the response.
func redirectWithStatus(response http.ResponseWriter, status int, pathOrURLSegments ...interface{}) {
	RedirectTo(response, pathOrURLSegments...)
	Status(response, status)
}

// Redirect writes the Location header and sets the http.StatusFound
// response.
func Redirect(response http.ResponseWriter, pathOrURLSegments ...interface{}) {
	redirectWithStatus(response, http.StatusFound, pathOrURLSegments...)
}

// RedirectTemp writes the Location header and sets the http.StatusTemporaryRedirect
// response.
func RedirectTemp(response http.ResponseWriter, pathOrURLSegments ...interface{}) {
	redirectWithStatus(response, http.StatusTemporaryRedirect, pathOrURLSegments...)
}

// RedirectPerm writes the Location header and sets the http.StatusMovedPermanently
// response.
func RedirectPerm(response http.ResponseWriter, pathOrURLSegments ...interface{}) {
	redirectWithStatus(response, http.StatusMovedPermanently, pathOrURLSegments...)
}
