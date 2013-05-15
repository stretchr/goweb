package responders

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
)

/*
  APIResponder represents objects capable of provide API responses.
*/
type APIResponder interface {

	/*
	   Codec services
	*/

	// SetCodecService sets the codec service to use.
	SetCodecService(codecservices.CodecService)

	// GetCodecService gets the codec service that will be used by this object.
	GetCodecService() codecservices.CodecService

	/*
	   Responding
	*/

	// Responds to the Context with the specified status, data and errors.
	Respond(ctx context.Context, status int, data interface{}, errors []string) error

	// WriteResponseObject writes the status code and response object to the HttpResponseWriter in
	// the specified context, in the format best suited based on the request.
	//
	// Goweb uses the WebCodecService to decide which codec to use when responding
	// see http://godoc.org/github.com/stretchrcom/codecs/services#WebCodecService for more information.
	//
	// This method should be used when the Goweb Standard Response Object does not satisfy the needs of
	// the API, but other Respond* methods are recommended.
	WriteResponseObject(ctx context.Context, status int, responseObject interface{}) error

	// RespondWithData responds with the specified data, no errors and a 200 StatusOK response.
	RespondWithData(ctx context.Context, data interface{}) error

	// RespondWithError responds with the specified error message and status code.
	RespondWithError(ctx context.Context, status int, err string) error
}
