package goweb

// Handler method to read an item by the specified ID
type RestReader interface {
	Read(id string, c *Context)
}

// Handler method to read many items
type RestManyReader interface {
	ReadMany(c *Context)
}

// Handler method to update a single item specified by the ID
type RestUpdater interface {
	Update(id string, c *Context)
}

// Handler method to update many items
type RestManyUpdater interface {
	UpdateMany(c *Context)
}

// Handler method to create a new item
type RestCreator interface {
	Create(c *Context)
}

// Handler method to delete an item specified by the ID
type RestDeleter interface {
	Delete(id string, c *Context)
}

// Handler method to delete a collection of items
type RestManyDeleter interface {
	DeleteMany(c *Context)
}

// Interface for RESTful controllers
type RestController interface{}
