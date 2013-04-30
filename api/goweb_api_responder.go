package api

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
)

type GowebAPIResponder struct {
	codecService codecservices.CodecService
}

func (a *GowebAPIResponder) SetCodecService(service codecservices.CodecService) {
	a.codecService = service
}

func (a *GowebAPIResponder) GetCodecService() codecservices.CodecService {

	if a.codecService == nil {
		a.codecService = new(codecservices.WebCodecService)
	}

	return a.codecService
}

// WriteResponseObject writes the status code and response object to the HttpResponseWriter in
// the specified context.
func (a *GowebAPIResponder) WriteResponseObject(ctx *context.Context, status int, responseObject interface{}) {

	service := a.GetCodecService()
	codec, _ := service.GetCodec("application/json")
	output, _ := codec.Marshal(responseObject, nil)
	ctx.HttpResponseWriter().Write(output)
	ctx.HttpResponseWriter().WriteHeader(status)

}

func (a *GowebAPIResponder) Respond(ctx *context.Context, status int, data interface{}, errors []string) {

	sro := map[string]interface{}{"s": status}

	if data != nil {
		sro["d"] = data
	}
	if len(errors) > 0 {
		sro["e"] = errors
	}

	a.WriteResponseObject(ctx, status, sro)

}
