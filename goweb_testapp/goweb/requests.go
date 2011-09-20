package goweb

import (
	"http"
	"strings"
)

// Constant string for HTML format
const HTML_FORMAT string = "HTML"

// Constant string for XML format
const XML_FORMAT string = "XML"

// Constant string for JSON format
const JSON_FORMAT string = "JSON"

// The fallback format if one cannot be determined by the request
const DEFAULT_FORMAT string = JSON_FORMAT

// Gets a string describing the format of the request.
func getFormatForRequest(request *http.Request) string {
	
	if (request.URL == nil) { return DEFAULT_FORMAT }
	
	// use the file extension as the format
	ext := strings.ToUpper(getFileExtension(request.URL.Path))
	if ext != "" {
		
		// manual overrides
		if ext == "HTM" { return HTML_FORMAT }
		
		return ext
	}
	
	// we don't know, so use the default one
	return DEFAULT_FORMAT
	
}