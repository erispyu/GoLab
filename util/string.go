package util

import (
	"reflect"
	"unsafe"
)

func String2Bytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh)) //nolint
}
