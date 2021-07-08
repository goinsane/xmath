package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleStepper() {
	s, err := xmath.NewStepper(0.341, 2.5, 0.0)
	if err != nil {
		panic(err)
	}
	//fmt.Println(s.Normalize(1.25))
	//return
	for i := 0; i < 100; i++ {
		x, _ := s.Step(int64(i))
		//fmt.Println(x)
		_ = x
	}

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
	fmt.Println(s.Normalize(1.376))
	fmt.Println(s.Count())

	// Output:
	// bb
}

func ExampleStepper_Normalize_cc() {
	s, err := xmath.NewStepper(0.1, 3.0, 1.0)
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
