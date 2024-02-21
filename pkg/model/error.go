package model

import "fmt"

// RespStatus
// https://github.com/omniti-labs/jsend
type RespStatus string

const (
	// 2xx All went well, and (usually) some data was returned
	RespSuccess RespStatus = "success"
	// 4xx There was a problem with the data submitted, or some pre-condition of the API call wasn't satisfied
	RespFail RespStatus = "fail"
	// 5xx An error occurred in processing the request, i.e. an exception was thrown
	RespError RespStatus = "error"
)

type Response struct {
	// 0 success, otherwise error code
	Code int `json:"code"`
	// Success, Fail, Error
	Status RespStatus `json:"status"`
	// Data returned
	Data interface{} `json:"data,omitempty"`
	// Error message, if any
	Message string `json:"message,omitempty"`
}

type Error struct {
	HTTPStatus int         `json:"-"`
	Code       int         `json:"code"`
	Status     RespStatus  `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewError(httpStatus int, code int, status RespStatus, message string) *Error {
	return &Error{
		HTTPStatus: code,
		Code:       code,
		Status:     status,
		Message:    message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
