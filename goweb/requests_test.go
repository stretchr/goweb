package goweb

import (
	"testing"
	"http"
)

func TestFormatStrings(t *testing.T) {
	
	if DEFAULT_FORMAT != JSON_FORMAT {
		t.Errorf("Default format should be " + JSON_FORMAT)
	}
	if HTML_FORMAT != "HTML" {
		t.Errorf("HTML_FORMAT format should be HTML")
	}
	if XML_FORMAT != "XML" {
		t.Errorf("XML_FORMAT should be XML")
	}
	if JSON_FORMAT != "JSON" {
		t.Errorf("JSON_FORMAT should be JSON")
	}
	
}

func TestGetFormatForRequest(t *testing.T) {
	
	var request *http.Request
	
	request = new(http.Request)
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456.json")
	if getFormatForRequest(request) != JSON_FORMAT {
		t.Errorf("getFormatForRequest should be 'JSON' not '%s'", getFormatForRequest(request))
	}
	
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456.xml")
	if getFormatForRequest(request) != XML_FORMAT {
		t.Errorf("getFormatForRequest should be 'XML' not '%s'", getFormatForRequest(request))
	}
	
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456")
	if getFormatForRequest(request) != DEFAULT_FORMAT {
		t.Errorf("getFormatForRequest should be '%s' not '%s'", DEFAULT_FORMAT, getFormatForRequest(request))
	}
	
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456.html")
	if getFormatForRequest(request) != HTML_FORMAT {
		t.Errorf("getFormatForRequest should be '%s' not '%s'", HTML_FORMAT, getFormatForRequest(request))
	}
	request.URL, _ = http.ParseURL(testDomain + "/people/123/groups/456.htm")
	if getFormatForRequest(request) != HTML_FORMAT {
		t.Errorf("getFormatForRequest should be '%s' not '%s'", HTML_FORMAT, getFormatForRequest(request))
	}
	
	
	request.URL = nil
	if getFormatForRequest(request) != DEFAULT_FORMAT {
		t.Errorf("getFormatForRequest should be 'JSON' not '%s'", getFormatForRequest(request))
	}
	
}