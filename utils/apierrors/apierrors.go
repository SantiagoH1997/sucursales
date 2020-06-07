package apierrors

import (
	"net/http"
)

// APIError is the interface for the error sent to the user
type APIError interface {
	StatusCode() int
	Parse() map[string]interface{}
}

type genericError struct {
	statusCode int
	code       string
	message    string
}

// Parse parses the error
func (ge *genericError) Parse() map[string]interface{} {
	formattedErr := make(map[string]interface{})

	formattedErr["statusCode"] = ge.statusCode
	formattedErr["code"] = ge.code
	formattedErr["message"] = ge.message

	return formattedErr
}

// StatusCode returns the error's status code
func (ge *genericError) StatusCode() int {
	return ge.statusCode
}

// NewNotFound returns a Not Found error
func NewNotFound(message string) APIError {
	var err APIError = &genericError{
		statusCode: http.StatusNotFound,
		code:       "not_found",
		message:    message,
	}
	return err
}

// NewInternalServerError returns an internal server error
func NewInternalServerError(message string) APIError {
	var err APIError = &genericError{
		statusCode: http.StatusInternalServerError,
		code:       "internal_server_error",
		message:    message,
	}
	return err
}

// NewBadRequest returns a bad request error
func NewBadRequest(message string) APIError {
	var err APIError = &genericError{
		statusCode: http.StatusBadRequest,
		code:       "bad_request",
		message:    message,
	}
	return err
}

type validationError struct {
	statusCode int
	code       string
	message    string
	fields     []ErrorField
}

// Parse parses the validation error
func (ve *validationError) Parse() map[string]interface{} {
	formattedErr := make(map[string]interface{})

	formattedErr["statusCode"] = ve.statusCode
	formattedErr["code"] = ve.code
	formattedErr["message"] = ve.message
	formattedErr["fields"] = ve.fields

	return formattedErr
}

// StatusCode returns a 400
func (ve *validationError) StatusCode() int {
	return ve.statusCode
}

// NewValidationError returns a validationError with a BadRequestError inside
func NewValidationError(message string, fields []ErrorField) APIError {
	var err APIError = &validationError{
		statusCode: http.StatusBadRequest,
		code:       "bad_request",
		message:    message,
		fields:     fields,
	}
	return err
}

// ErrorField is a struct specifying a missing/invalid field and an error
type ErrorField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
