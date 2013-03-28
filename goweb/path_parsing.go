package goweb

import (
	"strings"
)

/*
	Path parsing
*/

// Tests whether a path segment is dynamic or not
func isDynamicSegment(segment string) bool {
	return strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}")
}

func isExtensionSegment(segment string) bool {
	return strings.HasPrefix(segment, ".")
}

// Splits a path into its segments
func getPathSegments(path string) []string {

	ext := getFileExtension(path)

	// trim off the extension (if it's there)
	if len(ext) > 0 {
		path = path[0 : len(path)-(len(ext)+1)]
	}

	// split path
	segments := strings.Split(strings.Trim(path, "/"), "/")

	if len(ext) > 0 {
		segments = append(segments, "."+ext)
	}

	return segments
}

// Generates a parameter value map from the specified parameter key map and path
func getParameterValueMap(keys ParameterKeyMap, path string) ParameterValueMap {

	var paramValues ParameterValueMap = make(ParameterValueMap)

	var segments []string = getPathSegments(path)
	for k, index := range keys {
		paramValues[k] = segments[index]
	}

	return paramValues

}

// Gets the file extension from a path
func getFileExtension(path string) string {

	// get the last segment
	segments := strings.Split(strings.Trim(path, "/"), "/")
	lastSegment := segments[len(segments)-1]
	lastDot := strings.LastIndex(lastSegment, ".")

	if lastDot == -1 {
		return "" /* no extension */
	} else {
		return lastSegment[lastDot+1:]
	}

	return ""

}
