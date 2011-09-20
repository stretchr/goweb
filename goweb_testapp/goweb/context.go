package goweb

import (
	"http"
	"os"
	"strings"
)

// Object holding details about the request and responses
type Context struct {
	
	// The underlying http.Request for this context
	Request *http.Request
	
	// The underlying http.ResponseWriter for this context
	ResponseWriter http.ResponseWriter
	
	// A ParameterValueMap containing path parameters
	PathParams ParameterValueMap
	
	// The format that the response should be
	Format string
	
}

// Helper function to make a new Context object
// with the specified http.Request, http.ResponseWriter and ParameterValueMap
func makeContext(request *http.Request, responseWriter http.ResponseWriter, pathParams ParameterValueMap) *Context {
	
	var context *Context = new(Context)
	
	// set the parameters
	context.Request = request
	context.ResponseWriter = responseWriter
	context.PathParams = pathParams
	
	// note the format
	context.Format = getFormatForRequest(request)
	
	return context
}

/*
	Form and parameter parsing
*/

// Gets the context value from the request
func (c *Context) GetRequestContext() string {
	request := c.Request
	if request.Form == nil {
		request.ParseForm()
	}
	return request.Form.Get(REQUEST_CONTEXT_PARAMETER)
}

// Gets the callback value from the request
func (c *Context) GetCallback() string {
	request := c.Request
	if request.Form == nil {
		request.ParseForm()
	}
	return request.Form.Get(REQUEST_CALLBACK_PARAMETER)
}

/*
	HTTP Method helper functions
*/

// Checks whether the HTTP method is GET or not
func (c *Context) IsGet() bool {
	return c.Request.Method == GET_HTTP_METHOD
}

// Checks whether the HTTP method is POST or not
func (c *Context) IsPost() bool {
	return c.Request.Method == POST_HTTP_METHOD
}

// Checks whether the HTTP method is PUT or not
func (c *Context) IsPut() bool {
	return c.Request.Method == PUT_HTTP_METHOD
}

// Checks whether the HTTP method is DELETE or not
func (c *Context) IsDelete() bool {
	return c.Request.Method == DELETE_HTTP_METHOD
}


/*
	RespondWith* methods
*/
func (c *Context) Respond(data interface{}, statusCode int, errors []string, context *Context) os.Error {

	// make the standard response object
	obj := makeStandardResponse()
	obj.E = errors
	obj.D = data
	obj.S = statusCode
	obj.C = c.GetRequestContext()

	return c.WriteResponse(obj, statusCode)

}

// Writes the specified object out (with the specified status code)
// using the appropriate formatter
func (c *Context) WriteResponse(obj interface{}, statusCode int) os.Error {
	
	var error os.Error

	// get the formatter
	formatter, error := GetFormatter(c.Format)
	
	if error != nil {
		c.writeInternalServerError(error, http.StatusNotFound)
		return error
	} else {
	
		// set the content type
		c.ResponseWriter.Header()["Content-Type"] = []string{ formatter.ContentType() }
	
		// format the output
		output, error := formatter.Format(obj)
		
		if error != nil {
			c.writeInternalServerError(error, http.StatusInternalServerError)
			return error
		} else {
			
			outputString := string(output)
			
			/*
				JSONP
			*/
			callback := c.GetCallback()
			if callback != "" {
				
				// wrap with function call
				
				requestContext := c.GetRequestContext()
				
				outputString = callback + "(" + outputString
				
				if requestContext != "" {
					outputString = outputString + ", \"" + requestContext + "\")"
				} else {
					outputString = outputString + ")"
				}
				
				// set the new content type
				c.ResponseWriter.Header()["Content-Type"] = []string{ JSONP_CONTENT_TYPE }
				
			}
			
			// write the status code
			if (strings.Index(c.Request.URL.Raw, REQUEST_ALWAYS200_PARAMETER) > -1) {
			
				// "always200"
				// write a fake 200 status code (regardless of what the actual code was)
				c.ResponseWriter.WriteHeader(http.StatusOK)
			
			} else {
				
				// write the actual status code
				c.ResponseWriter.WriteHeader(statusCode)
				
			}
			
			// write the output
			c.ResponseWriter.Write([]uint8(outputString))
		
		}
	
	}
	
	// success - no errors
	return nil
	
}

func (c *Context) writeInternalServerError(error os.Error, statusCode int) {
	http.Error(c.ResponseWriter, error.String(), statusCode)
}

// Responds with the specified HTTP status code defined in RFC 2616
// see http://golang.org/src/pkg/http/status.go for options
func (c *Context) RespondWithStatus(statusCode int) os.Error {
	return c.Respond(nil, statusCode, nil, c)
}

// Responds with the specified HTTP status code defined in RFC 2616
// and adds the description to the errors list
// see http://golang.org/src/pkg/http/status.go for options
func (c *Context) RespondWithError(statusCode int) os.Error {
	return c.RespondWithErrorMessage(http.StatusText(statusCode), statusCode)
}

func (c *Context) RespondWithErrorMessage(message string, statusCode int) os.Error {
	return c.Respond(nil, statusCode, []string{ message }, c)
}

// Responds with the specified data
func (c *Context) RespondWithData(data interface{}) os.Error {
	return c.Respond(data, http.StatusOK, nil, c)
}

// Responds with OK status (200) and no data
func (c *Context) RespondWithOK() os.Error {
	return c.RespondWithData(nil)
}

// Responds with 404 Not Found
func (c *Context) RespondWithNotFound() os.Error {
	return c.RespondWithError(http.StatusNotFound)
}

// Responds with 501 Not Implemented
func (c *Context) RespondWithNotImplemented() os.Error {
	return c.RespondWithError(http.StatusNotImplemented)
}