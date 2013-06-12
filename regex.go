package goweb

import (
	"github.com/stretchr/goweb/handlers"
)

/*
  This function is just a shortcut so people can do
  goweb.RegexPath instead of having to include the handlers package.
*/

// RegexPath returns a MatcherFunc that mathces the path based on the specified
// Regex pattern.
//
// To match a path that contains only numbers, you could do:
//
//     goweb.Map(executionFunc, goweb.RegexPath(`^[0-9]+$`))
func RegexPath(regexpPattern string) handlers.MatcherFunc {
	return handlers.RegexPath(regexpPattern)
}
