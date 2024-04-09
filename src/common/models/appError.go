package models

// AppError represents an application-specific error with additional context.
type AppError struct {
	Message string `json:"message"`
}

// Error returns the error message.
func (e *AppError) Error() string {
	return e.Message
}
