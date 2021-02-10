package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	ANError  string `json:"error,omitempty"`
}

func (err *apiError) Status() int {
	return err.AStatus
}

func (err *apiError) Message() string {
	return err.AMessage
}

func (err *apiError) Error() string {
	return err.ANError
}

func NewApiErrorFromBytes(bytes []byte) (ApiError, error) {
	var apiErr apiError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid api error json")
	}
	return &apiErr, nil
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
