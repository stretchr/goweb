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

	handler := &PathMatchHandler{path, executor}
	h.AppendHandler(handler)

	// ok
	return nil

}
