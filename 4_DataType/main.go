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

	// boolean data type (true/false)
	isExists := true
	fmt.Printf("Is Exists?: %t\n", isExists)
	fmt.Println(strings.Repeat("=", int(repeatLine)))

	// string data type
	message1 := "Hi this string use 'quote' \n this is endline"
	message2 := `Hi this string use 'backticks'
 this is endline
`

	fmt.Printf("String with \"Quote\": %s\n", message1)
	fmt.Printf(`String with "Backticks": %s`, message2)

	zeroValue := `
	Zero Value
	%s
	string  : ""
	bool 	: false
	int 	: 0
	fload	: 0.0
	map/slice/interface{}/channel/pointers/func : nil
	%s
	`
	fmt.Printf(zeroValue, strings.Repeat("=", int(repeatLine)), strings.Repeat("=", int(repeatLine))+"\n")
}
