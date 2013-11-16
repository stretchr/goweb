package controllers

import (
	"github.com/stretchr/goweb/context"
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
//
// This is usually mapped to the following kind of request:
//     POST /resources
type RestfulCreator interface {
	// Create creates a new resource, or new resources.
	Create(ctx context.Context) error
}

// RestfulReader represents a controller capable of reading a single resource.
//
// This is usually mapped to the following kind of request:
//     GET /resources/{id}
type RestfulReader interface {
	// Read reads a single resource.
	Read(id string, ctx context.Context) error
}

// RestfulManyReader represents a controller capable of reading many resources.
//
// This is usually mapped to the following kind of request:
//     GET /resources
type RestfulManyReader interface {
	// ReadMany reads many resources.
	ReadMany(ctx context.Context) error
}

// RestfulDeletor represents a controller capable of deleting a single resource.
//
// This is usually mapped to the following kind of request:
//     DELETE /resources/{id}
type RestfulDeletor interface {
	// Delete delets a single resource.
	Delete(id string, ctx context.Context) error
}

// RestfulManyDeleter represents a controller capable of deleting many resources.
//
// This is usually mapped to the following kind of request:
//     DELETE /resources
type RestfulManyDeleter interface {
	// DeleteMany deletes many resources.
	DeleteMany(ctx context.Context) error
}

// RestfulUpdater represents a controller capable of updating a single resource.
//
// This is usually mapped to the following kind of request:
//     PATCH /resources/{id}
type RestfulUpdater interface {
	// Update updates a single resource.
	Update(id string, ctx context.Context) error
}

// RestfulReplacer represents a controller capable of replacing a single resource.
//
// This is usually mapped to the following kind of request:
//     PUT /resources/{id}
type RestfulReplacer interface {
	// Replace replaces a single resource.
	Replace(id string, ctx context.Context) error
}

// RestfulManyUpdater represents a controller capable of updating many resources.
//
// This is usually mapped to the following kind of request:
//     PATCH /resources
type RestfulManyUpdater interface {
	// UpdateMany updates many resources at once.
	UpdateMany(ctx context.Context) error
}

// RestfulOptions represents a controller that explicitly handles an OPTIONS request.
//
// This is usually mapped to the following kind of requests:
//     OPTIONS /resources
//     OPTIONS /resources/{id}
type RestfulOptions interface {
	// Options returns valid options for this resource.
	Options(ctx context.Context) error
}

// RestfulHead represents a controller that explicitly handles a HEAD request.
//
// This is usually mapped to the following kind of requests:
//     HEAD /resources
//     HEAD /resources/{id}
type RestfulHead interface {
	// Head gets the headers only for the request.
	Head(ctx context.Context) error
}
