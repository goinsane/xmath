package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleNewStepper() {
	var err error
	_, err = xmath.NewStepper(2, 10, 0.1, 3.01, 2.31)
	fmt.Println(err)
	_, err = xmath.NewStepper(2, 10, 0.105, 3.01, 2.31)
	fmt.Println(err)
	_, err = xmath.NewStepper(2, 10, 0.1, 3.015, 2.31)
	fmt.Println(err)
	_, err = xmath.NewStepper(2, 10, 0.1, 3.01, 2.315)
	fmt.Println(err)
	_, err = xmath.NewStepper(2, 10, 0.15, 3.01, 2.31)
	fmt.Println(err)

	// Output:
	// <nil>
	// step overflow
	// max overflow
	// min overflow
	// range overflow
}

func ExampleStepper_Step() {
	s, err := xmath.NewStepper(2, 10, 0.1, 3.01, 2.31)
	if err != nil {
		panic(err)
	}
	for i := 0; i < s.Count(); i++ {
		fmt.Println(s.Step(i))
	}

	// Output:
	// 2.31 <nil>
	// 2.41 <nil>
	// 2.51 <nil>
	// 2.61 <nil>
	// 2.71 <nil>
	// 2.81 <nil>
	// 2.91 <nil>
	// 3.01 <nil>
}
