package xmath_test

import (
	"fmt"

	"github.com/goinsane/xmath"
)

func ExampleStepper() {
	s, err := xmath.NewStepper(2, 10, 0.1, 10, 5.21)
	if err != nil {
		panic(err)
	}
	for i := -1; i < s.Count()+1; i++ {
		fmt.Println(s.Step(i))
	}
	fmt.Println(s.Normalize(4.945))

	// Output:
	// aa
}
