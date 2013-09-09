package goweb

import (
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/responders"
)

// CodecService is the service class that provides codec capabilities to Goweb.
var CodecService codecsservices.CodecService = codecsservices.NewWebCodecService()

// Respond is a responders.HTTPResponder which provides the ability to make HTTP responses.
//
// This allows a simple interface for making normal web responses such as the following:
//
//     return goweb.Respond.WithStatus(ctx, 404)
var Respond responders.HTTPResponder = new(responders.GowebHTTPResponder)

// API is a responders.APIResponder which provides the ability to make API responses.
//
// This allows a simple interface for making API responses such as the following:
//
//     return goweb.API.Respond(ctx, 404, nil, []string{"File not found"})
var API responders.APIResponder = responders.NewGowebAPIResponder(CodecService, Respond)
