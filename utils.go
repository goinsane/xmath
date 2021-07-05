package xmath

import (
	"math/big"
)

func splitBigRat(a *big.Rat) (num, denom *big.Int) {
	/*num, denom = new(big.Int).Mul(big.NewInt(2), a.Num()), new(big.Int).Mul(big.NewInt(2), a.Denom())
	absNum := new(big.Int).Abs(num)
	neg := num.Cmp(absNum) != 0
	n, rem := new(big.Int).QuoRem(absNum, denom, new(big.Int))
	if rem.Cmp(big.NewInt(0)) != 0 {
		roundToStepBigInt(denom, rem)
	}
	num.Mul(n, denom)
	num.Add(num, rem)
	if neg {
		num.Sub(big.NewInt(0), num)
	}
	return*/
	num, denom = new(big.Int).Mul(big.NewInt(2), a.Num()), new(big.Int).Mul(big.NewInt(2), a.Denom())
	n, rem := new(big.Int).DivMod(num, denom, new(big.Int))
	if rem.Cmp(big.NewInt(0)) != 0 {
		roundToStepBigInt(denom, rem)
	}
	num.Mul(n, denom)
	num.Add(num, rem)
	//a = new(big.Rat).SetFrac(num, denom)
	//num, denom = new(big.Int).Mul(big.NewInt(2), a.Num()), new(big.Int).Mul(big.NewInt(2), a.Denom())
	return
}

func normalizeNumDenom(num, denom *big.Int) {
	num, denom = new(big.Int).Mul(big.NewInt(2), num), new(big.Int).Mul(big.NewInt(2), denom)
	n, rem := new(big.Int).DivMod(num, denom, new(big.Int))
	if rem.Cmp(big.NewInt(0)) != 0 {
		roundToStepBigInt(denom, rem)
	}
	num.Mul(n, denom)
	num.Add(num, rem)
}

func normalizeBigRat(a *big.Rat) *big.Rat {
	num, denom := splitBigRat(a)
	a.SetFrac(num, denom)
	return a
}

func roundToStepBigInt(n *big.Int, step *big.Int) *big.Int {
	if step.Cmp(big.NewInt(0)) < 0 {
		panic("step below zero")
	}
	stepHalf := new(big.Int).Quo(step, big.NewInt(2))
	r := new(big.Int).Rem(n, step)
	n.Sub(n, r)
	switch t := r.Sign(); {
	case t < 0:
		if r.Cmp(new(big.Int).Sub(big.NewInt(0), stepHalf)) < 0 {
			n.Sub(n, step)
		}
	case t > 0:
		if r.Cmp(new(big.Int).Add(stepHalf, big.NewInt(0))) >= 0 {
			n.Add(n, step)
		}
	}
	return n
}
