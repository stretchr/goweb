package goweb

import (
	"github.com/stretchrcom/goweb/context"
)

/*
  Map maps an executor to the specified PathPattern on the DefaultHttpHandler.
*/
func Map(pathPattern string, executor func(context.Context) error) error {
	return DefaultHttpHandler().Map(pathPattern, executor)
}
