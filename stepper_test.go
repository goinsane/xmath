package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleStepper_Normalize_step0_1_max5_0_min__2_5() {
	s, err := xmath.NewStepper(0.1, 5.0, -2.0)
	if err != nil {
		panic(err)
	}
	var x, y float64
	x = 3.04
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = 3.05
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = 3.06
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = 2.16
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = -1.64
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = -1.65
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)
	x = -1.66
	y, err = s.Normalize(x)
	fmt.Printf("normalize of %v: %v err=%v\n", x, y, err)

	y, err = s.Step(45)
	fmt.Printf("%v %v\n", y, err)

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
	fmt.Println(s.Normalize(0.8))
	fmt.Println(s.Count())
	fmt.Println(s.Step(5))

	// Output:
	// cc
}
