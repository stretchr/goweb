package goweb

import (
	"testing"
	"reflect"
)

func TestAddFormatter(t *testing.T) {
	ClearFormatters()

	var testFormatter = new(TestFormatter)

	AddFormatter(testFormatter)

	if len(formatters) != 1 {
		t.Errorf("len(formatters) should be 2, but was %d", len(formatters))
	}

	if formatters[0] != testFormatter {
		t.Error("formatters[0] should be testFormatter")
	}

}

func TestClearFormatters(t *testing.T) {

	AddFormatter(new(TestFormatter))
	AddFormatter(new(TestFormatter2))

	ClearFormatters()

	if formatters != nil {
		t.Error("ClearFormatter didn't do that!")
	}

}

func TestGetFormatter(t *testing.T) {
	ClearFormatters()

	var testFormatter = new(TestFormatter)
	var testFormatter2 = new(TestFormatter2)

	AddFormatter(testFormatter)
	AddFormatter(testFormatter2)

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

	testFormatterA := new(TestFormatter)
	testFormatterB := new(TestFormatter)

	AddFormatter(testFormatterA)
	AddFormatter(testFormatterB)

	// both formatters will be a hit, but we want the
	// latest one added to be considered first
	c := new(Context)
	c.Format = "ONE"
	formatter, err := GetFormatter(c)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if formatter != testFormatterB {
		t.Errorf("GetFormatter should look them up in reverse order expected: %p got %p", testFormatterB, formatter)
	}

}

func TestGetFormatter_ThrowsErrorIfNoFormatters(t *testing.T) {

	ClearFormatters()

	c := new(Context)

	_, err := GetFormatter(c)
	if err == nil {
		t.Error("Calling GetFormatter with no formatters should raise an error")
	}

}

func TestConfigureDefaultFormatterOptions(t *testing.T) {
	ClearFormatters()

	ConfigureDefaultFormatters()

	c := new(Context)
	c.Format = "encoding/json"

	formatter, _ := GetFormatter(c)

	if reflect.TypeOf(formatter).Elem().Name() != "JsonFormatter" {
		t.Error("ConfigureDefaultFormatters didn't set up the defualt JSON formatter")
	}

}
