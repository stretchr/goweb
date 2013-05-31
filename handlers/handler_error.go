package handlers

// HandlerError represents an error that was thrown by a particular Handler.
type HandlerError struct {
	// Handler is the Handler that returned the error.
	Handler Handler
	// OriginalError is the error returned by the Handler.
	OriginalError error
}

// Error gets the error string from the OriginalError.
func (e HandlerError) Error() string {
	return e.OriginalError.Error()
}
