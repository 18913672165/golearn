package strbyte

import (
	"reflect"
	"unsafe"
)

//convert string to []byte, no copy
func Str2Byte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := &reflect.SliceHeader{}
	sliceHeader.Data = stringHeader.Data
	sliceHeader.Len = stringHeader.Len
	sliceHeader.Cap = stringHeader.Len
	b := (*[]byte)(unsafe.Pointer(sliceHeader))
	return *b
}

//convert []byte to string, no copy
func Byte2Str(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringHeader := &reflect.StringHeader{}
	stringHeader.Data = sliceHeader.Data
	stringHeader.Len = sliceHeader.Len
	s := (*string)(unsafe.Pointer(stringHeader))
	return *s
}
