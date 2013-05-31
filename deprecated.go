package goweb

import (
	"fmt"
)

// deprecatedPanic panics with a deprecation warning and tip of how to resolve
// the issue.
func deprecatedPanic(method, alternativeTip string) {
	panic(fmt.Sprintf("goweb: (deprecated) %s is no longer supported. %s", method, alternativeTip))
}

// MapFunc is no longer supported in Goweb 2.
//
// Use goweb.Map instead.
func MapFunc(path string, controllerFunc interface{}, matcherFuncs ...interface{}) interface{} {
	deprecatedPanic("MapFunc", "Use goweb.Map instead.")
	return nil
}

// MapRest is no longer supported in Goweb 2.
//
// Use goweb.MapController instead.
func MapRest(pathPrefix string, controller interface{}) {
	deprecatedPanic("MapRest", "Use goweb.MapController instead.")
}
