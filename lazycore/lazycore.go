package lazycore

import (
	"reflect"
	"unsafe"
	"github.com/pkg/errors"
)

// Duplicate returns the pointer of duplicated value of i.
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

// Index returns the index s from v. If not found return -1. Panics v is not array or slice.
func Index(v interface{}, s interface{}) int {
	rv := reflect.ValueOf(v)
	if k := rv.Type().Kind(); k == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	if k := rv.Type().Kind(); k != reflect.Array && k != reflect.Slice {
		panic(errors.New("v must be Array or Slice"))
	}
	for i := 0; i < rv.Len(); i++ {
		if reflect.DeepEqual(rv.Index(i).Interface(), s) {
			return i
		}
	}
	return -1
}

// Each returns the keys slice and values slice from v. Panics v is not map.
func Each(v interface{}) ([]interface{}, []interface{}) {
	rv := reflect.ValueOf(v)
	if k := rv.Type().Kind(); k == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}
	if k := rv.Type().Kind(); k != reflect.Map {
		panic(errors.New("v must be Map"))
	}
	var keys []interface{}
	var vals []interface{}
	rkeys := rv.MapKeys()
	for _, rkey := range rkeys {
		keys = append(keys, rkey.Interface())
		vals = append(vals, rv.MapIndex(rkey).Interface())
	}
	return keys, vals
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
