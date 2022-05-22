package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

//402
func NewPaymentRequiredError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusPaymentRequired,
	}
}

//404
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusNotFound,
	}
}

//500
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusInternalServerError,
	}
}

func(e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

//422
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusUnprocessableEntity,
	}
}

//401
func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

//403
func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
