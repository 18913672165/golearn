package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"
)

type ErrCode int

type Fields map[string]interface{}

type Error struct {
	fields    Fields
	prefix    []string
	code      ErrCode
	err       error
	stackInfo string
	ptr       uintptr
}

func (err Error) Error() string {
	if len(err.prefix) != 0 {
		return strings.Join(err.prefix, ".") + ":" + err.err.Error()
	}
	return err.err.Error()
}

func (err Error) SetCode(code ErrCode) {
	err.code = code
}

func (err Error) Code() ErrCode {
	return err.code
}

func (err Error) MatchCode(other Error) bool {
	return err.code == other.code
}

func (err Error) EqualTo(other Error) bool {
	if err.ptr == other.ptr {
		return true
	}
	return false
}

func Equal(err1, err2 Error) bool {
	return err1.EqualTo(err2)
}

func MatchCode(err1, err2 Error) bool {
	return err1.MatchCode(err2)
}

func Trace(err error) error {
	return Wrap(2, err, nil)
}

func TracePrefix(err error, prefix string) error {
	errWrapped := Wrap(2, err, nil)
	errWrapped.prefix = append(errWrapped.prefix, prefix)
	return errWrapped
}

func Wrap(depth int, err error, fields Fields) *Error {
	pc, _, line, ok := runtime.Caller(depth)
	if !ok {
		return nil
	}
	e, ok := err.(Error)
	if !ok {
		errWrapped := Error{
			stackInfo: fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line),
			err:       err,
		}
		errWrapped.ptr = uintptr(unsafe.Pointer(&errWrapped))
		return &errWrapped
	}
	e.stackInfo = fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
	e.ptr = uintptr(unsafe.Pointer(&e))
	return &e
}
