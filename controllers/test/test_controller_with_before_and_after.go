package test

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/testify/mock"
)

type TestHandlerWithBeforeAndAfters struct {
	mock.Mock
}

func (c *TestHandlerWithBeforeAndAfters) Path() string {
	return "test"
}

func (c *TestHandlerWithBeforeAndAfters) Before(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

func (c *TestHandlerWithBeforeAndAfters) After(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

func (c *TestHandlerWithBeforeAndAfters) Create(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}

func (c *TestHandlerWithBeforeAndAfters) Replace(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
