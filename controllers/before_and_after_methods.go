package controllers

import (
	"github.com/stretchr/goweb/context"
)

// BeforeHandler represents a controller that has a before handler.
//
// Before handlers will be mapped to any actions of this controller and will
// be called before any of the main methods.
type BeforeHandler interface {

	// Before is called after any other mapped methods.
	Before(context.Context) error
}

// AfterHandler represents a controller that has an after handler.
//
// After handlers will be mapped to any actions of this controller and will
// be called after any of the main methods.
type AfterHandler interface {

	// After is called after any other mapped methods.
	After(context.Context) error
}
