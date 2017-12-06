package lazycore

import (
	"errors"
	"reflect"
	"unsafe"
)

// Duplicate returns the pointer of duplicated value of i.
func Duplicate(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t == nil {
		return nil
	}
	//src := reflect.Indirect(reflect.ValueOf(&i)).InterfaceData()[1]
	var src uintptr
	if t.Kind() != reflect.Ptr {
		src = *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(uintptr(0))))
	} else {
		src = uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(uintptr(0))
	}
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

// Index returns the index s from v. If not found return -1. Panics v is not array or slice.
func Index(v interface{}, s interface{}) int {
	if v == nil {
		return -1
	}
	rv := reflect.ValueOf(v)
	if k := rv.Type().Kind(); k == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	if k := rv.Type().Kind(); k != reflect.Array && k != reflect.Slice {
		panic(errors.New("v must be Array or Slice"))
	}
	for i, l := 0, rv.Len(); i < l; i++ {
		if reflect.DeepEqual(rv.Index(i).Interface(), s) {
			return i
		}
	}
	return -1
}

// Each returns the keys slice and values slice from v. Panics v is not map.
func Each(v interface{}) (keys []interface{}, values []interface{}) {
	if v == nil {
		return
	}
	rv := reflect.ValueOf(v)
	if k := rv.Type().Kind(); k == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	if k := rv.Type().Kind(); k != reflect.Map {
		panic(errors.New("v must be Map"))
	}
	rkeys := rv.MapKeys()
	for _, rkey := range rkeys {
		keys = append(keys, rkey.Interface())
		values = append(values, rv.MapIndex(rkey).Interface())
	}
	return
}

// Keys returns the keys slice from v. Panics v is not map.
func Keys(v interface{}) []interface{} {
	r, _ := Each(v)
	return r
}

// Values returns the values slice from v. Panics v is not map.
func Values(v interface{}) []interface{} {
	_, r := Each(v)
	return r
}
