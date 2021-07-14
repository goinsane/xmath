package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleReal() {
	n := new(xmath.Real)
	for i := 0; i < 20; i++ {
		k := 0.25 * float64(i)
		n.SetFloat64(k)
		fmt.Println(n, k)
	}

	// Output:
	// 0 0
	// 0 0.25
	// 1 0.5
	// 1 0.75
	// 1 1
	// 1 1.25
	// 2 1.5
	// 2 1.75
	// 2 2
	// 2 2.25
	// 3 2.5
	// 3 2.75
	// 3 3
	// 3 3.25
	// 4 3.5
	// 4 3.75
	// 4 4
	// 4 4.25
	// 5 4.5
	// 5 4.75
}

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
