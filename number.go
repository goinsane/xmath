package xmath

import (
	"fmt"
	"math"
	"math/big"
)

// Number is a big.Float like implementation for integers and real numbers.
// It rounds the big.Float to half away from zero after all of writing operations.
// A Number can be created with new(Number) or NewNumber and etc.
// A Number which is created by new(Number) has precision 0 and base 10.
// Both of Precision and base, can't change after the Number created.
type Number struct {
	prec int
	base int
	f    *big.Float
	k    *big.Float
}

// NewNumber returns a new Number with given precision and base.
// Calling the NewNumber, isn't necessary to create new Number.
func NewNumber(prec, base int) *Number {
	checkInvalidBase(base)
	return &Number{
		prec: prec,
		base: base,
		f:    big.NewFloat(0),
		k:    big.NewFloat(math.Pow(float64(base), float64(prec))),
	}
}

// NewBinary returns a new Number with base 2.
func NewBinary(prec int) *Number {
	return NewNumber(prec, 2)
}

// NewOctal returns a new Number with base 8.
func NewOctal(prec int) *Number {
	return NewNumber(prec, 8)
}

// NewDecimal returns a new Number with base 10.
func NewDecimal(prec int) *Number {
	return NewNumber(prec, 10)
}

// NewHexdecimal returns a new Number with base 16.
func NewHexdecimal(prec int) *Number {
	return NewNumber(prec, 16)
}

func (z *Number) init() {
	if z.base != 0 {
		return
	}
	z.base = 10
	z.f = big.NewFloat(0)
	z.k = big.NewFloat(1)
}

// Prec returns precision of the Number.
func (x *Number) Prec() int {
	x.init()
	return x.prec
}

// Base returns base of the Number.
func (x *Number) Base() int {
	x.init()
	return x.base
}

// Float returns a copy of big.Float value.
func (x *Number) Float() *big.Float {
	x.init()
	return new(big.Float).Copy(x.f)
}

// FloatMinPrec is similar with MinPrec method of big.Float.
func (x *Number) FloatMinPrec() uint {
	x.init()
	return x.f.MinPrec()
}

// FloatMode is similar with Mode method of big.Float.
func (x *Number) FloatMode() big.RoundingMode {
	x.init()
	return x.f.Mode()
}

// FloatPrec is similar with Prec method of big.Float.
func (x *Number) FloatPrec() uint {
	x.init()
	return x.f.Prec()
}

// SetFloat is similar with Set except that x is big.Float.
func (z *Number) SetFloat(x *big.Float) *Number {
	z.init()
	defer z.round()
	z.f.Set(x)
	return z
}

// SetFloatMode is similar with SetMode method of big.Float.
func (z *Number) SetFloatMode(mode big.RoundingMode) *Number {
	z.init()
	defer z.round()
	z.f.SetMode(mode)
	return z
}

// SetFloatPrec is similar with SetPrec method of big.Float.
func (z *Number) SetFloatPrec(prec uint) *Number {
	z.init()
	defer z.round()
	z.f.SetPrec(prec)
	return z
}

func (z *Number) round() {
	if z.f.IsInf() {
		return
	}
	z.f.Mul(z.f, z.k)
	z.f.Add(z.f, big.NewFloat(0.5))
	z.f.SetInt(FloorBigFloat(z.f))
	z.f.Quo(z.f, z.k)
}

// Abs is similar with Abs method of big.Float.
func (z *Number) Abs(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Abs(x.f)
	return z
}

// Acc is similar with Acc method of big.Float.
func (x *Number) Acc() big.Accuracy {
	x.init()
	return x.f.Acc()
}

// Append is similar with Append method of big.Float.
func (x *Number) Append(buf []byte, fmt byte, prec int) []byte {
	x.init()
	return x.f.Append(buf, fmt, prec)
}

// Add is similar with Add method of big.Float.
func (z *Number) Add(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Add(x.f, y.f)
	return z
}

// Cmp is similar with Cmp method of big.Float.
func (x *Number) Cmp(y *Number) int {
	x.init()
	return x.f.Cmp(y.f)
}

// Copy is similar with Copy method of big.Float.
func (z *Number) Copy(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Copy(x.f)
	return z
}

// Float32 is similar with Float32 method of big.Float.
func (x *Number) Float32() (float32, big.Accuracy) {
	x.init()
	return x.f.Float32()
}

// Float64 is similar with Float64 method of big.Float.
func (x *Number) Float64() (float64, big.Accuracy) {
	x.init()
	return x.f.Float64()
}

// Format is similar with Format method of big.Float.
func (x *Number) Format(s fmt.State, format rune) {
	x.init()
	x.f.Format(s, format)
}

// GobDecode is similar with GobDecode method of big.Float.
func (z *Number) GobDecode(buf []byte) error {
	z.init()
	defer z.round()
	return z.f.GobDecode(buf)
}

// GobEncode is similar with GobEncode method of big.Float.
func (x *Number) GobEncode() ([]byte, error) {
	x.init()
	return x.f.GobEncode()
}

// Int is similar with Int method of big.Float.
func (x *Number) Int(z *big.Int) (*big.Int, big.Accuracy) {
	x.init()
	return x.f.Int(z)
}

// Int64 is similar with Int64 method of big.Float.
func (x *Number) Int64() (int64, big.Accuracy) {
	x.init()
	return x.f.Int64()
}

// IsInf is similar with IsInf method of big.Float.
func (x *Number) IsInf() bool {
	x.init()
	return x.f.IsInf()
}

// IsInt is similar with IsInt method of big.Float.
func (x *Number) IsInt() bool {
	x.init()
	return x.f.IsInt()
}

// MantExp is similar with MantExp method of big.Float.
func (x *Number) MantExp(mant *Number) (exp int) {
	x.init()
	return x.f.MantExp(mant.f)
}

// MarshalText is similar with MarshalText method of big.Float.
func (x *Number) MarshalText() (text []byte, err error) {
	x.init()
	return x.f.MarshalText()
}

// Mul is similar with Mul method of big.Float.
func (z *Number) Mul(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Mul(x.f, y.f)
	return z
}

// Neg is similar with Neg method of big.Float.
func (z *Number) Neg(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Neg(x.f)
	return z
}

// Parse is similar with Parse method of big.Float.
func (z *Number) Parse(s string, base int) (n *Number, b int, err error) {
	z.init()
	defer z.round()
	_, b, err = z.f.Parse(s, base)
	return z, b, err
}

// Quo is similar with Quo method of big.Float.
func (z *Number) Quo(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Quo(x.f, y.f)
	return z
}

// Rat is similar with Rat method of big.Float.
func (x *Number) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	x.init()
	return x.f.Rat(z)
}

// Scan is similar with Scan method of big.Float.
func (z *Number) Scan(s fmt.ScanState, ch rune) error {
	z.init()
	defer z.round()
	return z.f.Scan(s, ch)
}

// Set is similar with Set method of big.Float.
func (z *Number) Set(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Set(x.f)
	return z
}

// SetFloat64 is similar with SetFloat64 method of big.Float.
func (z *Number) SetFloat64(x float64) *Number {
	z.init()
	defer z.round()
	z.f.SetFloat64(x)
	return z
}

// SetInf is similar with SetInf method of big.Float.
func (z *Number) SetInf(signbit bool) *Number {
	z.init()
	defer z.round()
	z.f.SetInf(signbit)
	return z
}

// SetInt is similar with SetInt method of big.Float.
func (z *Number) SetInt(x *big.Int) *Number {
	z.init()
	defer z.round()
	z.f.SetInt(x)
	return z
}

// SetInt64 is similar with SetInt64 method of big.Float.
func (z *Number) SetInt64(x int64) *Number {
	z.init()
	defer z.round()
	z.f.SetInt64(x)
	return z
}

// SetMantExp is similar with SetMantExp method of big.Float.
func (z *Number) SetMantExp(mant *Number, exp int) *Number {
	z.init()
	defer z.round()
	z.f.SetMantExp(mant.f, exp)
	return z
}

// SetRat is similar with SetRat method of big.Float.
func (z *Number) SetRat(x *big.Rat) *Number {
	z.init()
	defer z.round()
	z.f.SetRat(x)
	return z
}

// SetString is similar with SetString method of big.Float.
func (z *Number) SetString(s string) (n *Number, ok bool) {
	z.init()
	defer z.round()
	_, ok = z.f.SetString(s)
	return z, ok
}

// SetUint64 is similar with SetUint64 method of big.Float.
func (z *Number) SetUint64(x uint64) *Number {
	z.init()
	defer z.round()
	z.f.SetUint64(x)
	return z
}

// Sign is similar with Sign method of big.Float.
func (x *Number) Sign() int {
	x.init()
	return x.f.Sign()
}

// Signbit is similar with Signbit method of big.Float.
func (x *Number) Signbit() bool {
	x.init()
	return x.f.Signbit()
}

// Sqrt is similar with Sqrt method of big.Float.
func (z *Number) Sqrt(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Sqrt(x.f)
	return z
}

// String is similar with String method of big.Float.
func (x *Number) String() string {
	x.init()
	return x.f.String()
}

// Sub is similar with Sub method of big.Float.
func (z *Number) Sub(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Sub(x.f, y.f)
	return z
}

// Text is similar with Text method of big.Float.
func (x *Number) Text(format byte, prec int) string {
	x.init()
	return x.f.Text(format, prec)
}

// Uint64 is similar with Uint64 method of big.Float.
func (x *Number) Uint64() (uint64, big.Accuracy) {
	x.init()
	return x.f.Uint64()
}

// UnmarshalText is similar with UnmarshalText method of big.Float.
func (z *Number) UnmarshalText(text []byte) error {
	z.init()
	defer z.round()
	return z.f.UnmarshalText(text)
}
