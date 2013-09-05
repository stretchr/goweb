package goweb

import (
	"github.com/stretchr/goweb/handlers"
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
// Each matcherFunc argument can be one of three types:
//     1) handlers.MatcherFunc
//     2) []handlers.MatcherFunc
//     3) func(context.Context) (MatcherFuncDecision, error)
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
// of what the Path() method returns:
//
//     MapController(path, controller)
//
// Optionally, you can pass matcherFuncs as optional additional arguments.  See
// goweb.Map() for details on the types of arguments allowed.
func MapController(options ...interface{}) error {
	return DefaultHttpHandler().MapController(options...)
}

// MapStatic maps static files from the specified systemPath to the
// specified publicPath.
//
//     goweb.MapStatic("/static", "location/on/system/to/files")
//
// Goweb will automatically expand the above public path pattern from `/static` to
// `/static/***` to ensure subfolders are automatcially mapped.
//
// Paths
//
// The systemPath is relative to where you will actually run your app from, for
// example if I am doing `go run main.go` inside the `example_webapp` folder, then
// `./` would refer to the `example_webapp` folder itself.  Therefore, to map to the
// `static-files` subfolder, you just need to specify the name of the directory:
//
//     goweb.MapStatic("/static", "static-files")
//
// In some cases, your systemPath might be different for development and production
// and this is something to watch out for.
//
// Optionally, you can pass arguments of type MatcherFunc after the second
// argument.  Unlike goweb.Map, these can only be of type MatcherFunc.
func MapStatic(publicPath, systemPath string, matcherFuncs ...handlers.MatcherFunc) (handlers.Handler, error) {
	return DefaultHttpHandler().MapStatic(publicPath, systemPath, matcherFuncs...)
}

// MapStaticFile maps a static file from the specified staticFilePath to the
// specified publicPath.
//
//     goweb.MapStaticFile("favicon.ico", "/location/on/system/to/icons/favicon.ico")
//
// Only paths matching exactly publicPath will cause the single file specified in
// staticFilePath to be delivered to clients.
//
// Optionally, you can pass arguments of type MatcherFunc after the second
// argument.  Unlike goweb.Map, these can only be of type MatcherFunc.
func MapStaticFile(publicPath, staticFilePath string, matcherFuncs ...handlers.MatcherFunc) (handlers.Handler, error) {
	return DefaultHttpHandler().MapStaticFile(publicPath, staticFilePath, matcherFuncs...)
}

/*
  DEVNOTE: These functions are not tested because it simply passes the call on to the
  DefaultHttpHandler.
*/
