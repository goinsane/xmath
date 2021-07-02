package xmath

import (
	"errors"
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
	stepNumHalf   *big.Int
	stepDenomHalf *big.Int
	stepNum       *big.Int
	stepDenom     *big.Int
	maxRat        *big.Rat
	minRat        *big.Rat
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
	s.stepNumHalf, s.stepDenomHalf = s.stepRat.Num(), s.stepRat.Denom()
	s.stepNum, s.stepDenom = new(big.Int).Mul(big.NewInt(2), s.stepNumHalf), new(big.Int).Mul(big.NewInt(2), s.stepDenomHalf)
	s.intvlRat = new(big.Rat).Sub(s.maxRat, s.minRat)
	if q := new(big.Rat).Quo(s.intvlRat, s.stepRat); !q.IsInt() {
		f, _ := q.Float64()
		if r := math.Remainder(f, 1); !IsZero(r) {
			return nil, ErrStepperStepOverflow
		}
	}
	if q, r := new(big.Int).QuoRem(s.stepDenom, s.intvlRat.Denom(), new(big.Int)); r.Cmp(big.NewInt(0)) == 0 {
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
	f, exact := s.exactFromNum(n, 2)
	if !exact {
		return f, ErrStepperInexactValue
	}
	return f, nil
}

func (s *Stepper) toDiffNum(x float64) *big.Int {
	f := new(big.Rat).SetFloat64(x)
	d := new(big.Rat).Sub(f, s.minRat)
	m := new(big.Rat).Mul(d, new(big.Rat).SetInt(s.stepDenom))
	n, _ := IntBigRat(m)
	_, r := new(big.Int).QuoRem(n, s.stepNum, new(big.Int))
	n.Sub(n, r)
	switch t := r.Sign(); {
	case t < 0:
		if r.Cmp(new(big.Int).Sub(big.NewInt(0), s.stepNumHalf)) < 0 {
			n.Sub(n, s.stepNum)
		}
	case t > 0:
		if r.Cmp(s.stepNumHalf) >= 0 {
			n.Add(n, s.stepNum)
		}
	}
	return n
}

func (s *Stepper) fromDiffNum(n *big.Int) *big.Rat {
	return new(big.Rat).Add(s.minRat, new(big.Rat).SetFrac(n, s.stepDenom))
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
