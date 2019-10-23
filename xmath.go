// Package xmath provides some extended capabilities according to GoLang's math.
package xmath

import "math"

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

// Round returns the integer value rounded by x.
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundP returns the value rounded by x with specified precision.
func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k+0.5) / k
}

/*func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return Round(x*k) / k
}*/

// MinMax returns the min, max values in this order.
func MinMax(x, y float64) (float64, float64) {
	return math.Min(x, y), math.Max(x, y)
}

// Between checks x is between a and b
func Between(x float64, a, b float64) bool {
	min, max := math.Min(a, b), math.Max(a, b)
	return min < a && a < max
}

// BetweenIn checks x is in a and b
func BetweenIn(x float64, a, b float64) bool {
	min, max := math.Min(a, b), math.Max(a, b)
	return min <= a && a <= max
}
