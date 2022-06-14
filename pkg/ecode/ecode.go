package ecode

import (
	"gin_example_with_generic/pkg/errors"
	"net/http"
)

type ErrCode struct {
	ErrCode    int
	HttpStatus int
	Message    string
}

func (e *ErrCode) Code() int {
	return e.ErrCode
}

func (e *ErrCode) String() string {
	return e.Message
}

func (e *ErrCode) HTTPStatus() int {
	if e.HttpStatus == 0 {
		return http.StatusInternalServerError
	}
	return e.HttpStatus
}

func register(code int, httpStatus int, message string) {
	errors.MustRegister(&ErrCode{
		ErrCode:    code,
		HttpStatus: httpStatus,
		Message:    message,
	})
}
