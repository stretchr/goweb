package goweb

import (
	"github.com/stretchrcom/goweb/responders"
)

// Respond is a responders.HTTPResponder which provides the ability to make HTTP responses.
//
// This allows a simple interface for making normal web responses such as the following:
//
//     goweb.Respond.WithStatus(ctx, 404)
var Respond responders.HTTPResponder = new(responders.GowebHTTPResponder)

// API is a responders.APIResponder which provides the ability to make API responses.
//
// This allows a simple interface for making API resposnes such as the following:
//
//     goweb.API.Respond(ctx, 404, nil, []string{"File not found"})
var API responders.APIResponder = responders.NewGowebAPIResponder(Respond)
