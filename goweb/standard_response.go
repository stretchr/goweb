package goweb

import "http"

// The standard API response object
type standardResponse struct {
	
	// The context of the request that initiated this response
	C string
	
	// The HTTP Status code of this response
	S int
	
	// The data (if any) for this response
	D interface{}
	
	// A list of any errors that occurred while processing
	// the response
	E []string

}

// Makes a standardResponse object with the specified settings
func makeStandardResponse() *standardResponse {
	response := new(standardResponse)
	response.C = ""
	response.S = 200
	response.E = nil
	return response
}

// Makes a successful standardResponse object with the specified settings
func makeSuccessfulStandardResponse(context string, statusCode int, data interface{}) *standardResponse {
	response := makeStandardResponse()
	response.C = context
	response.S = statusCode
	response.D = data
	return response
}

// Makes an unsuccessful standardResponse object with the specified settings
func makeFailureStandardResponse(context string, statusCode int) *standardResponse {
	response := makeStandardResponse()
	response.C = context
	response.S = statusCode
	response.E = []string{ http.StatusText(statusCode) }
	return response
}