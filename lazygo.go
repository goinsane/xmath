package lazygo

import (
	"reflect"
	"unsafe"
)

func Duplicate(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	src := reflect.Indirect(reflect.ValueOf(&i)).InterfaceData()[1]
	r := reflect.New(t)
	dst := r.Pointer()
	en := src + t.Size()
	for src < en {
		*(*byte)(unsafe.Pointer(dst)) = *(*byte)(unsafe.Pointer(src))
		src++
		dst++
	}
	return r.Interface()
}
