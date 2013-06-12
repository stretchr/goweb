package responders

import (
	codecservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
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
		Transformers
	*/

	// TransformStandardResponseObject transforms the standard response object before it is written to the response if a
	// transformer func has been set via SetStandardResponseObjectTransformer.  Otherwise, it just returns the same
	// object that is passed in.
	TransformStandardResponseObject(ctx context.Context, object map[string]interface{}) (map[string]interface{}, error)

	// SetStandardResponseObjectTransformer sets the function to use to transform the standard response object before it is
	// written to the response.
	//
	// You should use this function to control what kind of response your API makes.
	SetStandardResponseObjectTransformer(transformer func(ctx context.Context, object map[string]interface{}) (map[string]interface{}, error))

	/*
	   Responding
	*/

	// Responds to the Context with the specified status, data and errors.
	Respond(ctx context.Context, status int, data interface{}, errors []string) error

	// WriteResponseObject writes the status code and response object to the HttpResponseWriter in
	// the specified context, in the format best suited based on the request.
	//
	// Goweb uses the WebCodecService to decide which codec to use when responding
	// see http://godoc.org/github.com/stretchr/codecs/services#WebCodecService for more information.
	//
	// This method should be used when the Goweb Standard Response Object does not satisfy the needs of
	// the API, but other Respond* methods are recommended.
	WriteResponseObject(ctx context.Context, status int, responseObject interface{}) error

	// RespondWithData responds with the specified data, no errors and a 200 StatusOK response.
	RespondWithData(ctx context.Context, data interface{}) error

	// RespondWithError responds with the specified error message and status code.
	RespondWithError(ctx context.Context, status int, err string) error
}
