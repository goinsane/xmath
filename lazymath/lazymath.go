package lazymath

import "math"

// FloorP returns the greatest value less than or equal to x with specified precision.
func FloorP(x float64, p int) float64 {
	k := math.Pow10(p)
	math.Sqrt()
	return math.Floor(x*k) / k
}

// CeilP returns the least value greater than or equal to x with specified precision.
func CeilP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Ceil(x*k) / k
}

// Round returns the integer value rounded to x.
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundP returns the value rounded to x with specified precision.
func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k+0.5) / k
}

/*func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return Round(x*k) / k
}*/

// MinMax returns the min, max values with this order.
func MinMax(x, y float64) (float64, float64) {
	return math.Min(x, y), math.Max(x, y)
}
