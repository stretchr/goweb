package goweb

import "testing"

func TestRegisteredFormatters(t *testing.T) {
	
	if formatters[JSON_FORMAT] != defaultJsonFormatter {
		t.Errorf("The registered JSON formatter is missing or incorrect")
	}
	
}

func TestGetFormatter(t *testing.T) {
	
	formatters = make(map[string]Formatter)
	var testFormatter *TestFormatter = new(TestFormatter)
	
	RegisterFormatter("TEST", testFormatter)
	
	formatter, error := GetFormatter("TEST")
	
	if error != nil {
		t.Errorf("error should be nil")
	}
	if formatter != testFormatter {
		t.Errorf("formatter is incorrect")
	}
	
}

func TestJsonFormatter(t *testing.T) {
	
	input := "Good morning"
	output, error := defaultJsonFormatter.Format(input)
	
	if error != nil {
		t.Errorf("No errors should have occurred")
	}
	
	if string(output) != "\"Good morning\"" {
		t.Errorf("defaultJsonFormatter.Format didn't do a very good job.")
	}
	
	if defaultJsonFormatter.ContentType() != "application/json" {
		t.Errorf("JsonFormatter content type shouldn't be '%s'.", defaultJsonFormatter.ContentType())
	}
	
}

func TestRegisterFormatter(t *testing.T) {

	formatters = make(map[string]Formatter)
	var testFormatter *TestFormatter = new(TestFormatter)
	
	success, error := RegisterFormatter("TEST", testFormatter)
	
	if formatters["TEST"] != testFormatter {
		t.Errorf("formatters['TEST'] should equal the test formatter instance")
	}
	if !success || error != nil {
		t.Errorf("RegisterFormatter should have been successful")
	}

}

func TestRegisterFormatterAgain(t *testing.T) {
	
	formatters = make(map[string]Formatter)
	var testFormatter *TestFormatter = new(TestFormatter)
	
	success, error := RegisterFormatter("TEST", testFormatter)
	success, error = RegisterFormatter("TEST", testFormatter)
	
	if success {
		t.Errorf("success should have been false")
	}
	if error == nil {
		t.Errorf("An error should have been returned")
	}
	
}

func TestUnregisterFormatter(t *testing.T) {
	
	formatters = make(map[string]Formatter)
	
	var testFormatter *TestFormatter = new(TestFormatter)
	RegisterFormatter("TEST", testFormatter)
	
	success, error := UnregisterFormatter("TEST")
	
	if len(formatters) != 0 {
		t.Errorf("There should be NO items in formatters (not %d)", len(formatters))
	}
	
	if !success || error != nil {
		t.Errorf("UnregisterFormatter should have been successful")
	}
	
	success, error = UnregisterFormatter("TEST")
	
	if success || error == nil {
		t.Errorf("UnregisterFormatter should have been unsuccessful")
	}
	
}