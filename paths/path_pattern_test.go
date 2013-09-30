package paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPathPattern(t *testing.T) {

	path := "/people/{id}/books"
	p, _ := NewPathPattern(path)

	if assert.NotNil(t, p) {
		assert.Equal(t, path, p.RawPath)
	}

}

func TestPathPattern_String(t *testing.T) {

	gp, _ := NewPathPattern("/people/{id}/books/{title}/chapters/{chapter}")
	assert.Equal(t, gp.String(), "{PathPattern:\"/people/{id}/books/{title}/chapters/{chapter}\"}")

}

func TestPathPattern_GetPathMatch_Parameters(t *testing.T) {

	gp, _ := NewPathPattern("/people/{id}/books/{title}/chapters/{chapter}")
	m := gp.GetPathMatch(NewPath("people/123/books/origin-of-species/chapters/2"))

	assert.True(t, m.Matches)
	assert.Equal(t, m.Parameters["id"], "123")
	assert.Equal(t, m.Parameters["title"], "origin-of-species")
	assert.Equal(t, m.Parameters["chapter"], "2")

}

func TestPathPattern_GetPathMatch_Extensions(t *testing.T) {

	gp, _ := NewPathPattern("/people/{id}/books/{title}/chapters/{chapter}")
	m := gp.GetPathMatch(NewPath("people/123/books/origin.of.species/chapters/2.json"))

	assert.True(t, m.Matches)
	assert.Equal(t, m.Parameters["id"], "123")
	assert.Equal(t, m.Parameters["title"], "origin.of.species")
	assert.Equal(t, m.Parameters["chapter"], "2")

}

func TestPathPattern_GetPathMatch_Edges(t *testing.T) {

	// everything
	gp, _ := NewPathPattern(MatchAllPaths)
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("")).Matches)

	// root
	gp, _ = NewPathPattern("/")
	assert.True(t, gp.GetPathMatch(NewPath("/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people")).Matches)

}

func TestPathPattern_GetPathMatch_Matches(t *testing.T) {

	// {variable}
	gp, _ := NewPathPattern("/people/{id}/books")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/nope/123/books")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)

	// ***
	gp, _ = NewPathPattern("/people/{id}/books/***")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/hello/how/do/you/do")).Matches, "/people/123/books/hello/how/do/you/do")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches, "/people/123/books/hello")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches, "/people/123/books")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/")).Matches, "/people/123/books/")

	// *
	gp, _ = NewPathPattern("/people/*/books")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/nope/123/books")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)

	// [optional]
	gp, _ = NewPathPattern("/people/[id]")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)

}

func TestPathPattern_GetPathMatchCatchallPrefixLiteral_Matches(t *testing.T) {
	// ***/literal
	gp, _ := NewPathPattern("/***/books")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/books")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("people/123/[books]/")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("people/123/{books}/")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)
}

func TestPathPattern_GetPathMatchCatchallPrefixSuffix_Matches(t *testing.T) {
	// ***/literal/***
	gp, _ := NewPathPattern("/***/books/***")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/lotr/chapters/one")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/books")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("people/123/[books]/")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("people/123/{books}/")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/novels/lotr/chapters/one")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/novels/hello")).Matches)

}
