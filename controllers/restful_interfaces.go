package controllers

import (
	"github.com/stretchrcom/goweb/context"
)

// RestfulController represents an object that satisfies key aspects of a 
// RESTful controller.
type RestfulController interface {
	// Path gets the path prefix for this controller.
	Path() string
}

/*
  RESTful actions
*/

// RestfulCreator represents a controller capable of creating RESTful resources.
type RestfulCreator interface {
	// Create creates a new resource, or new resources.
	Create(ctx context.Context) error
}

// RestfulReader represents a controller capable of reading a single resource.
type RestfulReader interface {
	// Read reads a single resource.
	Read(id string, ctx context.Context) error
}

// RestfulManyReader represents a controller capable of reading many resources.
type RestfulManyReader interface {
	// ReadMany reads many resources.
	ReadMany(ctx context.Context) error
}

// RestfulDeletor represents a controller capable of deleting a single resource.
type RestfulDeletor interface {
	// Delete delets a single resource.
	Delete(id string, ctx context.Context) error
}

// RestfulManyDeleter represents a controller capable of deleting many resources.
type RestfulManyDeleter interface {
	// DeleteMany deletes many resources.
	DeleteMany(ctx context.Context) error
}

// RestfulUpdater represents a controller capable of updating a single resource.
type RestfulUpdater interface {
	// Update updates a single resource.
	Update(id string, ctx context.Context) error
}

// RestfulReplacer represents a controller capable of replacing a single resource.
type RestfulReplacer interface {
	// Replace replaces a single resource.
	Replace(id string, ctx context.Context) error
}

// RestfulManyUpdater represents a controller capable of updating many resources.
type RestfulManyUpdater interface {
	// UpdateMany updates many resources at once.
	UpdateMany(ctx context.Context) error
}

// RestfulOptions represents a controller that explicitly handles an OPTIONS request.
type RestfulOptions interface {
	// Options returns valid options for this resource.
	Options(ctx context.Context) error
}

// RestfulHead represents a controller that explicitly handles a HEAD request.
type RestfulHead interface {
	// Head gets the headers only for the request.
	Head(ctx context.Context) error
}
