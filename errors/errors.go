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
	pc        uintptr
}

func (err Error) Error() {
	if len(err.prefix) != 0 {
		return strings.Join(err.prefix, ".") + ":" + s.err.Error()
	}
	return err.err.Error()
}

func (err Error) SetCode(code ErrCode) {
	err.code = code
}

func (err Error) Code() ErrCode {
	return err.code
}

func (err Error) MatchCode(other Error) {
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
	err = Wrap(2, err, nil)
	err.prefix = append(err.prefix, prefix)
	return err
}

func Wrap(depth int, err error, fields Fields) error {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return nil
	}
	e, ok := err.(Error)
	if !ok {
		err = Error{
			stackInfo: fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line),
			err:       err,
		}
		err.ptr = uintptr(unsafe.Pointer(&err))
		return err
	}
	err.stackInfo = fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
	err.ptr = uintptr(unsafe.Pointer(&err))
	return err
}
