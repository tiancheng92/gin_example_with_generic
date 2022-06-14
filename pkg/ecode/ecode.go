package ecode

import (
	"gin_example_with_generic/pkg/errors"
)

type ErrCode struct {
	ErrCode    int
	HttpStatus int
	Message    string
}

func (coder ErrCode) Code() int {
	return coder.ErrCode
}

func (coder ErrCode) String() string {
	return coder.Message
}

func (coder ErrCode) HTTPStatus() int {
	if coder.HttpStatus == 0 {
		return 500
	}
	return coder.HttpStatus
}

func register(code int, httpStatus int, message string) {
	errors.MustRegister(&ErrCode{
		ErrCode:    code,
		HttpStatus: httpStatus,
		Message:    message,
	})
}
