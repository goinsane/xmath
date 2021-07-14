package xmath

import (
	"errors"
	"math"
	"math/big"
)

var (
	ErrStepperStepOverflow    = errors.New("step overflow")
	ErrStepperMaxOverflow     = errors.New("max overflow")
	ErrStepperMinOverflow     = errors.New("min overflow")
	ErrStepperRangeOverflow   = errors.New("range overflow")
	ErrStepperUnorderedMaxMin = errors.New("unordered max min")
	ErrStepperMaxExceeded     = errors.New("max exceeded")
	ErrStepperMinExceeded     = errors.New("min exceeded")
)

// Stepper is a utility to
type Stepper struct {
	prec       int
	base       int
	stepReal   *Real
	maxReal    *Real
	minReal    *Real
	intrvlReal *Real
	step       float64
	max        float64
	min        float64
	count      int64
}

func NewStepper(prec, base int, step, max, min float64) (s *Stepper, err error) {
	checkInvalidBase(base)
	s = &Stepper{
		prec: prec,
		base: base,
	}
	s.stepReal = s.newReal().SetFloat64(step)
	if f, acc := s.stepReal.Float64(); f != step || acc != big.Exact {
		return nil, ErrStepperStepOverflow
	} else {
		s.step = f
	}
	s.maxReal = s.newReal().SetFloat64(max)
	if f, acc := s.maxReal.Float64(); f != max || acc != big.Exact || (!math.IsInf(max, 0) && math.Nextafter(max, math.Inf(+1))-max >= step) {
		return nil, ErrStepperMaxOverflow
	} else {
		s.max = f
	}
	s.minReal = s.newReal().SetFloat64(min)
	if f, acc := s.minReal.Float64(); f != min || acc != big.Exact || (!math.IsInf(min, 0) && min-math.Nextafter(min, math.Inf(-1)) >= step) {
		return nil, ErrStepperMinOverflow
	} else {
		s.min = f
	}
	s.intrvlReal = s.newReal().Sub(s.maxReal, s.minReal)
	r := s.newReal().Quo(s.intrvlReal, s.stepReal)
	if r.Cmp(s.newReal()) < 0 {
		return nil, ErrStepperUnorderedMaxMin
	}
	if !r.IsInf() {
		count, acc := r.Int64()
		if !r.IsInt() || acc != big.Exact {
			return nil, ErrStepperRangeOverflow
		}
		count++
		s.count = count
	}
	return s, nil
}

func (s *Stepper) newReal() *Real {
	return NewReal(s.prec, s.base)
}

func (s *Stepper) Prec() int {
	return s.prec
}

func (s *Stepper) Base() int {
	return s.base
}

func (s *Stepper) Count() int {
	if s.count > MaxIntValue {
		return MaxIntValue
	}
	return int(s.count)
}

func (s *Stepper) Count64() int64 {
	return s.count
}

func (s *Stepper) Step(index int) (float64, error) {
	return s.Step64(int64(index))
}

func (s *Stepper) Step64(index int64) (float64, error) {
	minReal := s.newReal()
	if !s.intrvlReal.IsInf() {
		if index >= s.count {
			return s.max, ErrStepperMaxExceeded
		}
		if index < 0 {
			return s.min, ErrStepperMinExceeded
		}
		minReal.Copy(s.minReal)
	}
	f, acc := s.newReal().Add(s.newReal().Mul(s.newReal().SetInt64(index), s.stepReal), minReal).Float64()
	if acc != big.Exact {
		panic("bug: result not exact")
	}
	return f, nil
}

func (s *Stepper) Normalize(f float64) (float64, error) {
	minReal := s.newReal()
	if !s.intrvlReal.IsInf() {
		if math.IsInf(f, +1) {
			return f, ErrStepperMaxExceeded
		}
		if math.IsInf(f, -1) {
			return f, ErrStepperMinExceeded
		}
		minReal.Copy(s.minReal)
	}
	if math.IsInf(f, 0) {
		return f, nil
	}
	if math.IsNaN(f) {
		return f, nil
	}
	return s.Step64(RoundBigFloat(s.newReal().Quo(s.newReal().Sub(s.newReal().SetFloat64(f), minReal), s.stepReal).Float()).Int64())
}
