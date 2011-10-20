package goweb

import (
	"os"
	"json"
)

// Interface describing an object responsible for formatting
// other objects
type Formatter interface {
	Format(input interface{}) ([]uint8, os.Error)
	ContentType() string
}

// Function responsible for deciding if a formatter should be
// used or not
type FormatterDecider func(*Context) (bool, os.Error)

// Struct to hold a potential formatter option.  The Decider will
// be used to determine if the specified Formatter will be used
type FormatterOption struct {
	Decider FormatterDecider
	Formatter Formatter
}

// Makes a new formatter option with the specified decider and formatter
func makeFormatterOption(decider FormatterDecider, formatter Formatter) *FormatterOption {
	var option *FormatterOption = new(FormatterOption)
	option.Decider = decider
	option.Formatter = formatter
	return option
}

// Internal collection of formatter options
var formatterOptions []*FormatterOption

// Adds a formatter decider 
func AddFormatter(decider FormatterDecider, formatter Formatter) {
	formatterOptions = append(formatterOptions, makeFormatterOption(decider, formatter))
}

// Clears all formatters (including default internal ones)
func ClearFormatters() {
	formatterOptions = make([]*FormatterOption, 0)
}

// Gets the relevant formatter for the specified context or
// returns an error if no formatter is found
func GetFormatter(context *Context) (Formatter, os.Error) {
	
	for i := len(formatterOptions); i > 0; i-- {
		
		hit, error := formatterOptions[i-1].Decider(context)
		
		// return the error if there was one
		if error != nil {
			return nil, error
		}
		
		// if it was a hit, return this formatter
		if hit {
			return formatterOptions[i-1].Formatter, nil
		}
		
	}
	
	return nil, os.NewError("No suitable Formatter could be found to deal with that request, consider calling ConfigureDefaultFormatters() or AddFormatter().  See http://code.google.com/p/goweb/wiki/APIDocumentation#Formatters")
	
}


/*

	Default internal formatters

*/

// Formatter for JSON
type JsonFormatter struct{}

// Converts a data object into JSON
func (f *JsonFormatter) Format(input interface{}) ([]uint8, os.Error) {
	output, error := json.Marshal(input)
	return []uint8(output), error
}

// Gets the "application/json" content type
func (f *JsonFormatter) ContentType() string {
	return "application/json"
}

// Default instance of the JSON formatter
var defaultJsonFormatter *JsonFormatter = new(JsonFormatter)

// Adds the default formatters to goweb so that
func ConfigureDefaultFormatters() {
	
	AddFormatter(func(c *Context) (bool, os.Error) {	
		return c.Format == "JSON", nil
	}, defaultJsonFormatter)
	
}