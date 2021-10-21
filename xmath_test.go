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

func ExampleZero() {
	var sign int

	sign = -1
	fmt.Printf("zero with sign %+d: %+.2f\n", sign, xmath.Zero(sign))

	sign = 0
	fmt.Printf("zero with sign %+d: %+.2f\n", sign, xmath.Zero(sign))

	sign = 1
	fmt.Printf("zero with sign %+d: %+.2f\n", sign, xmath.Zero(sign))

	// Output:
	// zero with sign -1: -0.00
	// zero with sign +0: +0.00
	// zero with sign +1: +0.00
}

func ExampleSign() {
	var x float64

	x = -1.3
	fmt.Printf("sign of %+.2f: %+d\n", x, xmath.Sign(x))

	x = xmath.Zero(-1)
	fmt.Printf("sign of %+.2f: %+d\n", x, xmath.Sign(x))

	x = xmath.Zero(+1)
	fmt.Printf("sign of %+.2f: %+d\n", x, xmath.Sign(x))

	x = 2.7
	fmt.Printf("sign of %+.2f: %+d\n", x, xmath.Sign(x))

	// Output:
	// sign of -1.30: -1
	// sign of -0.00: +0
	// sign of +0.00: +0
	// sign of +2.70: +1
}

func ExampleSignInt() {
	var x int64

	x = -3
	fmt.Printf("sign of integer %+d: %+d\n", x, xmath.SignInt(x))

	x = 0
	fmt.Printf("sign of integer %+d: %+d\n", x, xmath.SignInt(x))

	x = 7
	fmt.Printf("sign of integer %+d: %+d\n", x, xmath.SignInt(x))

	// Output:
	// sign of integer -3: -1
	// sign of integer +0: +0
	// sign of integer +7: +1
}

func ExampleSum() {
	x := []float64{-5.6, 2.1, 4.5, -10.9, -3.4}
	fmt.Printf("sum of %v: %v\n", x, xmath.Sum(x...))

	// Output:
	// sum of [-5.6 2.1 4.5 -10.9 -3.4]: -13.3
}

func ExampleAvg() {
	x := []float64{-5.6, 2.1, 4.5, -10.9, -3.4}
	fmt.Printf("avg of %v: %v\n", x, xmath.Avg(x...))

	// Output:
	// avg of [-5.6 2.1 4.5 -10.9 -3.4]: -2.66
}

func ExampleSumInt() {
	x := []int64{-5, 2, 4, -10, -3}
	fmt.Printf("sum of integers %v: %v\n", x, xmath.SumInt(x...))

	// Output:
	// sum of integers [-5 2 4 -10 -3]: -12
}

func ExampleAvgInt() {
	x := []int64{-5, 2, 4, -10, -3}
	fmt.Printf("avg of integers %v: %v\n", x, xmath.AvgInt(x...))

	// Output:
	// avg of integers [-5 2 4 -10 -3]: -2.4
}

func ExampleSumUint() {
	x := []uint64{5, 2, 4, 10, 3}
	fmt.Printf("sum of unsigned integers %v: %v\n", x, xmath.SumUint(x...))

	// Output:
	// sum of unsigned integers [5 2 4 10 3]: 24
}

func ExampleAvgUint() {
	x := []uint64{5, 2, 4, 10, 3}
	fmt.Printf("avg of unsigned integers %v: %v\n", x, xmath.AvgUint(x...))

	// Output:
	// avg of unsigned integers [5 2 4 10 3]: 4.8
}

func ExampleSumInt2() {
	x := []int64{-5, 2, 4, -10, -3}
	r, o := xmath.SumInt2(x...)
	fmt.Printf("sum of integers %v: %v overflow %v\n", x, r, o)

	x = []int64{5, math.MaxInt64, -2}
	r, o = xmath.SumInt2(x...)
	fmt.Printf("sum of integers %v: %v overflow %v\n", x, r, o)

	x = []int64{-7, math.MinInt64, 1}
	r, o = xmath.SumInt2(x...)
	fmt.Printf("sum of integers %v: %v overflow %v\n", x, r, o)

	// Output:
	// sum of integers [-5 2 4 -10 -3]: -12 overflow false
	// sum of integers [5 9223372036854775807 -2]: -9223372036854775806 overflow true
	// sum of integers [-7 -9223372036854775808 1]: 9223372036854775802 overflow true
}

func ExampleAvgInt2() {
	x := []int64{-5, 2, 4, -10, -3}
	r, o := xmath.AvgInt2(x...)
	fmt.Printf("avg of integers %v: %v overflow %v\n", x, r, o)

	x = []int64{5, math.MaxInt64, -2}
	r, o = xmath.AvgInt2(x...)
	fmt.Printf("avg of integers %v: %v overflow %v\n", x, r, o)

	x = []int64{-7, math.MinInt64, 1}
	r, o = xmath.AvgInt2(x...)
	fmt.Printf("avg of integers %v: %v overflow %v\n", x, r, o)

	// Output:
	// avg of integers [-5 2 4 -10 -3]: -2 overflow false
	// avg of integers [5 9223372036854775807 -2]: -3074457345618258602 overflow true
	// avg of integers [-7 -9223372036854775808 1]: 3074457345618258600 overflow true
}

func ExampleSumUint2() {
	x := []uint64{5, 2, 4, 10, 3}
	r, o := xmath.SumUint2(x...)
	fmt.Printf("sum of unsigned integers %v: %v overflow %v\n", x, r, o)

	x = []uint64{5, math.MaxUint64}
	r, o = xmath.SumUint2(x...)
	fmt.Printf("sum of unsigned integers %v: %v overflow %v\n", x, r, o)

	// Output:
	// sum of unsigned integers [5 2 4 10 3]: 24 overflow false
	// sum of unsigned integers [5 18446744073709551615]: 4 overflow true
}

func ExampleAvgUint2() {
	x := []uint64{5, 2, 4, 10, 3}
	r, o := xmath.AvgUint2(x...)
	fmt.Printf("avg of unsigned integers %v: %v overflow %v\n", x, r, o)

	x = []uint64{5, math.MaxUint64}
	r, o = xmath.AvgUint2(x...)
	fmt.Printf("avg of unsigned integers %v: %v overflow %v\n", x, r, o)

	// Output:
	// avg of unsigned integers [5 2 4 10 3]: 4 overflow false
	// avg of unsigned integers [5 18446744073709551615]: 2 overflow true
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
