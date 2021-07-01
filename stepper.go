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
	stepRat   *big.Rat
	stepNum   *big.Int
	stepDenom *big.Int
	maxRat    *big.Rat
	minRat    *big.Rat
	diffRat   *big.Rat
	diffNum   *big.Int
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
	s.stepNum, s.stepDenom = s.stepRat.Num(), s.stepRat.Denom()
	s.diffRat = new(big.Rat).Sub(s.maxRat, s.minRat)
	if q, r := new(big.Int).QuoRem(s.stepDenom, s.diffRat.Denom(), new(big.Int)); r.Cmp(big.NewInt(0)) == 0 {
		s.diffNum = new(big.Int).Mul(s.diffRat.Num(), q)
	} else {
		return nil, ErrStepperStepOverflow
	}
	return s, nil
}

func (s *Stepper) Normalize(x float64) (float64, error) {
	f := new(big.Rat).SetFloat64(x)
	d := new(big.Rat).Sub(f, s.minRat)
	m := new(big.Rat).Mul(d, new(big.Rat).SetInt(s.stepDenom))
	numHalf, denomHalf := m.Num(), m.Denom()
	n, r := new(big.Int).QuoRem(new(big.Int).Mul(big.NewInt(2), numHalf), new(big.Int).Mul(big.NewInt(2), denomHalf), new(big.Int))
	switch t := r.Sign(); {
	case t > 0:
		if r.Cmp(denomHalf) >= 0 {
			n.Add(n, big.NewInt(1))
		}
	case t < 0:
		if r.Cmp(new(big.Int).Mul(big.NewInt(-1), denomHalf)) < 0 {
			n.Add(n, big.NewInt(-1))
		}
	}
	if n.Cmp(big.NewInt(0)) < 0 {
		result, _ := s.minRat.Float64()
		return result, ErrStepperMinExceeded
	}
	if n.Cmp(s.diffNum) > 0 {
		result, _ := s.maxRat.Float64()
		return result, ErrStepperMaxExceeded
	}
	y := new(big.Rat).Add(s.minRat, new(big.Rat).SetFrac(n, s.stepDenom))
	result, exact := y.Float64()
	if !exact {
		return result, ErrStepperInexactValue
	}
	return result, nil
}
