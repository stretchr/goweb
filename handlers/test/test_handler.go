package handlers_test

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/testify/mock"
)

type TestHandler struct {
	mock.Mock
}

func (h *TestHandler) WillHandle(c *context.Context) (bool, error) {

	args := h.Called(c)

	if args[1] == nil {
		return args[0].(bool), nil
	}

	return args[0].(bool), args[1].(error)

}

func (h *TestHandler) Handle(c *context.Context) error {
	e := h.Called(c)[0]
	if e == nil {
		return nil
	}
	return e.(error)
}
