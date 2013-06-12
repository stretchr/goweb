package handlers_test

import (
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/testify/mock"
)

type TestHandler struct {
	mock.Mock
}

func (h *TestHandler) WillHandle(c context.Context) (bool, error) {

	args := h.Called(c)

	if args[1] == nil {
		return args[0].(bool), nil
	}

	return args[0].(bool), args[1].(error)

}

func (h *TestHandler) Handle(c context.Context) (bool, error) {
	args := h.Called(c)
	if args.Error(1) == nil {
		return args.Bool(0), nil
	}
	return args.Bool(0), args.Error(1)
}
