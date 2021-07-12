package xmath

import (
	"errors"
	"math"
	"math/big"
)

var (
	ErrStepOverflow    = errors.New("step overflow")
	ErrMaxOverflow     = errors.New("max overflow")
	ErrMinOverflow     = errors.New("min overflow")
	ErrRangeOverflow   = errors.New("range overflow")
	ErrUnorderedMaxMin = errors.New("unordered max min")
	ErrMaxExceeded     = errors.New("max exceeded")
	ErrMinExceeded     = errors.New("min exceeded")
)

type Stepper struct {
	prec    int
	base    int
	stepNum *Number
	maxNum  *Number
	minNum  *Number
	count   int64
}

func NewStepper(prec, base int, step, max, min float64) (s *Stepper, err error) {
	checkInvalidBase(base)
	s = &Stepper{
		prec: prec,
		base: base,
	}
	s.stepNum = s.newNumber().SetFloat64(step)
	if f, acc := s.stepNum.Float64(); f != step || acc != big.Exact {
		return nil, ErrStepOverflow
	}
	s.maxNum = s.newNumber().SetFloat64(max)
	if f, acc := s.maxNum.Float64(); f != max || acc != big.Exact || math.Nextafter(max, math.Inf(+1))-max >= step {
		return nil, ErrMaxOverflow
	}
	s.minNum = s.newNumber().SetFloat64(min)
	if f, acc := s.minNum.Float64(); f != min || acc != big.Exact || min-math.Nextafter(min, math.Inf(-1)) >= step {
		return nil, ErrMinOverflow
	}
	n := RoundBigFloat(s.newNumber().Quo(s.newNumber().Sub(s.maxNum, s.minNum), s.stepNum).Float())
	if !n.IsInt64() {
		return nil, ErrRangeOverflow
	}
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, ErrUnorderedMaxMin
	}
	s.count = n.Int64()
	return s, nil
}

func (s *Stepper) newNumber() *Number {
	return NewNumber(s.prec, s.base)
}

func (s *Stepper) Count() int {
	return int(s.count)
}

func (s *Stepper) Count64() int64 {
	return s.count
}

func (s *Stepper) Step(index int) (float64, error) {
	return s.Step64(int64(index))
}

func (s *Stepper) Step64(index int64) (float64, error) {
	if index >= s.count {
		return math.NaN(), ErrMaxExceeded
	}
	if index < 0 {
		return math.NaN(), ErrMinExceeded
	}
	f, acc := s.newNumber().Add(s.minNum, s.newNumber().Mul(s.newNumber().SetInt64(index), s.stepNum)).Float64()
	if acc != big.Exact {
		panic("bug: result not exact")
	}
	return f, nil
}

func (s *Stepper) Normalize(f float64) (float64, error) {
	if math.IsInf(f, +1) {
		return math.Inf(+1), ErrMaxExceeded
	}
	if math.IsInf(f, -1) {
		return math.Inf(-1), ErrMinExceeded
	}
	if math.IsNaN(f) {
		return math.NaN(), nil
	}
	n := RoundBigFloat(s.newNumber().Quo(s.newNumber().Sub(s.newNumber().SetFloat64(f), s.minNum), s.stepNum).Float())
	return s.Step64(n.Int64())
}
