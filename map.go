package goweb

// Map adds a new mapping to the DefaultHttpHandler.
// 
// The Map function has many flavours.
//
// The path pattern, and the handler function, followed by any optional matcher funcs will cause
// the func to be executed when the path pattern matches.
//
//     (pathPattern, func [, matcherFuncs])
//
// The HTTP method, the path pettern, the handler function, followed by optional matcher funcs will cause
// the func to be executed when the path and HTTP method match.
//
//     (method, pathPattern, func [, matcherFuncs])
//
// Just the handler function, and any optional matcher funcs will add a catch-all handler.  This should
// be the last call to Map that you make.
//
//     (func [, matcherFuncs])
//
// Examples
//
// The following code snippets are real examples of how to use the Map function:
//
//     handler.Map(http.MethodPost, "/events", func(c context.Context) error {
//
//       // TODO: Add an event
//
//       // no errors
//       return nil
//
//     })
//
//     handler.Map("/articles/{year}/{month}/{day}", func(c context.Context) error {
//  
//       day := c.PathParams().Get("day")
//       month := c.PathParams().Get("month")
//       year := c.PathParams().Get("year")
//
//       // show the articles for the specified day
//
//       // no errors
//       return nil
//
//     })
//
//     handler.Map(func(c context.Context) error {
//     
//       // everything else is a 404
//       goweb.Respond.WithStatus(c, http.StatusNotFound)
//  
//       // no errors
//       return nil
//
//     })
func Map(options ...interface{}) error {
	return DefaultHttpHandler().Map(options...)
}

// MapController maps a controller in the handler.
//
// A controller is any object that inherits from one or more of the controllers.Restful*
// interfaces.
//
// They include:
//
//     RestfulController.Path(context.Context) string
//     RestfulCreator.Create(context.Context) error
//     RestfulReader.Read(context.Context) error
//     RestfulManyReader.ReadMany(context.Context) error
//     RestfulDeletor.Delete(id string, context.Context) error
//     RestfulManyDeleter.DeleteMany(context.Context) error
//     RestfulUpdater.Update(id string, context.Context) error
//     RestfulReplacer.Replace(id string, context.Context) error
//     RestfulManyUpdater.UpdateMany(context.Context) error
//     RestfulOptions.Options(context.Context) error
//     RestfulHead.Head(context.Context) error
//
// This code will map the controller and use the Path() method on the
// controller to determine the path prefix:
//
//     MapController(controller)
//
// This code will map the controller to the specified path prefix regardless
// of what the Path() methods returns:
//
//     MapController(path, controller)
func MapController(options ...interface{}) error {
	return DefaultHttpHandler().MapController(options...)
}

/*
  DEVNOTE: These functions are not tested because it simply passes the call on to the
  DefaultHttpHandler.
*/
