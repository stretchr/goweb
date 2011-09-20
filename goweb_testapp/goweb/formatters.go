package goweb

import (
	"os"
	"json"
	"fmt"
)

type Formatter interface {
	Format(input interface{}) ([]uint8, os.Error)
	ContentType() string
}


// Formatter for JSON
type JsonFormatter struct {}
func (f *JsonFormatter) Format(input interface{}) ([]uint8, os.Error) {
	output, error := json.Marshal(input)
	return []uint8(output), error
}
func (f *JsonFormatter) ContentType() string {
	return "application/json"
}
var defaultJsonFormatter *JsonFormatter = new(JsonFormatter)


// Internal map between formats and the classes that perform
// that formatting
var formatters map[string]Formatter = map[string]Formatter{ "JSON": defaultJsonFormatter }

// Gets the formatter object responsible for the specified
// format
func GetFormatter(format string) (Formatter, os.Error) {
	
	if formatters[format] == nil {
		return nil, fmt.Errorf("I don't know how to make files of that kind - Try changing the extension to .json")
	}
	
	return formatters[format], nil
	
}

// Registers a formatter object to a format 
func RegisterFormatter(format string, formatter Formatter) (bool, os.Error) {
	
	if formatters[format] != nil {
		return false, fmt.Errorf("A formatter is already registered for \"%s\".", format)
	}
	
	// register the formatter
	formatters[format] = formatter
	
	// return success
	return true, nil
}

// Unregisters a formatter object from a format
func UnregisterFormatter(format string) (bool, os.Error) {
	
	if formatters[format] == nil {
		return false, fmt.Errorf("Cannot unregister. No formatter is registered for \"%s\".", format)
	}
	
	formatters[format] = nil, false
	
	return true, nil
}


