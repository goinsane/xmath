package main

import (
	"fmt"
	"github.com/orkunkaraduman/lazygo/lazymath"
	"math"
)

func main() {
	fmt.Printf("Pi is %f\n", math.Pi)
	fmt.Printf("Pi floor with precision 4: %f\n", lazymath.FloorP(math.Pi, 4))
	fmt.Printf("Pi ceil with precision 4: %f\n", lazymath.CeilP(math.Pi, 4))
	fmt.Printf("Pi round with precision 4: %f\n", lazymath.RoundP(math.Pi, 4))
	fmt.Printf("Pi round with precision 2: %f\n", lazymath.RoundP(math.Pi, 2))
	fmt.Printf("Pi round: %f\n", lazymath.Round(math.Pi))
	fmt.Printf("Square root of 2 is %f\n", math.Sqrt2)
	mn, mx := lazymath.MinMax(math.Pi, math.Sqrt2)
	fmt.Printf("Between Pi and Sqrt2: Min=%f, Max=%f\n", mn, mx)
}
