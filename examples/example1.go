// +build ignore

package main

import (
	"fmt"
	"math"

	"github.com/goinsane/xmath"
)

func main() {
	fmt.Printf("Pi is %f\n", math.Pi)
	fmt.Printf("Square root of 2 (Sqrt2) is %f\n", math.Sqrt2)
	fmt.Printf("Pi floor with precision 4: %f\n", xmath.FloorP(math.Pi, 4))
	fmt.Printf("Pi ceil with precision 4: %f\n", xmath.CeilP(math.Pi, 4))
	fmt.Printf("Pi round with precision 4: %f\n", xmath.RoundP(math.Pi, 4))
	fmt.Printf("Pi round with precision 2: %f\n", xmath.RoundP(math.Pi, 2))
	fmt.Printf("Pi round: %f\n", xmath.Round(math.Pi))
	fmt.Printf("Min of Pi and Sqrt2 is %f\n", xmath.Min(math.Pi, math.Sqrt2))
	fmt.Printf("Max of Pi and Sqrt2 is %f\n", xmath.Max(math.Pi, math.Sqrt2))
	min, max := xmath.MinMax(math.Pi, math.Sqrt2)
	fmt.Printf("Pi and Sqrt2: Min=%f, Max=%f\n", min, max)
	fmt.Printf("Is 2.4 between 3 and 2.4: %v\n", xmath.Between(2.4, 3, 2.4))
	fmt.Printf("Is 2.4 between in 3 and 2.4: %v\n", xmath.BetweenIn(2.4, 3, 2.4))
}
