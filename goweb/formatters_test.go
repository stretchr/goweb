package goweb

import (
	"testing"
	"os"
)

func TestAddFormatter(t *testing.T) {
	ClearFormatters()
	
	var testFormatter = new(TestFormatter)
	var deciderForOne FormatterDecider = func(c *Context) (bool, os.Error) {	
		return c.Format == "ONE", nil
	}
	
	AddFormatter(deciderForOne, testFormatter)
	
	if len(formatterOptions) != 1 {
		t.Errorf("len(formatterOptions) should be 1, but was %d", len(formatterOptions))
	}
	
	if formatterOptions[0].Formatter != testFormatter {
		t.Error("formatters[0].Formatter should be testFormatter")
	}
	if formatterOptions[0].Decider != deciderForOne {
		t.Error("formatters[0].Decider should be deciderForOne")
	}
	
	
}

func TestClearFormatters(t *testing.T) {
	
	var testFormatter = new(TestFormatter)
	var testFormatter2 = new(TestFormatter2)
	
	var deciderForOne FormatterDecider = func(c *Context) (bool, os.Error) {	
		return c.Format == "ONE", nil
	}
	var deciderForTwo FormatterDecider = func(c *Context) (bool, os.Error) {	
		return c.Format == "TWO", nil
	}
	
	AddFormatter(deciderForOne, testFormatter)
	AddFormatter(deciderForTwo, testFormatter2)
	
	ClearFormatters()
	
	if len(formatterOptions) != 0 {
		t.Error("ClearFormatter didn't do that!")
	}
	
}

func TestGetFormatter(t *testing.T) {
	ClearFormatters()
	
	var testFormatter = new(TestFormatter)
	var testFormatter2 = new(TestFormatter2)
	
	var deciderForOne FormatterDecider = func(c *Context) (bool, os.Error) {	
		return c.Format == "ONE", nil
	}
	var deciderForTwo FormatterDecider = func(c *Context) (bool, os.Error) {	
		return c.Format == "TWO", nil
	}
	
	AddFormatter(deciderForOne, testFormatter)
	AddFormatter(deciderForTwo, testFormatter2)
	
	context := new(Context)
	
	context.Format = "ONE"
	actualFormatter, err := GetFormatter(context)
	if actualFormatter != testFormatter {
		t.Error("testFormatter expected")
	}
	
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	
	context.Format = "TWO"
	actualFormatter, err = GetFormatter(context)
	if actualFormatter != testFormatter2 {
		t.Error("testFormatter2 expected")
	}
	
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

}

func TestGetFormatter_ChecksInReverseOrder(t *testing.T) {
	ClearFormatters()
	
	var testFormatter = new(TestFormatter)
	var testFormatter2 = new(TestFormatter2)
	
	var deciderForOne FormatterDecider = func(c *Context) (bool, os.Error) {	
		return true, nil
	}
	var deciderForTwo FormatterDecider = func(c *Context) (bool, os.Error) {	
		return true, nil
	}
	
	AddFormatter(deciderForOne, testFormatter)
	AddFormatter(deciderForTwo, testFormatter2)
	
	// both formatter will be a hit, but we want the
	// latest one to be considered first
	c := new(Context)
	formatter, err := GetFormatter(c)
	
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if formatter != testFormatter2 {
		t.Error("GetFormatter should look them up in reverse order")
	}
	
}

func TestGetFormatter_ThrowsErrorIfNoFormatters(t *testing.T) {
	
	ClearFormatters()
	
	c := new(Context)
	formatter, err := GetFormatter(c)

	if err == nil || formatter != nil {
		t.Error("Calling GetFormatter with no formatters should raise an error")
	}
	
}

func TestGetFormatter_ThrowsErrorIfDeciderDoes(t *testing.T) {
	ClearFormatters()
	
	var error os.Error = os.NewError("Something went wrong!")
	var errorDecider FormatterDecider = func(c *Context) (bool, os.Error) {
		return false, error
	}
	
	AddFormatter(errorDecider, nil)
	
	c := new(Context)
	_, returnedError := GetFormatter(c)
	
	if returnedError != error {
		t.Error("Any errors returned by the decider functions should be returned by GetFormatter")
	}
	
}


func TestConfigureDefaultFormatterOptions(t *testing.T) {
	ClearFormatters()
	
	ConfigureDefaultFormatters()
	
	c := new(Context)
	c.Format = "JSON"
	
	formatter, _ := GetFormatter(c)
	
	if formatter != defaultJsonFormatter {
		t.Error("ConfigureDefaultFormatters didn't set up the defualt JSON formatter")
	}
	
}

