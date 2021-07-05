package xmath

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	ErrStepperStepZero        = errors.New("step zero")
	ErrStepperStepBelowZero   = errors.New("step below zero")
	ErrStepperUnorderedMaxMin = errors.New("unordered max min")
	ErrStepperMaxOverflow     = errors.New("max overflow")
	ErrStepperMinOverflow     = errors.New("min overflow")
	ErrStepperStepOverflow    = errors.New("step overflow")
	ErrStepperMaxExceeded     = errors.New("max exceeded")
	ErrStepperMinExceeded     = errors.New("min exceeded")
	ErrStepperInexactValue    = errors.New("inexact value")
)

type Stepper struct {
	stepRat       *big.Rat
	stepNum       *big.Int
	stepDenom     *big.Int
	maxRat        *big.Rat
	maxNum        *big.Int
	minRat        *big.Rat
	minNum        *big.Int
	intvlRat      *big.Rat
	intvlNum      *big.Int
}

func NewStepper(step, max, min float64) (s *Stepper, err error) {
	if step == 0 {
		return nil, ErrStepperStepZero
	}
	if step < 0 {
		return nil, ErrStepperStepBelowZero
	}
	if max < min {
		return nil, ErrStepperUnorderedMaxMin
	}
	if math.Nextafter(max, math.Inf(+1))-max > step {
		return nil, ErrStepperMaxOverflow
	}
	if min-math.Nextafter(min, math.Inf(-1)) > step {
		return nil, ErrStepperMinOverflow
	}
	s = &Stepper{
		stepRat: new(big.Rat).SetFloat64(step),
		maxRat:  new(big.Rat).SetFloat64(max),
		minRat:  new(big.Rat).SetFloat64(min),
	}
	s.stepNum, s.stepDenom = new(big.Int).Mul(big.NewInt(2), s.stepRat.Num()), new(big.Int).Mul(big.NewInt(2), s.stepRat.Denom())
	//s.stepNum, s.stepDenom = splitBigRat(s.stepRat)
	//s.stepNum, s.stepDenom = s.stepRat.Num(), s.stepRat.Denom()

	if q, r := new(big.Int).QuoRem(s.stepDenom, s.maxRat.Denom(), new(big.Int)); r.Cmp(big.NewInt(0)) == 0 {
		s.maxNum = new(big.Int).Mul(s.maxRat.Num(), q)
		//s.minNum = new(big.Int).Mul(s.minRat.Num(), q)
	} else {
		return nil, ErrStepperMaxOverflow
	}
	if q, r := new(big.Int).QuoRem(s.stepDenom, s.minRat.Denom(), new(big.Int)); r.Cmp(big.NewInt(0)) == 0 {
		//s.maxNum = new(big.Int).Mul(s.maxRat.Num(), q)
		s.minNum = new(big.Int).Mul(s.minRat.Num(), q)
	} else {
		return nil, ErrStepperMinOverflow
	}


	s.intvlRat = new(big.Rat).Sub(s.maxRat, s.minRat)
	if q := new(big.Rat).Quo(s.intvlRat, s.stepRat); !q.IsInt() {
		f, _ := q.Float64()
		if r := math.Remainder(f, 1); !IsZero(r) {
			return nil, ErrStepperStepOverflow
		}
	}
	if q, r := new(big.Int).QuoRem(s.stepDenom, s.intvlRat.Denom(), new(big.Int)); r.Cmp(big.NewInt(0)) == 0 {
		//s.maxNum = new(big.Int).Mul(s.maxRat.Num(), q)
		//s.minNum = new(big.Int).Mul(s.minRat.Num(), q)
		s.intvlNum = new(big.Int).Mul(s.intvlRat.Num(), q)
	} else {
		return nil, ErrStepperStepOverflow
	}
	return s, nil
}

func (s *Stepper) Count() int64 {
	return new(big.Int).Quo(s.intvlNum, s.stepNum).Int64()
}

func (s *Stepper) Step(index int64) (float64, error) {
	return s.normalize(new(big.Int).Mul(big.NewInt(index), s.stepNum))
}

func (s *Stepper) Normalize(x float64) (float64, error) {
	return s.normalize(s.toDiffNum(x))
}

func (s *Stepper) normalize(n *big.Int) (float64, error) {
	if n.Cmp(s.intvlNum) > 0 {
		f, _ := s.maxRat.Float64()
		return f, ErrStepperMaxExceeded
	}
	if n.Cmp(big.NewInt(0)) < 0 {
		f, _ := s.minRat.Float64()
		return f, ErrStepperMinExceeded
	}
	f, exact := s.exactFromNum(n, 0)
	if !exact {
		return f, ErrStepperInexactValue
	}
	return f, nil
}

func (s *Stepper) toDiffNum(x float64) *big.Int {
	/*a := new(big.Rat).SetFloat64(x)
	k := new(big.Int).Mul(a.Denom(), big.NewInt(2))
	//k := big.NewInt(1)
	m := new(big.Rat).Mul(a, new(big.Rat).SetInt(new(big.Int).Mul(s.stepDenom, k)))
	n := RoundBigRat(m)
	n.Sub(n, new(big.Int).Mul(s.minNum, k))
	//roundToStepBigInt(n, k)
	n.Quo(n, k)
	roundToStepBigInt(n, s.stepNum)
	fmt.Println(n, s.stepDenom)*/

	b := new(big.Rat).SetFloat64(x)
	//a := new(big.Rat).Sub(b, s.minRat)
	//fmt.Println(a.Float64())
	a := b
	//fmt.Println(a)
	num, denom := new(big.Int).Mul(a.Num(), big.NewInt(2)), new(big.Int).Mul(a.Denom(), big.NewInt(2))
	//num, denom := splitBigRat(a)
	//fmt.Println(num, denom)
	//num, denom := a.Num(), a.Denom()
	//fmt.Println(num, denom, a)
	newNum, newDenom := new(big.Int).Mul(num, s.stepDenom), new(big.Int).Mul(denom, s.stepDenom)
	_, _ = newNum, newDenom
	k := new(big.Int).Set(newNum)
	//fmt.Println(num, denom)
	//roundToStepBigInt(k, denom)
	//roundToStepBigInt(k, s.stepNum)
	//fmt.Println(k)
	roundToStepBigInt(k, new(big.Int).Mul(s.stepNum, denom))
	//fmt.Println(k)
	k.Quo(k, denom)
	//k.Quo(k, s.stepDenom)
	k.Sub(k, s.minNum)
	fmt.Println(k, s.stepDenom)
	roundToStepBigInt(k, s.stepNum)
	return k
}

func (s *Stepper) fromDiffNum(n *big.Int) *big.Rat {
	//fmt.Println(new(big.Int).Add(s.minNum, n), s.stepNum)
	//k := new(big.Rat).SetFrac(new(big.Int).Add(s.minNum, n), s.stepDenom)
	//normalizeBigRat(k)
	//return k
	roundToStepBigInt(n, s.stepNum)
	num, denom := new(big.Int).Add(s.minNum, n), new(big.Int).Set(s.stepDenom)
	fmt.Println(num, denom)
	//normalizeNumDenom(num, denom)
	//fmt.Println(num, denom)
	//num, denom = splitBigRat(new(big.Rat).SetFrac(num, denom))
	return new(big.Rat).SetFrac(num, denom)
}

func (s *Stepper) exactFromNum(n *big.Int, iterCount int) (f float64, exact bool) {
	g, exact := s.fromDiffNum(n).Float64()
	if exact {
		f = g
		return
	}
	for i := 2 * 1; i < 2*(iterCount+1); i++ {
		j := i / 2
		if i%2 >= 1 {
			j = -j
		}
		k := new(big.Int).Add(n, big.NewInt(int64(j)))
		f, exact = s.fromDiffNum(k).Float64()
		if exact {
			return
		}
	}
	return g, false
}
