package errors

import (
	"fmt"
	"net/http"
	"sync"
)

var codes sync.Map

type Coder interface {
	HTTPStatus() int
	String() string
	Code() int
}

type defaultCoder struct{}

func (*defaultCoder) Code() int {
	return 110000
}

func (*defaultCoder) String() string {
	return "error code is undefined"
}

func (*defaultCoder) HTTPStatus() int {
	return http.StatusInternalServerError
}

func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code '0' is ErrUnknown error code")
	}
	if _, ok := codes.Load(coder.Code()); ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}
	codes.Store(coder.Code(), coder)
}

func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}
	if v, ok := err.(*withCode); ok {
		if res, ok := codes.Load(v.code); ok {
			return res.(Coder)
		}
	}
	return new(defaultCoder)
}

func init() {
	Register(new(defaultCoder))
}
