package strbyte

import (
	"fmt"
	"testing"
)

func TestByte2Str(t *testing.T) {
	b := []byte{'t', 'e', 's', 't'}
	s := Byte2Str(b)
	fmt.Println(s)
}

func TestStr2Byte(t *testing.T) {
	s := "test"
	b := Str2Byte(s)
	fmt.Println(b)
}
