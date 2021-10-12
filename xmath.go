// Package xmath provides some extended capabilities according to GoLang's math.
package xmath

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

const (
	MinBase = 2
	MaxBase = 36

	MaxInt8Value   = 1<<7 - 1
	MinInt8Value   = -1 << 7
	MaxInt16Value  = 1<<15 - 1
	MinInt16Value  = -1 << 15
	MaxInt32Value  = 1<<31 - 1
	MinInt32Value  = -1 << 31
	MaxInt64Value  = 1<<63 - 1
	MinInt64Value  = -1 << 63
	MaxUint8Value  = 1<<8 - 1
	MaxUint16Value = 1<<16 - 1
	MaxUint32Value = 1<<32 - 1
	MaxUint64Value = 1<<64 - 1

	MaxIntValue  = 1<<(UintSize-1) - 1
	MinIntValue  = -1 << (UintSize - 1)
	MaxUintValue = 1<<UintSize - 1

	UintSize = 32 << (^uint(0) >> 32 & 1)
)

// FloorP returns the greatest value less than or equal to x with specified decimal precision.
func FloorP(x float64, prec int) float64 {
	k := math.Pow10(prec)
	return math.Floor(x*k) / k
}

// FloorPB returns the greatest value less than or equal to x with specified precision of base.
// It panics unless base is in valid range.
func FloorPB(x float64, prec int, base int) float64 {
	checkInvalidBase(base)
	k := math.Pow(float64(base), float64(prec))
	return math.Floor(x*k) / k
}

// CeilP returns the least value greater than or equal to x with specified decimal precision.
func CeilP(x float64, prec int) float64 {
	k := math.Pow10(prec)
	return math.Ceil(x*k) / k
}

// CeilPB returns the least value greater than or equal to x with specified precision of base.
// It panics unless base is in valid range.
func CeilPB(x float64, prec int, base int) float64 {
	checkInvalidBase(base)
	k := math.Pow(float64(base), float64(prec))
	return math.Ceil(x*k) / k
}

// Round returns the nearest integer value, rounding half away from zero.
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// RoundP returns the nearest integer value, rounding half away from zero with specified decimal precision.
func RoundP(x float64, prec int) float64 {
	k := math.Pow10(prec)
	return math.Floor(x*k+0.5) / k
}

// RoundPB returns the nearest integer value, rounding half away from zero with specified precision of base.
// It panics unless base is in valid range.
func RoundPB(x float64, prec int, base int) float64 {
	checkInvalidBase(base)
	k := math.Pow(float64(base), float64(prec))
	return math.Floor(x*k+0.5) / k
}

// Max returns the larger of x...
//
// Special cases are:
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
//	Max(x) = x
//	Max() = +Inf
func Max(x ...float64) float64 {
	if len(x) <= 0 {
		return math.Inf(+1)
	}
	result := math.Inf(-1)
	for _, a := range x {
		result = math.Max(a, result)
	}
	return result
}

// Min returns the smaller of x...
//
// Special cases are:
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
//	Min(x) = x
//	Min() = -Inf
func Min(x ...float64) float64 {
	if len(x) <= 0 {
		return math.Inf(-1)
	}
	result := math.Inf(+1)
	for _, a := range x {
		result = math.Min(a, result)
	}
	return result
}

// MaxMin returns the max, min values in this order, similar with Max and Min functions.
//
// Special cases are:
//	MaxMin(x) = x, x
//	MaxMin() = +Inf, -Inf
func MaxMin(x ...float64) (max float64, min float64) {
	min, max = MinMax(x...)
	return
}

// MinMax returns the min, max values in this order, similar with Min and Max functions.
//
// Special cases are:
//	MinMax(x) = x, x
//	MinMax() = -Inf, +Inf
func MinMax(x ...float64) (min float64, max float64) {
	if len(x) <= 0 {
		return math.Inf(-1), math.Inf(+1)
	}
	min = math.Inf(+1)
	max = math.Inf(-1)
	for _, a := range x {
		min = math.Min(a, min)
		max = math.Max(a, max)
	}
	return
}

// MaxInt returns the larger integer of x...
//
// Special cases are:
//	MaxInt(x) = x
//	MaxInt() = math.MaxInt64
func MaxInt(x ...int64) int64 {
	if len(x) <= 0 {
		return int64(math.MaxInt64)
	}
	result := int64(math.MinInt64)
	for _, a := range x {
		if a > result {
			result = a
		}
	}
	return result
}

// MinInt returns the smaller integer of x...
//
// Special cases are:
//	MinInt(x) = x
//	MinInt() = math.MinInt64
func MinInt(x ...int64) int64 {
	if len(x) <= 0 {
		return int64(math.MinInt64)
	}
	result := int64(math.MaxInt64)
	for _, a := range x {
		if a < result {
			result = a
		}
	}
	return result
}

// MaxMinInt returns the max, min integers in this order, similar with MaxInt and MinInt functions.
//
// Special cases are:
//	MaxMinInt(x) = x, x
//	MaxMinInt() = math.MaxInt64, math.MinInt64
func MaxMinInt(x ...int64) (max int64, min int64) {
	min, max = MinMaxInt(x...)
	return
}

// MinMaxInt returns the min, max integers in this order, similar with MinInt and MaxInt functions.
//
// Special cases are:
//	MinMaxInt(x) = x, x
//	MinMaxInt() = math.MinInt64, math.MaxInt64
func MinMaxInt(x ...int64) (min int64, max int64) {
	if len(x) <= 0 {
		return int64(math.MinInt64), int64(math.MaxInt64)
	}
	min = int64(math.MaxInt64)
	max = int64(math.MinInt64)
	for _, a := range x {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return
}

// MaxUint returns the larger unsigned integer of x...
//
// Special cases are:
//	MaxUint(x) = x
//	MaxUint() = math.MaxUint64
func MaxUint(x ...uint64) uint64 {
	if len(x) <= 0 {
		return uint64(math.MaxUint64)
	}
	result := uint64(0)
	for _, a := range x {
		if a > result {
			result = a
		}
	}
	return result
}

// MinUint returns the smaller unsigned integer of x...
//
// Special cases are:
//	MinUint(x) = x
//	MinUint() = 0
func MinUint(x ...uint64) uint64 {
	if len(x) <= 0 {
		return uint64(0)
	}
	result := uint64(math.MaxUint64)
	for _, a := range x {
		if a < result {
			result = a
		}
	}
	return result
}

// MaxMinUint returns the max, min unsigned integers in this order, similar with MaxUint and MinUint functions.
//
// Special cases are:
//	MaxMinUint(x) = x, x
//	MaxMinUint() = math.MaxUint64, 0
func MaxMinUint(x ...uint64) (max uint64, min uint64) {
	min, max = MinMaxUint(x...)
	return
}

// MinMaxUint returns the min, max unsigned integers in this order, similar with MinUint and MaxUint functions.
//
// Special cases are:
//	MinMaxUint(x) = x, x
//	MinMaxUint() = 0, math.MaxUint64
func MinMaxUint(x ...uint64) (min uint64, max uint64) {
	if len(x) <= 0 {
		return uint64(0), uint64(math.MaxUint64)
	}
	min = uint64(math.MaxUint64)
	max = uint64(0)
	for _, a := range x {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return
}

// Between checks x is between a and b
func Between(x float64, a, b float64) bool {
	min, max := MinMax(a, b)
	return min < x && x < max
}

// BetweenIn checks x is in a and b
func BetweenIn(x float64, a, b float64) bool {
	min, max := MinMax(a, b)
	return min <= x && x <= max
}

// SafeDiv divides x to y. For 'division by zero', it returns 0 if allowNaN is false.
// The GoLang's default behaviour is same with SafeDiv(x, y, true).
// Special cases are:
//	SafeDiv(0, ±n, true) = ±0
//	SafeDiv(0, ±n, false) = ±0
//	SafeDiv(±n, 0, true) = ±Inf
//	SafeDiv(±n, 0, false) = ±Inf
//	SafeDiv(0, 0, true) = NaN
//	SafeDiv(0, 0, false) = 0
func SafeDiv(x, y float64, allowNaN bool) float64 {
	if y == 0 {
		if x < 0 {
			return math.Inf(-1)
		}
		if x > 0 {
			return math.Inf(+1)
		}
		if allowNaN {
			return math.NaN()
		}
		return 0
	}
	return x / y
}

// CryptoRandInt returns a random integer in [0, max).
// It returns -1 when error occurs.
func CryptoRandInt(max int64) int64 {
	if max <= 0 {
		return -1
	}
	num, _ := crand.Int(crand.Reader, big.NewInt(max))
	if num == nil {
		return -1
	}
	return num.Int64()
}

// CryptoRandFloat returns a random decimal number in [0, 1).
// It returns -1 when error occurs.
func CryptoRandFloat() float64 {
	r := CryptoRandInt(math.MaxInt64)
	if r < 0 {
		return -1
	}
	return float64(r) / math.MaxInt64
}

// CryptoRand is synonym with CryptoRandFloat.
func CryptoRand() float64 {
	return CryptoRandFloat()
}

// CryptoRandCode generates random code in [10^(n-1), 10^n).
// It returns -1 when error occurs.
func CryptoRandCode(n int) int64 {
	if n < 1 || n > 18 {
		return -1
	}
	start := int64(1)
	for i := 0; i < n-1; i++ {
		start *= 10
	}
	r := CryptoRandInt(start*10 - start)
	if r < 0 {
		return -1
	}
	return start + r
}

// AlmostEqualP64 checks almost equality of all given 64-bit floating points values.
// Argument p is measure of precision. If p is 0, it checks exact equality.
// It returns true if all values are almost equal.
//
// Special cases are:
//	AlmostEqualP64(p) = false
//	AlmostEqualP64(p, x) = true
//	AlmostEqualP64(p, NaN) = false
//	AlmostEqualP64(p, NaN, x) = false
//	AlmostEqualP64(p, x, NaN) = false
//	AlmostEqualP64(p, +Inf, +Inf) = true
//	AlmostEqualP64(p, -Inf, -Inf) = true
func AlmostEqualP64(p uint64, x ...float64) bool {
	if len(x) <= 0 {
		return false
	}
	var a float64
	for i, b := range x {
		if math.IsNaN(b) {
			return false
		}
		if i > 0 {
			c, d := math.Float64bits(a), math.Float64bits(b)
			if c < d {
				c, d = d, c
			}
			if (c>>52 != d>>52) || c-d > p {
				return false
			}
		}
		a = b
	}
	return true
}

// AlmostEqualP32 checks almost equality of all given 32-bit floating points values.
// Argument p is measure of precision. If p is 0, it checks exact equality.
// It returns true if all values are almost equal.
//
// Special cases are:
//	AlmostEqualP32(p) = false
//	AlmostEqualP32(p, x) = true
//	AlmostEqualP32(p, NaN) = false
//	AlmostEqualP32(p, NaN, x) = false
//	AlmostEqualP32(p, x, NaN) = false
//	AlmostEqualP32(p, +Inf, +Inf) = true
//	AlmostEqualP32(p, -Inf, -Inf) = true
func AlmostEqualP32(p uint32, x ...float32) bool {
	if len(x) <= 0 {
		return false
	}
	var a float32
	for i, b := range x {
		if math.IsNaN(float64(b)) {
			return false
		}
		if i > 0 {
			c, d := math.Float32bits(a), math.Float32bits(b)
			if c < d {
				c, d = d, c
			}
			if (c>>23 != d>>23) || c-d > p {
				return false
			}
		}
		a = b
	}
	return true
}

// AlmostEqualP is synonym with AlmostEqualP64.
func AlmostEqualP(p uint64, x ...float64) bool {
	return AlmostEqualP64(p, x...)
}

// AlmostEqual64 is synonym with AlmostEqualP64(1, x...).
func AlmostEqual64(x ...float64) bool {
	return AlmostEqualP64(1, x...)
}

// AlmostEqual32 is synonym with AlmostEqualP32(1, x...).
func AlmostEqual32(x ...float32) bool {
	return AlmostEqualP32(1, x...)
}

// AlmostEqual is synonym with AlmostEqualP(1, x...).
func AlmostEqual(x ...float64) bool {
	return AlmostEqualP(1, x...)
}

// AlmostEqualD64 checks almost equality of all given 64-bit floating points values.
// Argument d is the least difference value of inequality. If d is 0, it checks exact equality.
// It returns true if all values are almost equal.
//
// Special cases are:
//	AlmostEqualD64(d) = false
//	AlmostEqualD64(d, x) = true
//	AlmostEqualD64(d, NaN) = false
//	AlmostEqualD64(d, NaN, x) = false
//	AlmostEqualD64(d, x, NaN) = false
//	AlmostEqualD64(d, +Inf, +Inf) = true
//	AlmostEqualD64(d, -Inf, -Inf) = true
func AlmostEqualD64(d float64, x ...float64) bool {
	if len(x) <= 0 {
		return false
	}
	var a float64
	for i, b := range x {
		if math.IsNaN(b) {
			return false
		}
		if i > 0 && math.Abs(a-b) >= d {
			return false
		}
		a = b
	}
	return true
}

// AlmostEqualD32 checks almost equality of all given 32-bit floating points values.
// Argument d is the least difference value of inequality. If d is 0, it checks exact equality.
// It returns true if all values are almost equal.
//
// Special cases are:
//	AlmostEqualD32(d) = false
//	AlmostEqualD32(d, x) = true
//	AlmostEqualD32(d, NaN) = false
//	AlmostEqualD32(d, NaN, x) = false
//	AlmostEqualD32(d, x, NaN) = false
//	AlmostEqualD32(d, +Inf, +Inf) = true
//	AlmostEqualD32(d, -Inf, -Inf) = true
func AlmostEqualD32(d float32, x ...float32) bool {
	if len(x) <= 0 {
		return false
	}
	var a float32
	for i, b := range x {
		if math.IsNaN(float64(b)) {
			return false
		}
		if i > 0 && math.Abs(float64(a-b)) >= float64(d) {
			return false
		}
		a = b
	}
	return true
}

// AlmostEqualD is synonym with AlmostEqualD64.
func AlmostEqualD(d float64, x ...float64) bool {
	return AlmostEqualD64(d, x...)
}

// Equal64 checks exact equality of all given 64-bit floating points values by comparing.
// It returns true if all values are equal.
//
// Special cases are:
//	Equal64() = false
//	Equal64(x) = true
//	Equal64(NaN) = false
//	Equal64(NaN, x) = false
//	Equal64(x, NaN) = false
//	Equal64(+Inf, +Inf) = true
//	Equal64(-Inf, -Inf) = true
func Equal64(x ...float64) bool {
	if len(x) <= 0 {
		return false
	}
	var a float64
	for i, b := range x {
		if math.IsNaN(b) {
			return false
		}
		if i > 0 {
			if !math.IsNaN(a-b) && (a > b || a < b) {
				return false
			}
		}
		a = b
	}
	return true
}

// Equal32 checks exact equality of all given 32-bit floating points values by comparing.
// It returns true if all values are equal.
//
// Special cases are:
//	Equal32() = false
//	Equal32(x) = true
//	Equal32(NaN) = false
//	Equal32(NaN, x) = false
//	Equal32(x, NaN) = false
//	Equal32(+Inf, +Inf) = true
//	Equal32(-Inf, -Inf) = true
func Equal32(x ...float32) bool {
	if len(x) <= 0 {
		return false
	}
	var a float32
	for i, b := range x {
		if math.IsNaN(float64(b)) {
			return false
		}
		if i > 0 {
			if !math.IsNaN(float64(a-b)) && (a > b || a < b) {
				return false
			}
		}
		a = b
	}
	return true
}

// Equal is synonym with Equal64.
func Equal(x ...float64) bool {
	return Equal64(x...)
}

// IsZero checks whether fraction of the given value is zero.
// It may return true even exponential isn't zero.
func IsZero(x float64) bool {
	return math.Float64bits(x)<<12 == 0
}

// Zero returns zero floating point value by given sign.
//	-0.0 if sign <  0
//	+0.0 if sign is 0
//	+0.0 if sign >  0
func Zero(sign int) float64 {
	return math.Copysign(0, float64(sign))
}

// Sign returns:
//	-1 if x <   0
//	 0 if x is ±0
//	+1 if x >   0
// Sign panics if x is NaN.
func Sign(x float64) int {
	if math.IsNaN(x) {
		panicForNaN(x)
	}
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	}
	return 0
}

// SignInt returns:
//	-1 if x <   0
//	 0 if x is  0
//	+1 if x >   0
func SignInt(x int64) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	}
	return 0
}

// Sum returns the sum of x...
func Sum(x ...float64) (sum float64) {
	for _, y := range x {
		sum += y
	}
	return
}

// Avg returns the arithmetic mean of x...
func Avg(x ...float64) (avg float64) {
	k := float64(len(x))
	for _, y := range x {
		avg += y / k
	}
	return
}

// SumInt returns the floating point of sum of x...
func SumInt(x ...int64) (sum float64) {
	for _, y := range x {
		sum += float64(y)
	}
	return
}

// AvgInt returns the floating point of arithmetic mean of x...
func AvgInt(x ...int64) (avg float64) {
	k := float64(len(x))
	for _, y := range x {
		avg += float64(y) / k
	}
	return
}

// SumUint returns the floating point of sum of x...
func SumUint(x ...uint64) (sum float64) {
	for _, y := range x {
		sum += float64(y)
	}
	return
}

// AvgUint returns the floating point of arithmetic mean of x...
func AvgUint(x ...uint64) (avg float64) {
	k := float64(len(x))
	for _, y := range x {
		avg += float64(y) / k
	}
	return
}

// SumInt2 returns the sum of x...
// If the result overflows, it returns overflow is true.
func SumInt2(x ...int64) (sum int64, overflow bool) {
	var last int64
	for _, y := range x {
		sum += y
		signLast, signSum, signY := SignInt(last), SignInt(sum), SignInt(y)
		if !overflow && signLast != signSum && signLast == signY {
			overflow = true
		}
		last = sum
	}
	return
}

// AvgInt2 returns the arithmetic mean of x...
// If the result overflows, it returns overflow is true.
func AvgInt2(x ...int64) (avg int64, overflow bool) {
	var sum int64
	sum, overflow = SumInt2(x...)
	if count := len(x); count > 0 {
		avg = sum / int64(count)
	}
	return
}

// SumUint2 returns the sum of x...
// If the result overflows, it returns overflow is true.
func SumUint2(x ...uint64) (sum uint64, overflow bool) {
	var last uint64
	for _, y := range x {
		sum += y
		if !overflow && sum < last {
			overflow = true
		}
		last = sum
	}
	return
}

// AvgUint2 returns the arithmetic mean of x...
// If the result overflows, it returns overflow is true.
func AvgUint2(x ...uint64) (avg uint64, overflow bool) {
	var sum uint64
	sum, overflow = SumUint2(x...)
	if count := len(x); count > 0 {
		avg = sum / uint64(count)
	}
	return
}

func checkInvalidBase(base int) {
	if base < MinBase || base > MaxBase {
		panic("invalid base")
	}
}

func panicForNaN(x float64) {
	if math.IsNaN(x) {
		panic("NaN value")
	}
}
