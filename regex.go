package goweb

import (
	"github.com/stretchrcom/goweb/handlers"
)

/*
  This function is just a shortcut so people can do
  goweb.RegexPath instead of having to include the handlers package.
*/

// RegexPath returns a MatcherFunc that mathces the path based on the specified
// Regex pattern.
func RegexPath(regexpPattern string) handlers.MatcherFunc {
	return handlers.RegexPath(regexpPattern)
}
