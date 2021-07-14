package xmath

import (
	"fmt"
	"math"
	"math/big"
)

// Real is a big.Float like implementation for real numbers.
// It rounds the big.Float to half away from zero by given precision and base after all of writing operations.
// A Real can be created with new(Real) or NewReal and etc.
// A Real which is created by new(Real) has precision 0 and base 10.
// Both of precision and base, can't change after the Real created.
type Real struct {
	prec int
	base int
	f    *big.Float
	k    *big.Float
}

// NewReal returns a new Real with given precision and base.
// Calling the NewReal, isn't necessary to create new Real.
// It panics unless base is in valid range.
func NewReal(prec, base int) *Real {
	checkInvalidBase(base)
	return &Real{
		prec: prec,
		base: base,
		f:    big.NewFloat(0),
		k:    big.NewFloat(math.Pow(float64(base), float64(prec))),
	}
}

// NewBinary returns a new Real with base 2.
func NewBinary(prec int) *Real {
	return NewReal(prec, 2)
}

// NewOctal returns a new Real with base 8.
func NewOctal(prec int) *Real {
	return NewReal(prec, 8)
}

// NewDecimal returns a new Real with base 10.
func NewDecimal(prec int) *Real {
	return NewReal(prec, 10)
}

// NewHexadecimal returns a new Real with base 16.
func NewHexadecimal(prec int) *Real {
	return NewReal(prec, 16)
}

// NewHexdecimal is synonym with NewHexadecimal.
func NewHexdecimal(prec int) *Real {
	return NewHexadecimal(prec)
}

func (z *Real) init() {
	if z.base != 0 {
		return
	}
	z.base = 10
	z.f = big.NewFloat(0)
	z.k = big.NewFloat(1)
}

// Prec returns precision of the Real.
func (x *Real) Prec() int {
	x.init()
	return x.prec
}

// Base returns base of the Real.
func (x *Real) Base() int {
	x.init()
	return x.base
}

// Float returns a copy of big.Float value.
func (x *Real) Float() *big.Float {
	x.init()
	return new(big.Float).Copy(x.f)
}

// FloatMinPrec is similar with MinPrec method of big.Float.
func (x *Real) FloatMinPrec() uint {
	x.init()
	return x.f.MinPrec()
}

// FloatMode is similar with Mode method of big.Float.
func (x *Real) FloatMode() big.RoundingMode {
	x.init()
	return x.f.Mode()
}

// FloatPrec is similar with Prec method of big.Float.
func (x *Real) FloatPrec() uint {
	x.init()
	return x.f.Prec()
}

// SetFloat is similar with Set except that x is big.Float.
func (z *Real) SetFloat(x *big.Float) *Real {
	z.init()
	defer z.round()
	z.f.Set(x)
	return z
}

// SetFloatMode is similar with SetMode method of big.Float.
func (z *Real) SetFloatMode(mode big.RoundingMode) *Real {
	z.init()
	defer z.round()
	z.f.SetMode(mode)
	return z
}

// SetFloatPrec is similar with SetPrec method of big.Float.
func (z *Real) SetFloatPrec(prec uint) *Real {
	z.init()
	defer z.round()
	z.f.SetPrec(prec)
	return z
}

func (z *Real) round() {
	if z.f.IsInf() {
		return
	}
	z.f.Mul(z.f, z.k)
	z.f.Add(z.f, big.NewFloat(0.5))
	z.f.SetInt(FloorBigFloat(z.f))
	z.f.Quo(z.f, z.k)
}

// Abs is similar with Abs method of big.Float.
func (z *Real) Abs(x *Real) *Real {
	z.init()
	defer z.round()
	z.f.Abs(x.f)
	return z
}

// Acc is similar with Acc method of big.Float.
func (x *Real) Acc() big.Accuracy {
	x.init()
	return x.f.Acc()
}

// Append is similar with Append method of big.Float.
func (x *Real) Append(buf []byte, fmt byte, prec int) []byte {
	x.init()
	return x.f.Append(buf, fmt, prec)
}

// Add is similar with Add method of big.Float.
func (z *Real) Add(x, y *Real) *Real {
	z.init()
	defer z.round()
	z.f.Add(x.f, y.f)
	return z
}

// Cmp is similar with Cmp method of big.Float.
func (x *Real) Cmp(y *Real) int {
	x.init()
	return x.f.Cmp(y.f)
}

// Copy is similar with Copy method of big.Float.
func (z *Real) Copy(x *Real) *Real {
	z.init()
	defer z.round()
	z.f.Copy(x.f)
	return z
}

// Float32 is similar with Float32 method of big.Float.
func (x *Real) Float32() (float32, big.Accuracy) {
	x.init()
	return x.f.Float32()
}

// Float64 is similar with Float64 method of big.Float.
func (x *Real) Float64() (float64, big.Accuracy) {
	x.init()
	return x.f.Float64()
}

// Format is similar with Format method of big.Float.
func (x *Real) Format(s fmt.State, format rune) {
	x.init()
	x.f.Format(s, format)
}

// GobDecode is similar with GobDecode method of big.Float.
func (z *Real) GobDecode(buf []byte) error {
	z.init()
	defer z.round()
	return z.f.GobDecode(buf)
}

// GobEncode is similar with GobEncode method of big.Float.
func (x *Real) GobEncode() ([]byte, error) {
	x.init()
	return x.f.GobEncode()
}

// Int is similar with Int method of big.Float.
func (x *Real) Int(z *big.Int) (*big.Int, big.Accuracy) {
	x.init()
	return x.f.Int(z)
}

// Int64 is similar with Int64 method of big.Float.
func (x *Real) Int64() (int64, big.Accuracy) {
	x.init()
	return x.f.Int64()
}

// IsInf is similar with IsInf method of big.Float.
func (x *Real) IsInf() bool {
	x.init()
	return x.f.IsInf()
}

// IsInt is similar with IsInt method of big.Float.
func (x *Real) IsInt() bool {
	x.init()
	return x.f.IsInt()
}

// MantExp is similar with MantExp method of big.Float.
func (x *Real) MantExp(mant *Real) (exp int) {
	x.init()
	return x.f.MantExp(mant.f)
}

// MarshalText is similar with MarshalText method of big.Float.
func (x *Real) MarshalText() (text []byte, err error) {
	x.init()
	return x.f.MarshalText()
}

// Mul is similar with Mul method of big.Float.
func (z *Real) Mul(x, y *Real) *Real {
	z.init()
	defer z.round()
	z.f.Mul(x.f, y.f)
	return z
}

// Neg is similar with Neg method of big.Float.
func (z *Real) Neg(x *Real) *Real {
	z.init()
	defer z.round()
	z.f.Neg(x.f)
	return z
}

// Parse is similar with Parse method of big.Float.
func (z *Real) Parse(s string, base int) (r *Real, b int, err error) {
	z.init()
	defer z.round()
	_, b, err = z.f.Parse(s, base)
	return z, b, err
}

// Quo is similar with Quo method of big.Float.
func (z *Real) Quo(x, y *Real) *Real {
	z.init()
	defer z.round()
	z.f.Quo(x.f, y.f)
	return z
}

// Rat is similar with Rat method of big.Float.
func (x *Real) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	x.init()
	return x.f.Rat(z)
}

// Scan is similar with Scan method of big.Float.
func (z *Real) Scan(s fmt.ScanState, ch rune) error {
	z.init()
	defer z.round()
	return z.f.Scan(s, ch)
}

// Set is similar with Set method of big.Float.
func (z *Real) Set(x *Real) *Real {
	z.init()
	defer z.round()
	z.f.Set(x.f)
	return z
}

// SetFloat64 is similar with SetFloat64 method of big.Float.
func (z *Real) SetFloat64(x float64) *Real {
	z.init()
	defer z.round()
	z.f.SetFloat64(x)
	return z
}

// SetInf is similar with SetInf method of big.Float.
func (z *Real) SetInf(signbit bool) *Real {
	z.init()
	defer z.round()
	z.f.SetInf(signbit)
	return z
}

// SetInt is similar with SetInt method of big.Float.
func (z *Real) SetInt(x *big.Int) *Real {
	z.init()
	defer z.round()
	z.f.SetInt(x)
	return z
}

// SetInt64 is similar with SetInt64 method of big.Float.
func (z *Real) SetInt64(x int64) *Real {
	z.init()
	defer z.round()
	z.f.SetInt64(x)
	return z
}

// SetMantExp is similar with SetMantExp method of big.Float.
func (z *Real) SetMantExp(mant *Real, exp int) *Real {
	z.init()
	defer z.round()
	z.f.SetMantExp(mant.f, exp)
	return z
}

// SetRat is similar with SetRat method of big.Float.
func (z *Real) SetRat(x *big.Rat) *Real {
	z.init()
	defer z.round()
	z.f.SetRat(x)
	return z
}

// SetString is similar with SetString method of big.Float.
func (z *Real) SetString(s string) (r *Real, ok bool) {
	z.init()
	defer z.round()
	_, ok = z.f.SetString(s)
	return z, ok
}

// SetUint64 is similar with SetUint64 method of big.Float.
func (z *Real) SetUint64(x uint64) *Real {
	z.init()
	defer z.round()
	z.f.SetUint64(x)
	return z
}

// Sign is similar with Sign method of big.Float.
func (x *Real) Sign() int {
	x.init()
	return x.f.Sign()
}

// Signbit is similar with Signbit method of big.Float.
func (x *Real) Signbit() bool {
	x.init()
	return x.f.Signbit()
}

// Sqrt is similar with Sqrt method of big.Float.
func (z *Real) Sqrt(x *Real) *Real {
	z.init()
	defer z.round()
	z.f.Sqrt(x.f)
	return z
}

// String is similar with String method of big.Float.
func (x *Real) String() string {
	x.init()
	return x.f.String()
}

// Sub is similar with Sub method of big.Float.
func (z *Real) Sub(x, y *Real) *Real {
	z.init()
	defer z.round()
	z.f.Sub(x.f, y.f)
	return z
}

// Text is similar with Text method of big.Float.
func (x *Real) Text(format byte, prec int) string {
	x.init()
	return x.f.Text(format, prec)
}

// Uint64 is similar with Uint64 method of big.Float.
func (x *Real) Uint64() (uint64, big.Accuracy) {
	x.init()
	return x.f.Uint64()
}

// UnmarshalText is similar with UnmarshalText method of big.Float.
func (z *Real) UnmarshalText(text []byte) error {
	z.init()
	defer z.round()
	return z.f.UnmarshalText(text)
}
