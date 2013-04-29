package context

import (
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/stew/objects"
)

/**
  Context represents a single request and provides methods for responding.
*/
type Context struct {
	path *paths.Path
	data objects.Map
}

func NewContext(path string) *Context {

	c := new(Context)

	c.data = make(objects.Map)
	c.path = paths.NewPath(path)

	return c

}

func (c *Context) Path() *paths.Path {
	return c.path
}

func (c *Context) Data() objects.Map {
	return c.data
}
