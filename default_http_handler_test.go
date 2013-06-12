package goweb

import (
	"github.com/stretchr/goweb/handlers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetDefaultHttpHandler(t *testing.T) {

	handler := new(handlers.HttpHandler)

	if assert.NotEqual(t, handler, defaultHttpHandler) {

		SetDefaultHttpHandler(handler)

		assert.Equal(t, handler, defaultHttpHandler)

	}

}
