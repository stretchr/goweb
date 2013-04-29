package paths

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

/*
   GoWeb paths

   /literal
   /{placeholder}
   /[optional placeholder]
   /* - matches everything following this
   /people/{id:int} - specific types
   /people/{id:string}

*/

func assertRegexPattern(t *testing.T, gowebPath, expectedRegex string) bool {
	p, _ := NewGowebPath(gowebPath)
	return assert.Equal(t, expectedRegex, p.RegexPattern())
}

func TestNewGowebPath(t *testing.T) {

	path := "/people/{id}/books"
	p, _ := NewGowebPath(path)

	if assert.NotNil(t, p) {
		assert.Equal(t, path, p.RawPath)
	}

}

func TestGowebPath_GetPathMatch_(t *testing.T) {

	// {variable}
	gp, _ := NewGowebPath("/people/{id}/books")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/nope/123/books")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)

	// ...
	gp, _ = NewGowebPath("/people/{id}/books/...")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books/hello/how/do/you/do")).Matches)

	// *
	gp, _ = NewGowebPath("/people/*/books")

	assert.True(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/PEOPLE/123/BOOKS")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/People/123/Books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("people/123/books/")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/nope/123/books")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books/hello")).Matches)

	// [optional]
	gp, _ = NewGowebPath("/people/[id]")
	assert.True(t, gp.GetPathMatch(NewPath("/people/123")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people/")).Matches)
	assert.True(t, gp.GetPathMatch(NewPath("/people")).Matches)

	assert.False(t, gp.GetPathMatch(NewPath("/people/123/books")).Matches)

}
