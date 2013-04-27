package paths

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPath_RawPath(t *testing.T) {

	p := NewPath("people/123/books")

	assert.Equal(t, "people/123/books", p.RawPath)

}

func TestPath_Segments(t *testing.T) {

	p := NewPath("people/123/books")

	s, _ := p.Segments()

	assert.Equal(t, "people", s[0])
	assert.Equal(t, "123", s[1])
	assert.Equal(t, "books", s[2])

}
