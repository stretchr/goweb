package goweb

// interfaces for several levels of nested controllers
// currently arbitarily limited to 3 deep

// Handler method to read an item by the specified ID
type RestReader interface { Read(id string, c *Context) }
type NestedRestReader interface { Read(parentId string, id string, c *Context) }

// Handler method to read many items
type RestManyReader interface { ReadMany(c *Context) }
type NestedRestManyReader interface { ReadMany(parentId string, c *Context) }

// Handler method to update a single item specified by the ID
type RestUpdater interface { Update(id string, c *Context) }
type NestedRestUpdater interface { Update(parentId string, id string, c *Context) }

// Handler method to update many items
type RestManyUpdater interface { UpdateMany(c *Context) }
type NestedRestManyUpdater interface { UpdateMany(parentId string, c *Context) }

// Handler method to create a new item
type RestCreator interface { Create(c *Context) }
type NestedRestCreator interface { Create(parentId string, c *Context) }

// Handler method to delete an item specified by the ID
type RestDeleter interface { Delete(id string, c *Context) }
type NestedRestDeleter interface { Delete(parentId string, id string, c *Context) }

// Handler method to delete a collection of items
type RestManyDeleter interface { DeleteMany(c *Context) }
type NestedRestManyDeleter interface { DeleteMany(parentId string, c *Context) }

// Interface for RESTful controllers
type RestController interface { }

