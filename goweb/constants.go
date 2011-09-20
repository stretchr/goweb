package goweb

/*
	Constants
*/

// Regex placeholder pattern
var ROUTE_REGEX_PLACEHOLDER string  = "(.*)"

// HTTP Methods
const GET_HTTP_METHOD string = "GET"
const POST_HTTP_METHOD string = "POST"
const PUT_HTTP_METHOD string = "PUT"
const DELETE_HTTP_METHOD string = "DELETE"

/*
	Error messages
*/
const ERR_STANDARD_PREFIX string = "Oops, something went wrong: "
const ERR_NO_CONTROLLER string = "Routes must have a valid Controller"
const ERR_NO_MATCHING_ROUTE string = "No route found for that path"

/*
	API Constants
*/

// Parameter name for Request Context
const REQUEST_CONTEXT_PARAMETER string = "context"

// Parameter name for callback
const REQUEST_CALLBACK_PARAMETER string = "callback"

// Parameter name for Always200
const REQUEST_ALWAYS200_PARAMETER string = "always200"

// Parameter name for method override
const REQUEST_METHOD_OVERRIDE_PARAMETER string = "method"

// JSONP Content-Type header value
const JSONP_CONTENT_TYPE string = "text/javascript"