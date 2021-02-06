package sql

import (
	"fmt"
	"testing"
)

func TestIdUtils(t *testing.T) {
	var a float32 = 123
	var b float64 = 123
	var e int8 = 123
	var f int16 = 123
	var g uint8 = 123
	var c complex128 = 123
	fmt.Println(ConstituteWherePhrase("qq", EQ, "123"))
	fmt.Println(ConstituteWherePhrase("qq", GE, a))
	fmt.Println(ConstituteWherePhrase("qq", GT, b))
	fmt.Println(ConstituteWherePhrase("qq", LT, e))
	fmt.Println(ConstituteWherePhrase("qq", LE, f))
	fmt.Println(ConstituteWherePhrase("qq", IS_NOT_NULL, g))
	fmt.Println(ConstituteWherePhrase("qq", IS_NULL, c))
}
