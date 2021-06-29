// Package xmath provides some extended capabilities according to GoLang's math.
package xmath

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

// FloorP returns the greatest value less than or equal to x with specified precision.
func FloorP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k) / k
}

// CeilP returns the least value greater than or equal to x with specified precision.
func CeilP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Ceil(x*k) / k
}

// Round returns the nearest integer value, rounding half away from zero.
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundP returns the nearest integer value, rounding half away from zero with specified precision.
func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k+0.5) / k
}

// Max returns the larger of x...
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
//	Max(x) = x
//	Max() = +Inf
func Max(x ...float64) float64 {
	if len(x) <= 0 {
		return math.Inf(+1)
	}
	result := math.Inf(-1)
	for _, a := range x {
		result = math.Max(a, result)
	}
	return result
}

// Min returns the smaller of x...
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
//	Min(x) = x
//	Min() = -Inf
func Min(x ...float64) float64 {
	if len(x) <= 0 {
		return math.Inf(-1)
	}
	result := math.Inf(+1)
	for _, a := range x {
		result = math.Min(a, result)
	}
	return result
}

// MaxMin returns the max, min values in this order, similar with Max and Min functions.
//
// Special cases are:
//	MaxMin(x) = x, x
//	MaxMin() = +Inf, -Inf
func MaxMin(x ...float64) (max float64, min float64) {
	min, max = MinMax(x...)
	return
}

// MinMax returns the min, max values in this order, similar with Min and Max functions.
//
// Special cases are:
//	MinMax(x) = x, x
//	MinMax() = -Inf, +Inf
func MinMax(x ...float64) (min float64, max float64) {
	if len(x) <= 0 {
		return math.Inf(-1), math.Inf(+1)
	}
	min = math.Inf(+1)
	max = math.Inf(-1)
	for _, a := range x {
		min = math.Min(a, min)
		max = math.Max(a, max)
	}
	return
}

// MaxInt returns the larger integer of x...
//
// Special cases are:
//	MaxInt(x) = x
//	MaxInt() = math.MaxInt64
func MaxInt(x ...int64) int64 {
	if len(x) <= 0 {
		return int64(math.MaxInt64)
	}
	result := int64(math.MinInt64)
	for _, a := range x {
		if a > result {
			result = a
		}
	}
	return result
}

// MinInt returns the smaller integer of x...
//
// Special cases are:
//	MinInt(x) = x
//	MinInt() = math.MinInt64
func MinInt(x ...int64) int64 {
	if len(x) <= 0 {
		return int64(math.MinInt64)
	}
	result := int64(math.MaxInt64)
	for _, a := range x {
		if a < result {
			result = a
		}
	}
	return result
}

// MaxMinInt returns the max, min integers in this order, similar with MaxInt and MinInt functions.
//
// Special cases are:
//	MaxMinInt(x) = x, x
//	MaxMinInt() = math.MaxInt64, math.MinInt64
func MaxMinInt(x ...int64) (max int64, min int64) {
	min, max = MinMaxInt(x...)
	return
}

// MinMaxInt returns the min, max integers in this order, similar with MinInt and MaxInt functions.
//
// Special cases are:
//	MinMaxInt(x) = x, x
//	MinMaxInt() = math.MinInt64, math.MaxInt64
func MinMaxInt(x ...int64) (min int64, max int64) {
	if len(x) <= 0 {
		return int64(math.MinInt64), int64(math.MaxInt64)
	}
	min = int64(math.MaxInt64)
	max = int64(math.MinInt64)
	for _, a := range x {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return
}

// Between checks x is between a and b
func Between(x float64, a, b float64) bool {
	min, max := MinMax(a, b)
	return min < x && x < max
}

// BetweenIn checks x is in a and b
func BetweenIn(x float64, a, b float64) bool {
	min, max := MinMax(a, b)
	return min <= x && x <= max
}

// SafeDiv divides x to y without 'division by zero' error.
// Using this function is not necessary in GoLang. Because GoLang's default behaviour is same with this function.
// Special cases are:
//	SafeDiv(0, ±n, true) = ±0
//	SafeDiv(0, ±n, false) = ±0
//	SafeDiv(±n, 0, true) = ±Inf
//	SafeDiv(±n, 0, false) = ±Inf
//	SafeDiv(0, 0, true) = NaN
//	SafeDiv(0, 0, false) = 0
func SafeDiv(x, y float64, allowNaN bool) float64 {
	if y == 0 {
		if x < 0 {
			return math.Inf(-1)
		}
		if x > 0 {
			return math.Inf(+1)
		}
		if allowNaN {
			return math.NaN()
		}
		return 0
	}
	return x / y
}

// CryptoRand returns a random decimal number in [0, 1).
// It returns -1 when error occurs.
func CryptoRand() float64 {
	r := CryptoRandInt(math.MaxInt64)
	if r < 0 {
		return -1
	}
	return float64(r) / math.MaxInt64
}

// CryptoRandInt returns a random integer in [0, max).
// It returns -1 when error occurs.
func CryptoRandInt(max int64) int64 {
	if max <= 0 {
		return -1
	}
	num, _ := crand.Int(crand.Reader, big.NewInt(max))
	if num == nil {
		return -1
	}
	return num.Int64()
}

// CryptoRandFloat is synonym with CryptoRand.
func CryptoRandFloat() float64 {
	return CryptoRand()
}

// CryptoRandCode generates random code in [10^(n-1), 10^n).
// It returns -1 when error occurs.
func CryptoRandCode(n int) int64 {
	if n < 1 || n > 18 {
		return -1
	}
	start := int64(1)
	for i := 0; i < n-1; i++ {
		start *= 10
	}
	r := CryptoRandInt(start*10 - start)
	if r < 0 {
		return -1
	}
	return start + r
}

// AlmostEqual64 checks almost equality of all given 64-bit floating points values.
// It returns true if all values are equal.
//
// Special cases are:
//	AlmostEqual64() = false
//	AlmostEqual64(x) = true
//	AlmostEqual64(NaN) = false
//	AlmostEqual64(NaN, x) = false
//	AlmostEqual64(x, NaN) = false
//	AlmostEqual64(+Inf, +Inf) = true
//	AlmostEqual64(-Inf, -Inf) = true
func AlmostEqual64(x ...float64) bool {
	if len(x) <= 0 {
		return false
	}
	var a float64
	for i, b := range x {
		if math.IsNaN(b) {
			return false
		}
		if i > 0 {
			if d := math.Float64bits(a) - math.Float64bits(b); d > 1 && -d > 1 {
				return false
			}
		}
		a = b
	}
	return true
}

// AlmostEqual32 checks almost equality of all given 32-bit floating points values.
// It returns true if all values are equal.
//
// Special cases are:
//	AlmostEqual32() = false
//	AlmostEqual32(x) = true
//	AlmostEqual32(NaN) = false
//	AlmostEqual32(NaN, x) = false
//	AlmostEqual32(x, NaN) = false
//	AlmostEqual32(+Inf, +Inf) = true
//	AlmostEqual32(-Inf, -Inf) = true
func AlmostEqual32(x ...float32) bool {
	if len(x) <= 0 {
		return false
	}
	var a float32
	for i, b := range x {
		if math.IsNaN(float64(b)) {
			return false
		}
		if i > 0 {
			if d := math.Float32bits(a) - math.Float32bits(b); d > 1 && -d > 1 {
				return false
			}
		}
		a = b
	}
	return true
}

// AlmostEqual is synonym with AlmostEqual64.
func AlmostEqual(x ...float64) bool {
	return AlmostEqual64(x...)
}

// Equal64 checks exact equality of all given 64-bit floating points values by comparing.
// It returns true if all values are equal.
//
// Special cases are:
//	Equal64() = false
//	Equal64(x) = true
//	Equal64(NaN) = false
//	Equal64(NaN, x) = false
//	Equal64(x, NaN) = false
//	Equal64(+Inf, +Inf) = true
//	Equal64(-Inf, -Inf) = true
func Equal64(x ...float64) bool {
	if len(x) <= 0 {
		return false
	}
	var a float64
	for i, b := range x {
		if math.IsNaN(b) {
			return false
		}
		if i > 0 {
			if !math.IsNaN(a-b) && (a > b || a < b) {
				return false
			}
		}
		a = b
	}
	return true
}

// Equal32 checks exact equality of all given 32-bit floating points values by comparing.
// It returns true if all values are equal.
//
// Special cases are:
//	Equal32() = false
//	Equal32(x) = true
//	Equal32(NaN) = false
//	Equal32(NaN, x) = false
//	Equal32(x, NaN) = false
//	Equal32(+Inf, +Inf) = true
//	Equal32(-Inf, -Inf) = true
func Equal32(x ...float32) bool {
	if len(x) <= 0 {
		return false
	}
	var a float32
	for i, b := range x {
		if math.IsNaN(float64(b)) {
			return false
		}
		if i > 0 {
			if !math.IsNaN(float64(a-b)) && (a > b || a < b) {
				return false
			}
		}
		a = b
	}
	return true
}

// Equal is synonym with Equal64.
func Equal(x ...float64) bool {
	return Equal64(x...)
}
