package goweb

import "testing"

func TestGetPathSegments(t *testing.T) {
	
	var segments []string = getPathSegments(routePath)
	
	if (segments[0] != "people") {
		t.Errorf("'%s' expected to be 'people'", segments[0])
	}
	if (segments[1] != "{id}") {
		t.Errorf("'%s' expected to be '{id}'", segments[1])
	}
	if (segments[2] != "groups") {
		t.Errorf("'%s' expected to be 'groups'", segments[2])
	}
	if (segments[3] != "{group_id}") {
		t.Errorf("'%s' expected to be '{group_id}'", segments[3])
	}
	if (segments[4] != ".json") {
		t.Errorf("'%s' expected to be '.json'", segments[4])
	}
	
}

func TestGetPathSegments_WithoutExtension(t *testing.T) {
	
	var segments []string = getPathSegments(routePathWithoutExtension)
	
	if (segments[0] != "people") {
		t.Errorf("'%s' expected to be 'people'", segments[0])
	}
	if (segments[1] != "{id}") {
		t.Errorf("'%s' expected to be '{id}'", segments[1])
	}
	if (segments[2] != "groups") {
		t.Errorf("'%s' expected to be 'groups'", segments[2])
	}
	if (segments[3] != "{group_id}") {
		t.Errorf("'%s' expected to be '{group_id}'", segments[3])
	}
	
}

func TestIsDynamicSegment(t *testing.T) {
	
	if (isDynamicSegment("no")) {
		t.Errorf("'no' is not a dynamic segment")
	}
	if (isDynamicSegment("{not-quite")) {
		t.Errorf("'{not-quite' is not a dynamic segment")
	}
	if (isDynamicSegment("not-me-either}")) {
		t.Errorf("'not-me-either}' is not a dynamic segment")
	}
	
	if (!isDynamicSegment("{i-am}")) {
		t.Errorf("'i-am' is a dynamic segment")
	}
	
}

func TestIsExtensionSegment(t *testing.T) {
	
	if isExtensionSegment("nope") {
		t.Errorf("Should not be an extension segment")
	}
	if isExtensionSegment("{no-way}") {
		t.Errorf("Should not be an extension segment")
	}
	if !isExtensionSegment(".yes") {
		t.Errorf("Should be an extension segment")
	}
	
	
}

func TestGetParameterValueMapFromPath(t *testing.T) {
	
	var route *Route = makeRouteFromPath(routePath)
	var paramKeys ParameterKeyMap = route.parameterKeys
	var paramValues ParameterValueMap = getParameterValueMap(paramKeys, "/people/123/groups/456")
	
	if (len(paramValues) != 2) {
		t.Errorf("paramValues should have 2 items")
	}
	if (paramValues["id"] != "123") {
		t.Errorf("paramKeys['id'] expected to be '123', but was %s", paramValues["id"])
	}
	if (paramValues["group_id"] != "456") {
		t.Errorf("paramKeys['group_id'] expected to be '456', but was %s", paramValues["group_id"])
	}
	
}


func assertFileExtension(t *testing.T, path string, expected string) {
	if getFileExtension(path) != expected {
		t.Errorf("getFileExtension(\"%s\") should be \"%s\" but was \"%s\"", path, expected, getFileExtension(path))
	}
}
func TestGetFileExtension(t *testing.T) {
	
	assertFileExtension(t, "/people.json", "json")
	assertFileExtension(t, "http://www.test.com/people.yml", "yml")
	assertFileExtension(t, "/people/123/groups/177.Xml", "Xml")
	assertFileExtension(t, "/people/123/groups/177.XML", "XML")
	assertFileExtension(t, "/people.d", "d")
	assertFileExtension(t, "/people", "")
	assertFileExtension(t, "/people/123/groups/456", "")
	
}