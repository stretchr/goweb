package context

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestNewContext(t *testing.T) {

	c := NewContext("/people/123")

	if assert.NotNil(t, c) {

		assert.NotNil(t, c.data)
		assert.Equal(t, "people/123", c.Path().RawPath)

	}

}
