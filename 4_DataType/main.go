package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		uint8 = 0 => 255
		uit16 = 0 => 65535
		uit32 = 0 => 4294967295
		uint64 = 0 => 18446744073709551615
		uint = same as uint32 or uint64 depends value
		int8 = 128 => 127
		int16 = -32768 => 32767
		int32 = -2147483648 => 2147483647
		int64 = -9223372036854775808 => 9223372036854775807
		int = same as int32 or int64 depends value
		rune = same as int32

	*/

	const repeatLine uint8 = 20

	// numeric data type with non decimal value (Integer)
	var positiveNumber uint8 = 178
	var negativeNumber = -99999
	fmt.Printf("Positive Number: %+d\n", positiveNumber)
	fmt.Printf("Negative Number: %+d\n", negativeNumber)
	fmt.Println(strings.Repeat("=", int(repeatLine)))

	// numeric data type with decimal value (Float)
	var decimalNumber = 20.22
	fmt.Printf("Decimal Number: %f\n", decimalNumber)
	fmt.Printf("Decimal Number: %.3f\n", decimalNumber)
	fmt.Println(strings.Repeat("=", int(repeatLine)))
}
