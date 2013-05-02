/*
  Goweb v2 BETA - A simple, powerful web framework for Go.

  Routing

  Telling Goweb how to handle requests based on URLs is the most powerful feature of
  Goweb.

    goweb.Map(options)

  The options arguments can be either:

     1 argument: A single func(context.Context)error, which will handle all requests
    2 arguments: A path pattern string and a func(context.Context)error,
                 which will map the func to the path pattern

  Matching is attempted in the order in which they are mapped, so do the most specific mappings
  first and the most generic ones afterwards.

*/
package goweb
