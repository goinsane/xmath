package lazymath

import "math"

func RoundP(v float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(v*k+0.5) * k
}

func Round(v float64) float64 {
	return math.Floor(v + 0.5)
}

func MinMax(x, y float64) (float64, float64) {
	return math.Min(x, y), math.Max(x, y)
}
