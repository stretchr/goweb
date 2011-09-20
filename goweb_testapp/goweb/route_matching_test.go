package goweb

import (
	"testing"
	"http"
)

func TestRouteMatcherFuncValue(t *testing.T) {
	
	if DontCare != -1 {
		t.Errorf("DontCare should be -1")
	}
	if NoMatch != 0 {
		t.Errorf("NoMatch should be 0")
	}
	if Match != 1 {
		t.Errorf("Match should be 1")
	}
	
}

func TestRouteMatcher_xMethods(t *testing.T) {
	
	var request *http.Request = new(http.Request)
	request.Method = GET_HTTP_METHOD
	var context *Context = new(Context)
	context.Request = request
	
	if GetMethod(context) != Match {
		t.Errorf("GetMethod on a GET context should Match")
	}
	if PutMethod(context) != DontCare {
		t.Errorf("PutMethod on a GET context should DontCare")
	}
	if DeleteMethod(context) != DontCare {
		t.Errorf("DeleteMethod on a GET context should DontCare")
	}
	if PostMethod(context) != DontCare {
		t.Errorf("PostMethod on a GET context should DontCare")
	}
	
	request.Method = POST_HTTP_METHOD
	
	if GetMethod(context) != DontCare {
		t.Errorf("GetMethod on a POST context should DontCare")
	}
	if PutMethod(context) != DontCare {
		t.Errorf("PutMethod on a POST context should DontCare")
	}
	if DeleteMethod(context) != DontCare {
		t.Errorf("DeleteMethod on a POST context should DontCare")
	}
	if PostMethod(context) != Match {
		t.Errorf("PostMethod on a POST context should Match")
	}
	
	request.Method = PUT_HTTP_METHOD
	
	if GetMethod(context) != DontCare {
		t.Errorf("GetMethod on a PUT context should DontCare")
	}
	if PutMethod(context) != Match {
		t.Errorf("PutMethod on a PUT context should Match")
	}
	if DeleteMethod(context) != DontCare {
		t.Errorf("DeleteMethod on a PUT context should DontCare")
	}
	if PostMethod(context) != DontCare {
		t.Errorf("PostMethod on a PUT context should DontCare")
	}
	
	request.Method = DELETE_HTTP_METHOD
	
	if GetMethod(context) != DontCare {
		t.Errorf("GetMethod on a DELETE context should DontCare")
	}
	if PutMethod(context) != DontCare {
		t.Errorf("PutMethod on a DELETE context should Match")
	}
	if DeleteMethod(context) != Match {
		t.Errorf("DeleteMethod on a DELETE context should Match")
	}
	if PostMethod(context) != DontCare {
		t.Errorf("PostMethod on a DELETE context should DontCare")
	}
	
}