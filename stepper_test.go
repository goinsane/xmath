package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleStepper() {
	s, err := xmath.NewStepper(0.33, 2.25, 0.60)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Normalize(1.5))

	// Output:
	// aa
}

func ExampleStepper_Normalize_aa() {
	s, err := xmath.NewStepper(0.66, 2.7, -0.6)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Normalize(1.1))
	fmt.Println(s.Normalize(1.124))
	fmt.Println(s.Normalize(1.125))
	fmt.Println(s.Normalize(1.126))
	fmt.Println(s.Normalize(1.2))
	fmt.Println(s.Normalize(1.374))
	fmt.Println(s.Normalize(1.375))
	fmt.Println(s.Count())

	// Output:
	// aa
}

func ExampleStepper_Normalize_bb() {
	s, err := xmath.NewStepper(0.25, 3.0, 1.0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Normalize(1.1))
	fmt.Println(s.Normalize(1.124))
	fmt.Println(s.Normalize(1.125))
	fmt.Println(s.Normalize(1.126))
	fmt.Println(s.Normalize(1.2))
	fmt.Println(s.Normalize(1.374))
	fmt.Println(s.Normalize(1.375))
	fmt.Println(s.Count())

	// Output:
	// bb
}

func ExampleStepper_Normalize_cc() {
	s, err := xmath.NewStepper(0.5, 3.0, 1.0)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.Normalize(1.24))
	fmt.Println(s.Normalize(1.25))
	fmt.Println(s.Normalize(1.26))
	fmt.Println(s.Normalize(1.74))
	fmt.Println(s.Normalize(1.75))
	fmt.Println(s.Count())
	fmt.Println(s.Step(5))

	// Output:
	// cc
}
