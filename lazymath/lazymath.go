package lazymath

import "math"

func FloorP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k) / k
}

func CeilP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Ceil(x*k) / k
}

func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

/*func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return Round(x*k) / k
}*/
func RoundP(x float64, p int) float64 {
	k := math.Pow10(p)
	return math.Floor(x*k+0.5) / k
}

func MinMax(x, y float64) (float64, float64) {
	return math.Min(x, y), math.Max(x, y)
}
