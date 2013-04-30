package goweb

import (
	"github.com/stretchrcom/goweb/context"
)

func Map(pathPattern string, executor func(context.Context) error) error {
	return DefaultHttpHandler().Map(pathPattern, executor)
}
