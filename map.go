package goweb

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/handlers"
	"github.com/stretchrcom/goweb/paths"
)

func Map(pathPattern string, executor func(context.Context) error) error {

	path, pathErr := paths.NewPathPattern(pathPattern)

	if pathErr != nil {
		return pathErr
	}

	handler := &handlers.PathMatchHandler{path, executor}
	DefaultHttpHandler().AppendHandler(handler)

	// ok
	return nil

}
