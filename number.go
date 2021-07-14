package xmath

import (
	"fmt"
	"math"
	"math/big"
)

type Number struct {
	prec int
	base int
	f    *big.Float
	k    *big.Float
}

func NewNumber(prec, base int) *Number {
	checkInvalidBase(base)
	return &Number{
		prec: prec,
		base: base,
		f:    big.NewFloat(0),
		k:    big.NewFloat(math.Pow(float64(base), float64(prec))),
	}
}

func NewBinary(prec int) *Number {
	return NewNumber(prec, 2)
}

func NewOctal(prec int) *Number {
	return NewNumber(prec, 8)
}

func NewDecimal(prec int) *Number {
	return NewNumber(prec, 10)
}

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

func (x *Number) Prec() int {
	x.init()
	return x.prec
}

func (x *Number) Base() int {
	x.init()
	return x.base
}

func (x *Number) Float() *big.Float {
	x.init()
	return new(big.Float).Copy(x.f)
}

func (x *Number) FloatMinPrec() uint {
	x.init()
	return x.f.MinPrec()
}

func (x *Number) FloatMode() big.RoundingMode {
	x.init()
	return x.f.Mode()
}

func (x *Number) FloatPrec() uint {
	x.init()
	return x.f.Prec()
}

func (z *Number) SetFloat(x *big.Float) *Number {
	z.init()
	defer z.round()
	z.f.Set(x)
	return z
}

func (z *Number) SetFloatMode(mode big.RoundingMode) *Number {
	z.init()
	defer z.round()
	z.f.SetMode(mode)
	return z
}

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

func (z *Number) Abs(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Abs(x.f)
	return z
}

func (x *Number) Acc() big.Accuracy {
	x.init()
	return x.f.Acc()
}

func (x *Number) Append(buf []byte, fmt byte, prec int) []byte {
	x.init()
	return x.f.Append(buf, fmt, prec)
}

func (z *Number) Add(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Add(x.f, y.f)
	return z
}

func (x *Number) Cmp(y *Number) int {
	x.init()
	return x.f.Cmp(y.f)
}

func (z *Number) Copy(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Copy(x.f)
	return z
}

func (x *Number) Float32() (float32, big.Accuracy) {
	x.init()
	return x.f.Float32()
}

func (x *Number) Float64() (float64, big.Accuracy) {
	x.init()
	return x.f.Float64()
}

func (x *Number) Format(s fmt.State, format rune) {
	x.init()
	x.f.Format(s, format)
}

func (z *Number) GobDecode(buf []byte) error {
	z.init()
	defer z.round()
	return z.f.GobDecode(buf)
}

func (x *Number) GobEncode() ([]byte, error) {
	x.init()
	return x.f.GobEncode()
}

func (x *Number) Int(z *big.Int) (*big.Int, big.Accuracy) {
	x.init()
	return x.f.Int(z)
}

func (x *Number) Int64() (int64, big.Accuracy) {
	x.init()
	return x.f.Int64()
}

func (x *Number) IsInf() bool {
	x.init()
	return x.f.IsInf()
}

func (x *Number) IsInt() bool {
	x.init()
	return x.f.IsInt()
}

func (x *Number) MantExp(mant *Number) (exp int) {
	x.init()
	return x.f.MantExp(mant.f)
}

func (x *Number) MarshalText() (text []byte, err error) {
	x.init()
	return x.f.MarshalText()
}

func (z *Number) Mul(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Mul(x.f, y.f)
	return z
}

func (z *Number) Neg(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Neg(x.f)
	return z
}

func (z *Number) Parse(s string, base int) (n *Number, b int, err error) {
	z.init()
	defer z.round()
	_, b, err = z.f.Parse(s, base)
	return z, b, err
}

func (z *Number) Quo(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Quo(x.f, y.f)
	return z
}

func (x *Number) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	x.init()
	return x.f.Rat(z)
}

func (z *Number) Scan(s fmt.ScanState, ch rune) error {
	z.init()
	defer z.round()
	return z.f.Scan(s, ch)
}

func (z *Number) Set(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Set(x.f)
	return z
}

func (z *Number) SetFloat64(x float64) *Number {
	z.init()
	defer z.round()
	z.f.SetFloat64(x)
	return z
}

func (z *Number) SetInf(signbit bool) *Number {
	z.init()
	defer z.round()
	z.f.SetInf(signbit)
	return z
}

func (z *Number) SetInt(x *big.Int) *Number {
	z.init()
	defer z.round()
	z.f.SetInt(x)
	return z
}

func (z *Number) SetInt64(x int64) *Number {
	z.init()
	defer z.round()
	z.f.SetInt64(x)
	return z
}

func (z *Number) SetMantExp(mant *Number, exp int) *Number {
	z.init()
	defer z.round()
	z.f.SetMantExp(mant.f, exp)
	return z
}

func (z *Number) SetRat(x *big.Rat) *Number {
	z.init()
	defer z.round()
	z.f.SetRat(x)
	return z
}

func (z *Number) SetString(s string) (n *Number, ok bool) {
	z.init()
	defer z.round()
	_, ok = z.f.SetString(s)
	return z, ok
}

func (z *Number) SetUint64(x uint64) *Number {
	z.init()
	defer z.round()
	z.f.SetUint64(x)
	return z
}

func (x *Number) Sign() int {
	x.init()
	return x.f.Sign()
}

func (x *Number) Signbit() bool {
	x.init()
	return x.f.Signbit()
}

func (z *Number) Sqrt(x *Number) *Number {
	z.init()
	defer z.round()
	z.f.Sqrt(x.f)
	return z
}

func (x *Number) String() string {
	x.init()
	return x.f.String()
}

func (z *Number) Sub(x, y *Number) *Number {
	z.init()
	defer z.round()
	z.f.Sub(x.f, y.f)
	return z
}

func (x *Number) Text(format byte, prec int) string {
	x.init()
	return x.f.Text(format, prec)
}

func (x *Number) Uint64() (uint64, big.Accuracy) {
	x.init()
	return x.f.Uint64()
}

func (z *Number) UnmarshalText(text []byte) error {
	z.init()
	defer z.round()
	return z.f.UnmarshalText(text)
}
