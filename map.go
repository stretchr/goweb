package goweb

import (
	"github.com/stretchrcom/goweb/handlers"
)

// Map adds a new mapping to the DefaultHttpHandler.
//
// The Map function has many flavours.
//
// 1. Implementing and passing in your own handlers.Handler object will just map the handler
// directly.
//
// 2. The path pattern, and the handler function, followed by any optional matcher funcs will cause
// the func to be executed when the path pattern matches.
//
//     (pathPattern, func [, matcherFuncs])
//
// 3. The HTTP method (or an array of HTTP methods), the path pettern, the handler function, followed by
// optional matcher funcs will cause the func to be executed when the path and HTTP method match.
//
//     (method|methods, pathPattern, func [, matcherFuncs])
//
// 4. Just the handler function, and any optional matcher funcs will add a catch-all handler.  This should
// be the last call to Map that you make.
//
//     (func [, matcherFuncs])
//
// Examples
//
// The following code snippets are real examples of how to use the Map function:
//
//     // POST /events
//     handler.Map(http.MethodPost, "/events", func(c context.Context) error {
//
//       // TODO: Add an event
//
//       // no errors
//       return nil
//
//     })
//
//     // POST|PUT /events
//     handler.Map([]string{http.MethodPost, http.MethodPut}, "/events", func(c context.Context) error {
//
//       // TODO: Add an event
//
//       // no errors
//       return nil
//
//     })
//
//     // POST|PUT|DELETE|GET /articles/2013/05/01
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
//     // All requests
//     handler.Map(func(c context.Context) error {
//
//       // everything else is a 404
//       goweb.Respond.WithStatus(c, http.StatusNotFound)
//
//       // no errors
//       return nil
//
//     })
//
// For a full overview of valid paths, see the "Mapping paths" section above.
func Map(options ...interface{}) (handlers.Handler, error) {
	return DefaultHttpHandler().Map(options...)
}

// MapBefore adds a new mapping to the DefaultHttpHandler that
// will be executed before other handlers.
//
// The usage is the same as the goweb.Map function.
//
// Before handlers are called before any of the normal handlers,
// and before any processing has begun.  Setting headers is appropriate
// for before handlers, but be careful not to actually write anything or
// Goweb will likely end up trying to write the headers twice and headers set
// in the processing handlers will have no effect.
func MapBefore(options ...interface{}) (handlers.Handler, error) {
	return DefaultHttpHandler().MapBefore(options...)
}

// MapAfter adds a new mapping to the DefaultHttpHandler that
// will be executed after other handlers.
//
// The usage is the same as the goweb.Map function.
//
// After handlers are called after the normal processing handlers are
// finished, and usually after the response has been written.  Setting headers
// or writing additional bytes will have no effect in after handlers.
func MapAfter(options ...interface{}) (handlers.Handler, error) {
	return DefaultHttpHandler().MapAfter(options...)
}

// MapController maps a controller in the handler.
//
// A controller is any object that implements one or more of the controllers.Restful*
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
// Optionally, you can map Before and After methods too, to allow controller specific
// code to run at appropriate times:
//
//     BeforeHandler.Before(context.Context) error
//     AfterHandler.After(context.Context) error
//
// To implement any of these methods, you just need to provide a method with the
// same name and signature.  For example, a simple RESTful controller that just
// provides a simple GET might look like this:
//
//     type MyController struct{}
//
//     func (c *MyController) ReadMany(ctx context.Context) error {
//       // TODO: do something
//     }
//
// The above code would map only `GET /my`.
//
// Controller Paths
//
// This code will map the controller and use the Path() method on the
// controller to determine the path prefix (or it will try to guess the URL part
// based on the name of the struct):
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

// MapStatic maps static files from the specified systemPath to the
// specified publicPath.
//
//     goweb.MapStatic("/static", "/location/on/system/to/files")
//
// Goweb will automatically expand the above public path pattern from `/static` to
// `/static/***` to ensure subfolders are automatcially mapped.
func MapStatic(publicPath, systemPath string) (handlers.Handler, error) {
	return DefaultHttpHandler().MapStatic(publicPath, systemPath)
}

/*
  DEVNOTE: These functions are not tested because it simply passes the call on to the
  DefaultHttpHandler.
*/
