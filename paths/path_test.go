package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPath_RawPath(t *testing.T) {

	p := NewPath("people/123/books")

	assert.Equal(t, "people/123/books", p.RawPath)

}

func TestPath_PathFromSegments(t *testing.T) {

	assert.Equal(t, "people/123/books/true", PathFromSegments("people", 123, "books", true))

}

func TestPath_cleanPath(t *testing.T) {

	assert.Equal(t, "people/123/books", cleanPath("/people/123/books/"))
	assert.Equal(t, "people/123/books", cleanPath("//people/123/books/"))
	assert.Equal(t, "people/123/books", cleanPath("//people/123/books////"))

}

func TestPath_Segments(t *testing.T) {

	p := NewPath("people/123/books")
	s := p.Segments()
	assert.Equal(t, "people", s[0])
	assert.Equal(t, "123", s[1])
	assert.Equal(t, "books", s[2])

	p = NewPath("/people/123/books")
	s = p.Segments()
	assert.Equal(t, "people", s[0])
	assert.Equal(t, "123", s[1])
	assert.Equal(t, "books", s[2])

	p = NewPath("/people/123/books/")
	s = p.Segments()
	assert.Equal(t, "people", s[0])
	assert.Equal(t, "123", s[1])
	assert.Equal(t, "books", s[2])

	p = NewPath("/people/123/books.json")
	s = p.Segments()
	assert.Equal(t, "people", s[0])
	assert.Equal(t, "123", s[1])
	assert.Equal(t, "books", s[2])

}

func TestPath_RealFilePath(t *testing.T) {

	var publicPath string = "/static"
	var systemPath string = "/static-files"
	var urlPath string = "/static/hello-world"

	p := NewPath(publicPath)

	assert.Equal(t, "/static-files/hello-world", p.RealFilePath(systemPath, urlPath))

}
