package api

import (
	"github.com/stretchrcom/goweb/context"
)

type GowebAPIResponder struct{}

func (a *GowebAPIResponder) Respond(ctx *context.Context, status int, data interface{}, errors []string) {

}
