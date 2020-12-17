package errs

import "net/http"

type AppError struct {
	Message string `json:"message"`
	ID      int    `json:",omitempty"`
}

func (a AppError) AsMessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		ID:      http.StatusNotFound,
	}
}

func NewNUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		ID:      http.StatusInternalServerError,
	}
}
