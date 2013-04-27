package context_test

import (
	"github.com/stretchrcom/goweb/context"
)

func MakeTestContext() *context.Context {
	return MakeTestContextWithPath("/")
}

func MakeTestContextWithPath(path string) *context.Context {
	return new(context.Context)
}
