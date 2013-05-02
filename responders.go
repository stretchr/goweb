package goweb

import (
	"github.com/stretchrcom/goweb/responders"
)

// API is an responders.APIResponder which provides the ability to make API responses.
var API responders.APIResponder = new(responders.GowebAPIResponder)
