package main

import (
	"fmt"
	"github.com/orkunkaraduman/lazygo"
)

func DuplicateTest() {
	var pSrc, pDst *int64
	var src, dst int64
	pSrc = &src
	src = 3 * 1024 * 1024 * 1024 * 1024
	pDst = lazygo.Duplicate(src).(*int64)
	dst = *pDst
	fmt.Printf("pDst := Duplicate(src).(*int64) -> src=%d(0x%x) dst=%d(0x%x)\n", src, pSrc, dst, pDst)
}

func main() {
	DuplicateTest()
}
