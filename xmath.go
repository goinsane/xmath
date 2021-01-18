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

// MinMax returns the min, max values in this order.
func MinMax(x, y float64) (min float64, max float64) {
	if x < y {
		return x, y
	} else {
		return y, x
	}
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
//	SafeDiv(0, ±n) = ±0
//	SafeDiv(±n, 0) = ±Inf
//	SafeDiv(0, 0) = NaN
func SafeDiv(x, y float64) float64 {
	if y == 0 {
		if x == 0 {
			return math.NaN()
		}
		if x < 0 {
			return math.Inf(-1)
		}
		return math.Inf(+1)
	}
	return x / y
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

// CryptoRandFloat returns a random decimal number in [0, 1).
// It returns -1 when error occurs.
func CryptoRandFloat() float64 {
	r := CryptoRandInt(math.MaxInt64)
	if r < 0 {
		return -1
	}
	return float64(r) / math.MaxInt64
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
