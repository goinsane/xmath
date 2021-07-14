package xmath_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/goinsane/xmath"
)

func BenchmarkRound(b *testing.B) {
	if b.N > 1000000 {
		b.N = 1000000
	}
	numberSet := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		numberSet[i] = xmath.CryptoRandFloat()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xmath.Round(numberSet[i])
	}
}

func BenchmarkCryptoRandInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xmath.CryptoRandInt(math.MaxInt64)
	}
}

func BenchmarkCryptoRandFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xmath.CryptoRandFloat()
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xmath.CryptoRand()
	}
}

func BenchmarkCryptoRandCode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xmath.CryptoRandCode(8)
	}
}

func ExampleFloorP() {
	fmt.Printf("floor of +3.1416 with precision 3: %+6.4f\n", xmath.FloorP(+3.1416, 3))
	fmt.Printf("floor of -3.1416 with precision 2: %+6.4f\n", xmath.FloorP(-3.1416, 2))

	// Output:
	// floor of +3.1416 with precision 3: +3.1410
	// floor of -3.1416 with precision 2: -3.1500
}

func ExampleCeilP() {
	fmt.Printf("ceil of +3.1416 with precision 3: %+6.4f\n", xmath.CeilP(+3.1416, 3))
	fmt.Printf("ceil of -3.1416 with precision 2: %+6.4f\n", xmath.CeilP(-3.1416, 2))

	// Output:
	// ceil of +3.1416 with precision 3: +3.1420
	// ceil of -3.1416 with precision 2: -3.1400
}

func ExampleRoundP() {
	fmt.Printf("round of +3.1416 with precision 2: %+6.4f\n", xmath.RoundP(+3.1416, 2))
	fmt.Printf("round of -3.1416 with precision 3: %+6.4f\n", xmath.RoundP(-3.1416, 3))

	// Output:
	// round of +3.1416 with precision 2: +3.1400
	// round of -3.1416 with precision 3: -3.1420
}

func ExampleMaxMin() {
	list := []float64{-1.215, -1.4142, +3.1416}
	max, min := xmath.MaxMin(list...)
	fmt.Printf("max=%+6.4f and min=%+6.4f of %+6.4f\n", max, min, list)

	// Output:
	// max=+3.1416 and min=-1.4142 of [-1.2150 -1.4142 +3.1416]
}

func ExampleMaxMinInt() {
	list := []int64{-1, -8, +7, -3, +5}
	max, min := xmath.MaxMinInt(list...)
	fmt.Printf("max=%+d and min=%+d of %+d\n", max, min, list)

	// Output:
	// max=+7 and min=-8 of [-1 -8 +7 -3 +5]
}

func ExampleMaxMinUint() {
	list := []uint64{+1, +8, +7, +3, +5}
	max, min := xmath.MaxMinUint(list...)
	fmt.Printf("max=%+d and min=%+d of %+d\n", max, min, list)

	// Output:
	// max=+8 and min=+1 of [+1 +8 +7 +3 +5]
}

func ExampleBetween() {
	var x float64
	interval := []float64{+3.1416, +1.4142}
	x = +1.4141
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))
	x = +1.4142
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))
	x = +1.4143
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))
	x = +3.1415
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))
	x = +3.1416
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))
	x = +3.1417
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.Between(x, interval[0], interval[1]))

	// Output:
	// is +1.4141 between [+3.1416 +1.4142]: false
	// is +1.4142 between [+3.1416 +1.4142]: false
	// is +1.4143 between [+3.1416 +1.4142]: true
	// is +3.1415 between [+3.1416 +1.4142]: true
	// is +3.1416 between [+3.1416 +1.4142]: false
	// is +3.1417 between [+3.1416 +1.4142]: false
}

func ExampleBetweenIn() {
	var x float64
	interval := []float64{+3.1416, +1.4142}
	x = +1.4141
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))
	x = +1.4142
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))
	x = +1.4143
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))
	x = +3.1415
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))
	x = +3.1416
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))
	x = +3.1417
	fmt.Printf("is %+6.4f between %+6.4f: %t\n", x, interval, xmath.BetweenIn(x, interval[0], interval[1]))

	// Output:
	// is +1.4141 between [+3.1416 +1.4142]: false
	// is +1.4142 between [+3.1416 +1.4142]: true
	// is +1.4143 between [+3.1416 +1.4142]: true
	// is +3.1415 between [+3.1416 +1.4142]: true
	// is +3.1416 between [+3.1416 +1.4142]: true
	// is +3.1417 between [+3.1416 +1.4142]: false
}

func ExampleSafeDiv() {
	var x, y float64
	x, y = 0, 0
	fmt.Printf("safe divide of %+6.4f/%+6.4f with allowNaN is %+6.4f\n", x, y, xmath.SafeDiv(x, y, false))
	x, y = 0, 0
	fmt.Printf("safe divide of %+6.4f/%+6.4f without allowNaN is %+6.4f\n", x, y, xmath.SafeDiv(x, y, true))
	x, y = -5, 0
	fmt.Printf("safe divide of %+6.4f/%+6.4f is %+6.4f\n", x, y, xmath.SafeDiv(x, y, false))
	x, y = +7, 0
	fmt.Printf("safe divide of %+6.4f/%+6.4f is %+6.4f\n", x, y, xmath.SafeDiv(x, y, false))
	x, y = +11, 2
	fmt.Printf("safe divide of %+6.4f/%+6.4f is %+6.4f\n", x, y, xmath.SafeDiv(x, y, false))

	// Output:
	// safe divide of +0.0000/+0.0000 with allowNaN is +0.0000
	// safe divide of +0.0000/+0.0000 without allowNaN is   +NaN
	// safe divide of -5.0000/+0.0000 is   -Inf
	// safe divide of +7.0000/+0.0000 is   +Inf
	// safe divide of +11.0000/+2.0000 is +5.5000
}

func ExampleAlmostEqual64() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v almost equal: %t\n", x, y, xmath.AlmostEqual64(x, y))

	// Output:
	// a=1000 b=2.00000004
	// a+b+b=1004.0000000800001
	// b+b+a=1004.00000008
	// are 1004.0000000800001 and 1004.00000008 equal: false
	// are 1004.0000000800001 and 1004.00000008 almost equal: true
}

func ExampleAlmostEqual32() {
	var a, b float32
	a, b = 1000.0, 2.4
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v almost equal: %t\n", x, y, xmath.AlmostEqual32(x, y))

	// Output:
	// a=1000 b=2.4
	// a+b+b=1004.80005
	// b+b+a=1004.8
	// are 1004.80005 and 1004.8 equal: false
	// are 1004.80005 and 1004.8 almost equal: true
}

func ExampleAlmostEqual() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("is AlmostEqual synonym with AlmostEqual64: %t\n", xmath.AlmostEqual(x, y) == xmath.AlmostEqual64(x, y))

	// Output:
	// is AlmostEqual synonym with AlmostEqual64: true
}

func ExampleAlmostEqualD64() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v almost equal: %t\n", x, y, xmath.AlmostEqualD64(0.00000001, x, y))

	// Output:
	// a=1000 b=2.00000004
	// a+b+b=1004.0000000800001
	// b+b+a=1004.00000008
	// are 1004.0000000800001 and 1004.00000008 equal: false
	// are 1004.0000000800001 and 1004.00000008 almost equal: true
}

func ExampleAlmostEqualD32() {
	var a, b float32
	a, b = 1000.0, 2.4
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v almost equal: %t\n", x, y, xmath.AlmostEqualD32(0.1, x, y))

	// Output:
	// a=1000 b=2.4
	// a+b+b=1004.80005
	// b+b+a=1004.8
	// are 1004.80005 and 1004.8 equal: false
	// are 1004.80005 and 1004.8 almost equal: true
}

func ExampleAlmostEqualD() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("is AlmostEqualD synonym with AlmostEqualD64: %t\n", xmath.AlmostEqualD(x, y) == xmath.AlmostEqualD64(x, y))

	// Output:
	// is AlmostEqualD synonym with AlmostEqualD64: true
}

func ExampleEqual64() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v equal by function: %t\n", x, y, xmath.Equal64(x, y))

	// Output:
	// a=1000 b=2.00000004
	// a+b+b=1004.0000000800001
	// b+b+a=1004.00000008
	// are 1004.0000000800001 and 1004.00000008 equal: false
	// are 1004.0000000800001 and 1004.00000008 equal by function: false
}

func ExampleEqual32() {
	var a, b float32
	a, b = 1000.0, 2.4
	x, y := a+b+b, b+b+a
	fmt.Printf("a=%v b=%v\na+b+b=%v\nb+b+a=%v\n", a, b, x, y)
	fmt.Printf("are %v and %v equal: %t\n", x, y, x == y)
	fmt.Printf("are %v and %v equal by function: %t\n", x, y, xmath.Equal32(x, y))

	// Output:
	// a=1000 b=2.4
	// a+b+b=1004.80005
	// b+b+a=1004.8
	// are 1004.80005 and 1004.8 equal: false
	// are 1004.80005 and 1004.8 equal by function: false
}

func ExampleEqual() {
	var a, b float64
	a, b = 1000.0, 2.00000004
	x, y := a+b+b, b+b+a
	fmt.Printf("is Equal synonym with Equal64: %t\n", xmath.Equal(x, y) == xmath.Equal64(x, y))

	// Output:
	// is Equal synonym with Equal64: true
}

func ExampleIsZero() {
	fmt.Printf("is remainder of -2.4/0.4 zero: %t\n", xmath.IsZero(math.Remainder(-2.4, 0.4)))
	fmt.Printf("is remainder of -2.5/0.4 zero: %t\n", xmath.IsZero(math.Remainder(-2.5, 0.4)))

	// Output:
	// is remainder of -2.4/0.4 zero: true
	// is remainder of -2.5/0.4 zero: false
}

func Example_example1() {
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
