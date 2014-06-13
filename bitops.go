package bitops

import (
	"fmt"
	"math"
	"strconv"
)

func SetBit(integer *uint64, position int, value bool) {
	var mask uint64
	if value {
		mask = 1 << uint64(position)
		*integer = *integer | uint64(mask)
		return
	} else {
		mask = uint64(math.MaxUint64) - uint64(math.Pow(2, float64(position)))
		*integer = *integer & mask
	}
}

func QueryBit(integer *uint64, position int) bool {
	posUint := uint(position)
	return (*integer>>posUint)&1 > 0
}

func PrintUint64(integer *uint64) {
	fmt.Println(strconv.FormatUint(*integer, 2))
}

func Print8x8Uint64(integer *uint64) {
	stringInt := strconv.FormatUint(*integer, 2)
	for i := len(stringInt); i < 64; i++ {
		stringInt = "0" + stringInt
	}

	fmt.Println(stringInt[0:7])
	fmt.Println(stringInt[8:15])
	fmt.Println(stringInt[16:23])
	fmt.Println(stringInt[24:31])
	fmt.Println(stringInt[32:39])
	fmt.Println(stringInt[40:47])
	fmt.Println(stringInt[48:55])
	fmt.Println(stringInt[56:63])
}
