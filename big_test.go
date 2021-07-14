package xmath_test

import (
	"fmt"
	"math"
	"math/big"

	"github.com/goinsane/xmath"
)

func ExampleFloorBigFloat() {
	fmt.Printf("floor of %v: %v\n", 2.4, xmath.FloorBigFloat(big.NewFloat(2.4)))
	fmt.Printf("floor of %v: %v\n", 2.5, xmath.FloorBigFloat(big.NewFloat(2.5)))
	fmt.Printf("floor of %v: %v\n", -6.5, xmath.FloorBigFloat(big.NewFloat(-6.5)))
	fmt.Printf("floor of %v: %v\n", -6.6, xmath.FloorBigFloat(big.NewFloat(-6.6)))
	fmt.Printf("floor of %v: %v\n", 5, xmath.FloorBigFloat(big.NewFloat(5)))
	fmt.Printf("floor of %v: %v\n", -9, xmath.FloorBigFloat(big.NewFloat(-9)))

	// Output:
	// floor of 2.4: 2
	// floor of 2.5: 2
	// floor of -6.5: -7
	// floor of -6.6: -7
	// floor of 5: 5
	// floor of -9: -9
}

func ExampleCeilBigFloat() {
	fmt.Printf("ceil of %v: %v\n", 2.4, xmath.CeilBigFloat(big.NewFloat(2.4)))
	fmt.Printf("ceil of %v: %v\n", 2.5, xmath.CeilBigFloat(big.NewFloat(2.5)))
	fmt.Printf("ceil of %v: %v\n", -6.5, xmath.CeilBigFloat(big.NewFloat(-6.5)))
	fmt.Printf("ceil of %v: %v\n", -6.6, xmath.CeilBigFloat(big.NewFloat(-6.6)))
	fmt.Printf("ceil of %v: %v\n", 5, xmath.CeilBigFloat(big.NewFloat(5)))
	fmt.Printf("ceil of %v: %v\n", -9, xmath.CeilBigFloat(big.NewFloat(-9)))

	// Output:
	// ceil of 2.4: 3
	// ceil of 2.5: 3
	// ceil of -6.5: -6
	// ceil of -6.6: -6
	// ceil of 5: 5
	// ceil of -9: -9
}

func ExampleRoundBigFloat() {
	fmt.Printf("round of %v: %v\n", 2.4, xmath.RoundBigFloat(big.NewFloat(2.4)))
	fmt.Printf("round of %v: %v\n", 2.5, xmath.RoundBigFloat(big.NewFloat(2.5)))
	fmt.Printf("round of %v: %v\n", -6.5, xmath.RoundBigFloat(big.NewFloat(-6.5)))
	fmt.Printf("round of %v: %v\n", -6.6, xmath.RoundBigFloat(big.NewFloat(-6.6)))
	fmt.Printf("round of %v: %v\n", 5, xmath.RoundBigFloat(big.NewFloat(5)))
	fmt.Printf("round of %v: %v\n", -9, xmath.RoundBigFloat(big.NewFloat(-9)))

	// Output:
	// round of 2.4: 2
	// round of 2.5: 3
	// round of -6.5: -6
	// round of -6.6: -7
	// round of 5: 5
	// round of -9: -9
}

func ExampleIntBigRat() {
	var n *big.Int
	var acc big.Accuracy
	n, acc = xmath.IntBigRat(big.NewRat(24, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", 24, 10, n, acc)
	n, acc = xmath.IntBigRat(big.NewRat(25, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", 25, 10, n, acc)
	n, acc = xmath.IntBigRat(big.NewRat(-65, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", -65, 10, n, acc)
	n, acc = xmath.IntBigRat(big.NewRat(-66, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", -66, 10, n, acc)
	n, acc = xmath.IntBigRat(big.NewRat(50, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", 50, 10, n, acc)
	n, acc = xmath.IntBigRat(big.NewRat(-90, 10))
	fmt.Printf("integer of %v/%v: %v acc=%v\n", -90, 10, n, acc)

	// Output:
	// integer of 24/10: 2 acc=Below
	// integer of 25/10: 2 acc=Below
	// integer of -65/10: -6 acc=Above
	// integer of -66/10: -6 acc=Above
	// integer of 50/10: 5 acc=Exact
	// integer of -90/10: -9 acc=Exact
}

func ExampleInt64BigRat() {
	var k int64
	var acc big.Accuracy
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(big.NewInt(50), big.NewInt(10)))
	fmt.Printf("int64 of 50/10: %v acc=%v\n", k, acc)
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(big.NewInt(-90), big.NewInt(10)))
	fmt.Printf("int64 of -90/10: %v acc=%v\n", k, acc)
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(big.NewInt(math.MaxInt64), big.NewInt(5)))
	fmt.Printf("int64 of MaxInt64/5: %v acc=%v\n", k, acc)
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(big.NewInt(math.MinInt64), big.NewInt(5)))
	fmt.Printf("int64 of MinInt64/5: %v acc=%v\n", k, acc)
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(new(big.Int).Mul(big.NewInt(7), big.NewInt(math.MaxInt64)), big.NewInt(5)))
	fmt.Printf("int64 of 7*MaxInt64/5: %v acc=%v\n", k, acc)
	k, acc = xmath.Int64BigRat(new(big.Rat).SetFrac(new(big.Int).Mul(big.NewInt(7), big.NewInt(math.MinInt64)), big.NewInt(5)))
	fmt.Printf("int64 of 7*MinInt64/5: %v acc=%v\n", k, acc)

	// Output:
	// int64 of 50/10: 5 acc=Exact
	// int64 of -90/10: -9 acc=Exact
	// int64 of MaxInt64/5: 1844674407370955161 acc=Below
	// int64 of MinInt64/5: -1844674407370955161 acc=Above
	// int64 of 7*MaxInt64/5: 9223372036854775807 acc=Below
	// int64 of 7*MinInt64/5: -9223372036854775808 acc=Above
}

func ExampleUint64BigRat() {
	var k uint64
	var acc big.Accuracy
	k, acc = xmath.Uint64BigRat(new(big.Rat).SetFrac(big.NewInt(50), big.NewInt(10)))
	fmt.Printf("uint64 of 50/10: %v acc=%v\n", k, acc)
	k, acc = xmath.Uint64BigRat(new(big.Rat).SetFrac(big.NewInt(0), big.NewInt(5)))
	fmt.Printf("uint64 of 0/5: %v acc=%v\n", k, acc)
	k, acc = xmath.Uint64BigRat(new(big.Rat).SetFrac(big.NewInt(math.MaxInt64), big.NewInt(7)))
	fmt.Printf("uint64 of MaxInt64/7: %v acc=%v\n", k, acc)
	k, acc = xmath.Uint64BigRat(new(big.Rat).SetFrac(new(big.Int).Mul(big.NewInt(13), big.NewInt(math.MaxInt64)), big.NewInt(5)))
	fmt.Printf("uint64 of 13*MaxInt64/5: %v acc=%v\n", k, acc)

	// Output:
	// uint64 of 50/10: 5 acc=Exact
	// uint64 of 0/5: 0 acc=Exact
	// uint64 of MaxInt64/7: 1317624576693539401 acc=Exact
	// uint64 of 13*MaxInt64/5: 18446744073709551615 acc=Below
}

func ExampleFloorBigRat() {
	fmt.Printf("floor of %v/%v: %v\n", 24, 10, xmath.FloorBigRat(big.NewRat(24, 10)))
	fmt.Printf("floor of %v/%v: %v\n", 25, 10, xmath.FloorBigRat(big.NewRat(25, 10)))
	fmt.Printf("floor of %v/%v: %v\n", -65, 10, xmath.FloorBigRat(big.NewRat(-65, 10)))
	fmt.Printf("floor of %v/%v: %v\n", -66, 10, xmath.FloorBigRat(big.NewRat(-66, 10)))
	fmt.Printf("floor of %v/%v: %v\n", 50, 10, xmath.FloorBigRat(big.NewRat(50, 10)))
	fmt.Printf("floor of %v/%v: %v\n", -90, 10, xmath.FloorBigRat(big.NewRat(-90, 10)))

	// Output:
	// floor of 24/10: 2
	// floor of 25/10: 2
	// floor of -65/10: -7
	// floor of -66/10: -7
	// floor of 50/10: 5
	// floor of -90/10: -9
}

func ExampleCeilBigRat() {
	fmt.Printf("ceil of %v/%v: %v\n", 24, 10, xmath.CeilBigRat(big.NewRat(24, 10)))
	fmt.Printf("ceil of %v/%v: %v\n", 25, 10, xmath.CeilBigRat(big.NewRat(25, 10)))
	fmt.Printf("ceil of %v/%v: %v\n", -65, 10, xmath.CeilBigRat(big.NewRat(-65, 10)))
	fmt.Printf("ceil of %v/%v: %v\n", -66, 10, xmath.CeilBigRat(big.NewRat(-66, 10)))
	fmt.Printf("ceil of %v/%v: %v\n", 50, 10, xmath.CeilBigRat(big.NewRat(50, 10)))
	fmt.Printf("ceil of %v/%v: %v\n", -90, 10, xmath.CeilBigRat(big.NewRat(-90, 10)))

	// Output:
	// ceil of 24/10: 3
	// ceil of 25/10: 3
	// ceil of -65/10: -6
	// ceil of -66/10: -6
	// ceil of 50/10: 5
	// ceil of -90/10: -9
}

func ExampleRoundBigRat() {
	fmt.Printf("round of %v/%v: %v\n", 24, 10, xmath.RoundBigRat(big.NewRat(24, 10)))
	fmt.Printf("round of %v/%v: %v\n", 25, 10, xmath.RoundBigRat(big.NewRat(25, 10)))
	fmt.Printf("round of %v/%v: %v\n", -65, 10, xmath.RoundBigRat(big.NewRat(-65, 10)))
	fmt.Printf("round of %v/%v: %v\n", -66, 10, xmath.RoundBigRat(big.NewRat(-66, 10)))
	fmt.Printf("round of %v/%v: %v\n", 50, 10, xmath.RoundBigRat(big.NewRat(50, 10)))
	fmt.Printf("round of %v/%v: %v\n", -90, 10, xmath.RoundBigRat(big.NewRat(-90, 10)))

	// Output:
	// round of 24/10: 2
	// round of 25/10: 3
	// round of -65/10: -6
	// round of -66/10: -7
	// round of 50/10: 5
	// round of -90/10: -9
}
