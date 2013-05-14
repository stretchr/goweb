package test

import (
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/paths"
	"github.com/stretchrcom/testify/mock"
)

type TestController struct {
	mock.Mock
}

func (c *TestController) Path() string {
	return paths.PathPrefixForClass(c)
}

func (c *TestController) Create(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Read(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) ReadMany(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Delete(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) DeleteMany(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Update(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) UpdateMany(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Options(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Head(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
