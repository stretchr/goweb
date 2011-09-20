package goweb

import "strings"

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
		path = path[0:len(path) - (len(ext) + 1)]
	}
	
	// for this go release...
	segments := strings.Split(strings.Trim(path, "/"), "/", -1)
	
	// for the next Go release...
	//segments := strings.Split(strings.Trim(path, "/"), "/", -1)
	
	if len(ext) > 0 {
		segments = append(segments, "." + ext)
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

// Gets the file extension (in uppercase) from a path
func getFileExtension(path string) string {
	
	lastDot := strings.LastIndex(path, ".")
	
	if lastDot == -1 { 
		return "" /* no extension */ 
	} else {
		return path[lastDot+1:]
	}
	
	return ""
	
}
