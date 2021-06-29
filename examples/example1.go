// +build ignore

package main

import (
	"fmt"
	"math"

	"github.com/goinsane/xmath"
)

func main() {
	fmt.Printf("Pi is %v\n", math.Pi)
	fmt.Printf("Square root of 2 (Sqrt2) is %v\n", math.Sqrt2)

	fmt.Printf("floor of Pi with precision 4: %v\n", xmath.FloorP(math.Pi, 4))

	fmt.Printf("ceil of Pi with precision 4: %v\n", xmath.CeilP(math.Pi, 4))

	fmt.Printf("round of Pi with precision 4: %v\n", xmath.RoundP(math.Pi, 4))
	fmt.Printf("round of Pi with precision 2: %v\n", xmath.RoundP(math.Pi, 2))

	fmt.Printf("round of Pi: %v\n", xmath.Round(math.Pi))

	fmt.Printf("max of Pi and Sqrt2 is %v\n", xmath.Max(math.Pi, math.Sqrt2))

	fmt.Printf("min of Pi and Sqrt2 is %v\n", xmath.Min(math.Pi, math.Sqrt2))

	max, min := xmath.MaxMin(math.Pi, math.Sqrt2)
	fmt.Printf("Pi and Sqrt2: max=%v  min=%v\n", max, min)

	fmt.Printf("is 2.4 between 3 and 2.4: %t\n", xmath.Between(2.4, 3, 2.4))
	fmt.Printf("is 2.6 between 3 and 2.4: %t\n", xmath.Between(2.6, 3, 2.4))
	fmt.Printf("is 3 between 3 and 2.4: %t\n", xmath.Between(3, 3, 2.4))

	fmt.Printf("is 2.4 between in 3 and 2.4: %t\n", xmath.BetweenIn(2.4, 3, 2.4))
	fmt.Printf("is 2.6 between in 3 and 2.4: %t\n", xmath.BetweenIn(2.6, 3, 2.4))
	fmt.Printf("is 3 between in 3 and 2.4: %t\n", xmath.BetweenIn(3, 3, 2.4))

	// Output:
	// Pi is 3.141592653589793
	// Square root of 2 (Sqrt2) is 1.4142135623730951
	// floor of Pi with precision 4: 3.1415
	// ceil of Pi with precision 4: 3.1416
	// round of Pi with precision 4: 3.1416
	// round of Pi with precision 2: 3.14
	// round of Pi: 3
	// max of Pi and Sqrt2 is 3.141592653589793
	// min of Pi and Sqrt2 is 1.4142135623730951
	// Pi and Sqrt2: max=3.141592653589793  min=1.4142135623730951
	// is 2.4 between 3 and 2.4: false
	// is 2.6 between 3 and 2.4: true
	// is 3 between 3 and 2.4: false
	// is 2.4 between in 3 and 2.4: true
	// is 2.6 between in 3 and 2.4: true
	// is 3 between in 3 and 2.4: true
}
