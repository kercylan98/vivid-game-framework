package vgameerr

import (
	"errors"
	"fmt"
)

const (
	applicationError = 100000
)

var (
	_        error = (*VGameErr)(nil)
	code2err       = make(map[int]error)
)

func register(code int, msg string) error {
	vge := &VGameErr{
		code: code,
		msg:  msg,
	}

	code2err[code] = vge
	return vge
}

func GetError(code int) error {
	return code2err[code]
}

func GetCode(err error) int {
	var vge *VGameErr
	if errors.As(err, &vge) {
		return vge.code
	}
	return 0
}

func IsError(err error, code int) bool {
	return GetCode(err) == code
}

type VGameErr struct {
	code int
	msg  string
}

func (e *VGameErr) Code() int {
	return e.code
}

func (e *VGameErr) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.code, e.msg)
}
