package test

import (
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/testify/mock"
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
