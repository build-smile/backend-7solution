package utils

// CustomError holds the HTTP status code and message.
type CustomError struct {
	Code    int    // HTTP status code
	Message string // Error message
}

func (e *CustomError) Error() string {
	return e.Message
}

// Utility function to create a new CustomError
func NewCustomError(code int, message string) *CustomError {
	return &CustomError{Code: code, Message: message}
}
