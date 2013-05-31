// The handlers package contains tools for building pipelines of handlers which
// requests are run through in order to properly respond to requests.
//
// While the main goweb package wraps most of this functionality, the handlers
// package allows for more advanced use cases, such as:
//
//   * When you want to write your own Handler implementations
//   * When you want multiple HttpHandlers within the same project
//   * When you want to build your own MatcherFuncs to enable you to be
//     more selective about when to handler specific requests
package handlers
