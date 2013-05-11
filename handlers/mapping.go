package handlers

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
)

/*
  Map maps an executor to the specified PathPattern.
*/
func (h *HttpHandler) Map(options ...interface{}) error {

	var pathPattern string
	var executor HandlerExecutionFunc

	switch len(options) {
	case 0: // ()
		// no arguments is an error
		panic("goweb: Cannot call Map with no arguments")
		break
	case 1: // (func)
		// catch all assumption
		pathPattern = "***"
		executor = options[0].(func(context.Context) error)
		break
	case 2: // (string, func)
		// pattern and executor
		pathPattern = options[0].(string)
		executor = options[1].(func(context.Context) error)
		break
	}

	path, pathErr := paths.NewPathPattern(pathPattern)

	if pathErr != nil {
		return pathErr
	}

	handler := NewPathMatchHandler(path, executor)

	// check match funcs too
	if len(options) > 2 { // (string, func, matcherFuncs...)
		// pattern and executor
		pathPattern = options[0].(string)
		executor = options[1].(func(context.Context) error)

		// collect the matcher funcs too
		var ok bool
		matcherFuncs := make([]MatcherFunc, len(options)-2)
		for i := 2; i < len(options); i++ {
			if matcherFuncs[i-2], ok = options[i].(MatcherFunc); !ok {
				panic("goweb: [2...] arguments passed to Map must be of type MatcherFunc.")
			}
		}

		// set them
		handler.MatcherFuncs = matcherFuncs

	}

	h.AppendHandler(handler)

	// ok
	return nil

}
