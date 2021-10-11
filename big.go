package xmath

import (
	"math"
	"math/big"
)

// FloorBigFloat returns the greatest integer value less than or equal to x.
// It returns nil if x is an infinity.
func FloorBigFloat(x *big.Float) *big.Int {
	n, acc := x.Int(nil)
	if acc == big.Above {
		n.Add(n, big.NewInt(-1))
	}
	return n
}

// CeilBigFloat returns the least integer value greater than or equal to x.
// It returns nil if x is an infinity.
func CeilBigFloat(x *big.Float) *big.Int {
	n, acc := x.Int(nil)
	if acc == big.Below {
		n.Add(n, big.NewInt(+1))
	}
	return n
}

// RoundBigFloat returns the nearest integer, rounding half away from zero.
// It returns nil if x is an infinity.
func RoundBigFloat(x *big.Float) *big.Int {
	z := new(big.Float).Copy(x)
	return FloorBigFloat(z.Add(z, big.NewFloat(0.5)))
}

// IntBigRat returns the result of truncating x towards zero.
// The accuracy is big.Exact if x.IsInt(); otherwise it is big.Below or big.Above.
func IntBigRat(x *big.Rat) (*big.Int, big.Accuracy) {
	n, r := new(big.Int).QuoRem(x.Num(), x.Denom(), new(big.Int))
	switch t := r.Sign(); {
	case t < 0:
		return n, big.Above
	case t > 0:
		return n, big.Below
	}
	return n, big.Exact
}

// Int64BigRat returns the integer resulting from truncating x towards zero.
// If math.MinInt64 <= x <= math.MaxInt64, the accuracy is like IntBigRat.
// The result is (math.MinInt64, big.Above) for x < math.MinInt64, and (math.MaxInt64, big.Below) for x > math.MaxInt64.
func Int64BigRat(x *big.Rat) (int64, big.Accuracy) {
	n, a := IntBigRat(x)
	if n.IsInt64() {
		return n.Int64(), a
	}
	switch t := n.Sign(); {
	case t < 0:
		return math.MinInt64, big.Above
	case t > 0:
		return math.MaxInt64, big.Below
	}
	return 0, big.Exact
}

// Uint64BigRat returns the integer resulting from truncating x towards zero.
// If 0 <= x <= math.MaxUint64, the accuracy is like IntBigRat.
// The result is (0, big.Above) for x < 0, and (math.MaxUint64, big.Below) for x > math.MaxUint64.
func Uint64BigRat(x *big.Rat) (uint64, big.Accuracy) {
	n, a := IntBigRat(x)
	if n.IsUint64() {
		return n.Uint64(), a
	}
	switch t := n.Sign(); {
	case t < 0:
		return 0, big.Above
	case t > 0:
		return math.MaxUint64, big.Below
	}
	return 0, big.Exact
}

// FloorBigRat returns the greatest integer value less than or equal to x.
func FloorBigRat(x *big.Rat) *big.Int {
	n, r := new(big.Int).QuoRem(x.Num(), x.Denom(), new(big.Int))
	if r.Sign() < 0 {
		n.Add(n, big.NewInt(-1))
	}
	return n
}

// CeilBigRat returns the least integer value greater than or equal to x.
func CeilBigRat(x *big.Rat) *big.Int {
	n, r := new(big.Int).QuoRem(x.Num(), x.Denom(), new(big.Int))
	if r.Sign() > 0 {
		n.Add(n, big.NewInt(+1))
	}
	return n
}

// RoundBigRat returns the nearest integer, rounding half away from zero.
func RoundBigRat(x *big.Rat) *big.Int {
	numHalf, denomHalf := x.Num(), x.Denom()
	n, r := new(big.Int).QuoRem(new(big.Int).Mul(big.NewInt(2), numHalf), new(big.Int).Mul(big.NewInt(2), denomHalf), new(big.Int))
	switch t := r.Sign(); {
	case t < 0:
		if r.Cmp(new(big.Int).Sub(big.NewInt(0), denomHalf)) < 0 {
			n.Add(n, big.NewInt(-1))
		}
	case t > 0:
		if r.Cmp(denomHalf) >= 0 {
			n.Add(n, big.NewInt(+1))
		}
	}
	return n
}
