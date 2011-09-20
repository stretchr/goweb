package goweb

// Interface for RESTful controllers
type RestController interface {
	
	// Handler method to read an item by the specified ID
	Read(id string, c *Context)
	
	// Handler method to read many items
	ReadMany(c *Context)
	
	// Handler method to update a single item specified by the ID
	Update(id string, c *Context)
	
	// Handler method to update many items
	UpdateMany(c *Context)
	
	// Handler method to create a new item
	Create(c *Context)
	
	// Handler method to delete an item specified by the ID
	Delete(id string, c *Context)
	
	// Handler method to delete a collection of items
	DeleteMany(c *Context)
	
}
