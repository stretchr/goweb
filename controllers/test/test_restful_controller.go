package test

import (
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/paths"
	"github.com/stretchr/testify/mock"
)

// TestController ia a mocked (using testify/mock) objects that acts like
// a RESTful controller.  It is used internally for testing code, but you can
// use it yourself if you need to in your test code too.
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
func (c *TestController) Replace(id string, ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Options(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Head(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) Before(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
func (c *TestController) After(ctx context.Context) error {
	return c.Called(ctx).Error(0)
}
