package goweb

import (
	"github.com/stretchrcom/goweb/handlers"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestSetDefaultHttpHandler(t *testing.T) {

	handler := new(handlers.HttpHandler)

	if assert.NotEqual(t, handler, defaultHttpHandler) {

		SetDefaultHttpHandler(handler)

		assert.Equal(t, handler, defaultHttpHandler)

	}

}
