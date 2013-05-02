package goweb

import (
	"github.com/stretchrcom/goweb/responders"
)

// Respond is a responders.HTTPResponder which provides the ability to make HTTP responses.
var Respond responders.HTTPResponder = new(responders.GowebHTTPResponder)

// API is a responders.APIResponder which provides the ability to make API responses.
var API responders.APIResponder = responders.NewGowebAPIResponder(Respond)
