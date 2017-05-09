package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	fields    Fields
	prefix    []string
	code      Errcode
	err       error
	stackInfo string
}

func (err Error) Error() {
	if len(err.prefix) != 0 {
		return strings.Join(err.prefix, ".") + ":" + s.err.Error()
	}
	return err.err.Error()
}

type Fields map[string]interface{}

func Trace(err error) error {
	return Wrap(2, err, nil)
}

func TracePrefix(err error, prefix string) error {
	err := Wrap(2, err, nil)
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
		return Error{
			stackInfo: fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line),
			err:       err,
		}
	}
	return err.stackInfo = fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
}
