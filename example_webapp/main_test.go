package main

import (
	"github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/handlers"
	"github.com/stretchr/testify/assert"
	testifyhttp "github.com/stretchr/testify/http"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {

	// make a test HttpHandler and use it
	codecService := new(services.WebCodecService)
	handler := handlers.NewHttpHandler(codecService)
	goweb.SetDefaultHttpHandler(handler)

	// call the target code
	mapRoutes()

	goweb.Test(t, "GET people/me", func(t *testing.T, response *testifyhttp.TestResponseWriter) {

		// should be a redirect
		assert.Equal(t, http.StatusFound, response.StatusCode, "Status code should be correct")

	})

}
