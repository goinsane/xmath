package xmath

import (
	"math"
	"math/big"
)

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

func FloorBigRat(x *big.Rat) *big.Int {
	numHalf, denomHalf := x.Num(), x.Denom()
	n, r := new(big.Int).QuoRem(new(big.Int).Mul(big.NewInt(2), numHalf), new(big.Int).Mul(big.NewInt(2), denomHalf), new(big.Int))
	switch t := r.Sign(); {
	case t < 0:
		n.Add(n, big.NewInt(-1))
	case t > 0:
	}
	return n
}

func CeilBigRat(x *big.Rat) *big.Int {
	numHalf, denomHalf := x.Num(), x.Denom()
	n, r := new(big.Int).QuoRem(new(big.Int).Mul(big.NewInt(2), numHalf), new(big.Int).Mul(big.NewInt(2), denomHalf), new(big.Int))
	switch t := r.Sign(); {
	case t < 0:
	case t > 0:
		n.Add(n, big.NewInt(+1))
	}
	return n
}

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
