package test

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/testify/mock"
)

type TestSemiRestfulController struct {
	mock.Mock
}

func (c *TestSemiRestfulController) Create(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestSemiRestfulController) Read(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestSemiRestfulController) ReadMany(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
