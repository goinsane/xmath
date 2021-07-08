package xmath

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	ErrStepperStepOverflow    = errors.New("step overflow")
	ErrStepperMaxOverflow     = errors.New("max overflow")
	ErrStepperMinOverflow     = errors.New("min overflow")
	ErrStepperUnorderedMaxMin = errors.New("unordered max min")
	ErrStepperMaxExceeded     = errors.New("max exceeded")
	ErrStepperMinExceeded     = errors.New("min exceeded")
)

type Stepper struct {
	step   float64
	max    float64
	min    float64
	exp int
	stepFrac *big.Int
	//stepExp int
	maxFrac *big.Int
	//maxExp int
	minFrac *big.Int
	//minExp int
}

func NewStepper(step, max, min float64) (s *Stepper, err error) {
	/*factor := 1.0 / step
	if step <= 0 || IsZero(step) || step != 1.0/factor {
		return nil, ErrStepperStepOverflow
	}
	if math.Nextafter(max, math.Inf(+1))-max > step {
		return nil, ErrStepperMaxOverflow
	}
	if min-math.Nextafter(min, math.Inf(-1)) > step {
		return nil, ErrStepperMinOverflow
	}
	if max < min {
		return nil, ErrStepperUnorderedMaxMin
	}*/
	s = &Stepper{
		step:   step,
		max:    max,
		min:    min,
	}
	stepFrac, stepExp := math.Frexp(s.step)
	s.stepFrac = big.NewInt(int64(stepFrac*math.Exp2(53)))
	maxFrac, maxExp := math.Frexp(s.max)
	s.maxFrac = big.NewInt(int64(maxFrac*math.Exp2(53)))
	minFrac, minExp := math.Frexp(s.min)
	s.minFrac = big.NewInt(int64(minFrac*math.Exp2(53)))
	s.exp = stepExp
	if maxExp < s.exp {
		s.exp = maxExp
	}
	if minExp < s.exp {
		s.exp = minExp
	}
	s.exp--
	s.stepFrac.Mul(s.stepFrac, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(stepExp-s.exp)), nil))
	s.maxFrac.Mul(s.maxFrac, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(maxExp-s.exp)), nil))
	s.minFrac.Mul(s.minFrac, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(minExp-s.exp)), nil))
	return s, nil
}

func (s *Stepper) Count() int64 {
	return 0
}

func (s *Stepper) Step(index int64) (float64, error) {
	return s.Normalize(float64(index)*s.step)
}

func (s *Stepper) Normalize(x float64) (float64, error) {
	frac, exp := math.Frexp(x)
	bigFrac := big.NewInt(int64(frac*math.Exp2(53)))
	// exp may be below from s.exp !!!
	bigFrac.Mul(bigFrac, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(exp-s.exp)), nil))
	//halfFrac := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(s.exp-1)), nil)
	n, r := new(big.Int).DivMod(bigFrac, s.stepFrac, new(big.Int))
	halfFrac := new(big.Int).Quo(s.stepFrac, big.NewInt(2))
	bigFrac.Sub(bigFrac, r)
	if r.Cmp(halfFrac) >= 0 {
		bigFrac.Add(bigFrac, s.stepFrac)
		n.Add(n, big.NewInt(1))
	}
	//denom := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(53-s.exp)), nil)
	//rat := new(big.Rat).SetFrac(bigFrac, denom)
	//f, _ := rat.Float64()
	//fmt.Println(bigFrac, denom)
	//fmt.Println(math.Ldexp(float64(bigFrac.Int64()), -(53-s.exp)))
	fmt.Printf("%+b %v\n", bigFrac.Int64(), math.Ldexp(float64(bigFrac.Int64()), -(53-s.exp)))
	return 0, nil
}
