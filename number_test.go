package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleNewDecimal() {
	n := xmath.NewDecimal(1)
	for i := 0; i < 10; i++ {
		k := 0.33 * float64(i)
		n.SetFloat64(k)
		fmt.Println(n)
	}

	// Output:
	// 0
	// 0.3
	// 0.7
	// 1
	// 1.3
	// 1.7
	// 2
	// 2.3
	// 2.6
	// 3
}
