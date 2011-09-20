package goweb

import "testing"

func TestMakeStandardResponse(t *testing.T) {
	
	var r *standardResponse
	r = makeStandardResponse()
	
	if r.S != 200 {
		t.Errorf("r.S should be 200")
	}
	if r.E != nil {
		t.Errorf("r.E should start as nil")
	}
	
}

func TestMakeSuccessfulStandardResponse(t *testing.T) {
	
	var data string = "This is the data"
	
	var r *standardResponse
	r = makeSuccessfulStandardResponse("123", 200, data)
	
	if r.S != 200 {
		t.Errorf("r.s should be 200 not %d", r.S)
	}
	if r.C != "123" {
		t.Errorf("r.c should be '123' not '%s'", r.C)
	}
	if r.D != data {
		t.Errorf("r.d should be '%s' not '%s'", data, r.D)
	}
	if len(r.E) > 0 {
		t.Errorf("There should be no errors")
	}
	
}

func TestMakeFailureStandardResponse(t *testing.T) {
	
	var r *standardResponse
	r = makeFailureStandardResponse("123", 404)
	
	if r.S != 404 {
		t.Errorf("r.s should be 200 not %d", r.S)
	}
	if r.C != "123" {
		t.Errorf("r.c should be '123' not '%s'", r.C)
	}
	if len(r.E) != 1 {
		t.Errorf("There should be 1 error")
	}
	
}