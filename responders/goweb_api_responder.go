package responders

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
)

type GowebAPIResponder struct {
	httpResponder HTTPResponder
	codecService  codecservices.CodecService
}

func NewGowebAPIResponder(httpResponder HTTPResponder) *GowebAPIResponder {
	api := new(GowebAPIResponder)
	api.httpResponder = httpResponder
	return api
}

// SetCodecService sets the codec service to use.
func (a *GowebAPIResponder) SetCodecService(service codecservices.CodecService) {
	a.codecService = service
}

// GetCodecService gets the codec service that will be used by this object.
func (a *GowebAPIResponder) GetCodecService() codecservices.CodecService {

	if a.codecService == nil {
		a.codecService = new(codecservices.WebCodecService)
	}

	return a.codecService
}

// WriteResponseObject writes the status code and response object to the HttpResponseWriter in
// the specified context.
func (a *GowebAPIResponder) WriteResponseObject(ctx context.Context, status int, responseObject interface{}) error {

	service := a.GetCodecService()
	codec, codecError := service.GetCodec("application/json")

	if codecError != nil {
		return codecError
	}

	output, marshalErr := codec.Marshal(responseObject, nil)

	if marshalErr != nil {
		return marshalErr
	}

	// use the HTTP responder to respond
	a.httpResponder.With(ctx, status, output)

	return nil

}

// Responds to the Context with the specified status, data and errors.
func (a *GowebAPIResponder) Respond(ctx context.Context, status int, data interface{}, errors []string) error {

	sro := map[string]interface{}{"s": status}

	if data != nil {
		sro["d"] = data
	}
	if len(errors) > 0 {
		sro["e"] = errors
	}

	return a.WriteResponseObject(ctx, status, sro)

}
