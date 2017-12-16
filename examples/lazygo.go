package main

import (
	"fmt"
	"github.com/orkunkaraduman/lazygo"
)

func DuplicateTest() {
	var pSrc *int64
	var src int64
	pSrc = &src
	src = -3 * 1024 * 1024 * 1024 * 1024
	pDst := lazygo.Duplicate(src).(*int64)
	fmt.Printf("pDst := Duplicate(src).(*int64) -> src=%d(0x%x) *pDst=%d(0x%x)\n", src, pSrc, *pDst, pDst)
	pDst2 := lazygo.Duplicate(&src).(**int64)
	fmt.Printf("pDst2 := Duplicate(&src).(**int64) -> src=%d(0x%x) **pDst2=%d(0x%x)\n", src, pSrc, **pDst2, *pDst2)
}

func IndexTest() {
	s := []int{3, 4, 5, 6}
	fmt.Printf("Slice is %v\n", s)
	fmt.Printf("Index of 5 in slice: %d\n", lazygo.Index(s, 5))
	fmt.Printf("Index of 1 in slice: %d\n", lazygo.Index(s, 1))
}

func EachTest() {
	m := map[string]int{"a": 3, "b": 4, "c": 5, "d": 6}
	fmt.Printf("Map is %v\n", m)
	k, v := lazygo.Each(m)
	fmt.Printf("Keys and values in map with Each function (ordered): %v %v\n", k, v)
	k, v = lazygo.Keys(m), lazygo.Values(m)
	fmt.Printf("Keys and values in map with Keys, Values functions: %v %v\n", k, v)
}

func main() {
	DuplicateTest()
	IndexTest()
	EachTest()
}
