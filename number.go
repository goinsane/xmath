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

func NewNumber(prec int, base int) *Number {
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

func (x *Number) Base() int {
	return x.base
}

func (x *Number) Prec() int {
	return x.prec
}

func (x *Number) Float() *big.Float {
	return new(big.Float).Copy(x.f)
}

func (x *Number) FloatMinPrec() uint {
	return x.f.MinPrec()
}

func (x *Number) FloatMode() big.RoundingMode {
	return x.f.Mode()
}

func (x *Number) FloatPrec() uint {
	return x.f.Prec()
}

func (z *Number) SetFloat(x *big.Float) *Number {
	defer z.round()
	z.f.Set(x)
	return z
}

func (z *Number) SetFloatMode(mode big.RoundingMode) *Number {
	defer z.round()
	z.f.SetMode(mode)
	return z
}

func (z *Number) SetFloatPrec(prec uint) *Number {
	defer z.round()
	z.f.SetPrec(prec)
	return z
}

func (z *Number) round() {
	z.f.Mul(z.f, z.k)
	z.f.Add(z.f, big.NewFloat(0.5))
	z.f.SetInt(FloorBigFloat(z.f))
	z.f.Quo(z.f, z.k)
}

func (z *Number) Abs(x *Number) *Number {
	defer z.round()
	z.f.Abs(x.f)
	return z
}

func (x *Number) Acc() big.Accuracy {
	return x.f.Acc()
}

func (x *Number) Append(buf []byte, fmt byte, prec int) []byte {
	return x.f.Append(buf, fmt, prec)
}

func (z *Number) Add(x, y *Number) *Number {
	defer z.round()
	z.f.Add(x.f, y.f)
	return z
}

func (x *Number) Cmp(y *Number) int {
	return x.f.Cmp(y.f)
}

func (z *Number) Copy(x *Number) *Number {
	defer z.round()
	z.f.Copy(x.f)
	return z
}

func (x *Number) Float32() (float32, big.Accuracy) {
	return x.f.Float32()
}

func (x *Number) Float64() (float64, big.Accuracy) {
	return x.f.Float64()
}

func (x *Number) Format(s fmt.State, format rune) {
	x.f.Format(s, format)
}

func (z *Number) GobDecode(buf []byte) error {
	defer z.round()
	return z.f.GobDecode(buf)
}

func (x *Number) GobEncode() ([]byte, error) {
	return x.f.GobEncode()
}

func (x *Number) Int(z *big.Int) (*big.Int, big.Accuracy) {
	return x.f.Int(z)
}

func (x *Number) Int64() (int64, big.Accuracy) {
	return x.f.Int64()
}

func (x *Number) IsInf() bool {
	return x.f.IsInf()
}

func (x *Number) IsInt() bool {
	return x.f.IsInt()
}

func (x *Number) MantExp(mant *Number) (exp int) {
	return x.f.MantExp(mant.f)
}

func (x *Number) MarshalText() (text []byte, err error) {
	return x.f.MarshalText()
}

func (z *Number) Mul(x, y *Number) *Number {
	defer z.round()
	z.f.Mul(x.f, y.f)
	return z
}

func (z *Number) Neg(x *Number) *Number {
	defer z.round()
	z.f.Neg(x.f)
	return z
}

func (z *Number) Parse(s string, base int) (n *Number, b int, err error) {
	defer z.round()
	_, b, err = z.f.Parse(s, base)
	return z, b, err
}

func (z *Number) Quo(x, y *Number) *Number {
	defer z.round()
	z.f.Quo(x.f, y.f)
	return z
}

func (x *Number) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	return x.f.Rat(z)
}

func (z *Number) Scan(s fmt.ScanState, ch rune) error {
	defer z.round()
	return z.f.Scan(s, ch)
}

func (z *Number) Set(x *Number) *Number {
	defer z.round()
	z.f.Set(x.f)
	return z
}

func (z *Number) SetFloat64(x float64) *Number {
	defer z.round()
	z.f.SetFloat64(x)
	return z
}

func (z *Number) SetInf(signbit bool) *Number {
	defer z.round()
	z.f.SetInf(signbit)
	return z
}

func (z *Number) SetInt(x *big.Int) *Number {
	defer z.round()
	z.f.SetInt(x)
	return z
}

func (z *Number) SetInt64(x int64) *Number {
	defer z.round()
	z.f.SetInt64(x)
	return z
}

func (z *Number) SetMantExp(mant *Number, exp int) *Number {
	defer z.round()
	z.f.SetMantExp(mant.f, exp)
	return z
}

func (z *Number) SetRat(x *big.Rat) *Number {
	defer z.round()
	z.f.SetRat(x)
	return z
}

func (z *Number) SetString(s string) (n *Number, ok bool) {
	defer z.round()
	_, ok = z.f.SetString(s)
	return z, ok
}

func (z *Number) SetUint64(x uint64) *Number {
	defer z.round()
	z.f.SetUint64(x)
	return z
}

func (x *Number) Sign() int {
	return x.f.Sign()
}

func (x *Number) Signbit() bool {
	return x.f.Signbit()
}

func (z *Number) Sqrt(x *Number) *Number {
	defer z.round()
	z.f.Sqrt(x.f)
	return z
}

func (x *Number) String() string {
	return x.f.String()
}

func (z *Number) Sub(x, y *Number) *Number {
	defer z.round()
	z.f.Sub(x.f, y.f)
	return z
}

func (x *Number) Text(format byte, prec int) string {
	return x.f.Text(format, prec)
}

func (x *Number) Uint64() (uint64, big.Accuracy) {
	return x.Uint64()
}

func (z *Number) UnmarshalText(text []byte) error {
	defer z.round()
	return z.f.UnmarshalText(text)
}
