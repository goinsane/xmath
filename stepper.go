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

type Stepper struct {
	prec     int
	base     int
	stepReal *Real
	maxReal  *Real
	minReal  *Real
	step     float64
	max      float64
	min      float64
	count    int64
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
	if f, acc := s.maxReal.Float64(); f != max || acc != big.Exact || math.Nextafter(max, math.Inf(+1))-max >= step {
		return nil, ErrStepperMaxOverflow
	} else {
		s.max = f
	}
	s.minReal = s.newReal().SetFloat64(min)
	if f, acc := s.minReal.Float64(); f != min || acc != big.Exact || min-math.Nextafter(min, math.Inf(-1)) >= step {
		return nil, ErrStepperMinOverflow
	} else {
		s.min = f
	}
	r := s.newReal().Quo(s.newReal().Sub(s.maxReal, s.minReal), s.stepReal)
	count, acc := r.Int64()
	if !r.IsInt() || acc != big.Exact {
		return nil, ErrStepperRangeOverflow
	}
	if count < 0 {
		return nil, ErrStepperUnorderedMaxMin
	}
	count++
	s.count = count
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
	if index >= s.count {
		return s.max, ErrStepperMaxExceeded
	}
	if index < 0 {
		return s.min, ErrStepperMinExceeded
	}
	f, acc := s.newReal().Add(s.minReal, s.newReal().Mul(s.newReal().SetInt64(index), s.stepReal)).Float64()
	if acc != big.Exact {
		panic("bug: result not exact")
	}
	return f, nil
}

func (s *Stepper) Normalize(f float64) (float64, error) {
	if math.IsInf(f, +1) {
		return math.Inf(+1), ErrStepperMaxExceeded
	}
	if math.IsInf(f, -1) {
		return math.Inf(-1), ErrStepperMinExceeded
	}
	if math.IsNaN(f) {
		return math.NaN(), nil
	}
	return s.Step64(RoundBigFloat(s.newReal().Quo(s.newReal().Sub(s.newReal().SetFloat64(f), s.minReal), s.stepReal).Float()).Int64())
}
