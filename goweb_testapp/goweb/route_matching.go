package goweb

// Represents the return value for RouteMatcher functions
type RouteMatcherFuncValue int

// Functions used to decide whether a route matches a request or not
type RouteMatcherFunc func(c *Context) RouteMatcherFuncValue

// Indicates that the RouteMatcherFunc doesn't care whether the
// route is a Match or NoMatch
const DontCare RouteMatcherFuncValue = -1

// Indicates that the route should NOT match 
const NoMatch RouteMatcherFuncValue = 0

// Indicates that the route should match
const Match RouteMatcherFuncValue = 1


// Returns Match if the Method of the http.Request in the specified
// Context is GET, otherwise returns DontCare
func GetMethod(c *Context) RouteMatcherFuncValue {
	if c.IsGet() {
		return Match
	}
	return DontCare
}

// Returns Match if the Method of the http.Request in the specified
// Context is PUT, otherwise returns DontCare
func PutMethod(c *Context) RouteMatcherFuncValue {
	if c.IsPut() {
		return Match
	}
	return DontCare
}

// Returns Match if the Method of the http.Request in the specified
// Context is DELETE, otherwise returns DontCare
func DeleteMethod(c *Context) RouteMatcherFuncValue {
	if c.IsDelete() {
		return Match
	}
	return DontCare
}

// Returns Match if the Method of the http.Request in the specified
// Context is POST, otherwise returns DontCare
func PostMethod(c *Context) RouteMatcherFuncValue {
	if c.IsPost() {
		return Match
	}
	return DontCare
}