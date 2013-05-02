package goweb

import (
	"github.com/stretchrcom/goweb/api"
)

// API is an api.APIResponder which provides the ability to make API responses.
var API api.APIResponder = new(api.GowebAPIResponder)
