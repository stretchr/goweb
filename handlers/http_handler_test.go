package handlers

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestNewHttpHandler(t *testing.T) {

	h := NewHttpHandler()

	if assert.NotNil(t, h) {
		if assert.NotNil(t, h.handlers) {
			if assert.Equal(t, 3, len(h.handlers)) {

				assert.NotNil(t, h.handlers[0])
				assert.NotNil(t, h.handlers[1])
				assert.NotNil(t, h.handlers[2])

			}
		}
	}

}

func TestHandlers(t *testing.T) {

	h := NewHttpHandler()

	assert.Equal(t, h.handlers[0], h.BeforeHandlers())
	assert.Equal(t, h.handlers[1], h.Handlers())
	assert.Equal(t, h.handlers[2], h.AfterHandlers())

}
