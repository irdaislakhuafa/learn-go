package main

import (
	"fmt"
	"strings"
)

const (
	lineChar string = "="
	lineLength int8 = 35
)

func main() {
	// try use some operators in Go Language

	// aritmatics operators (+, -, /, *, %)
	section("aritmatics operators")
	increase := 2 + 2
	fmt.Println("2 + 2 :", increase)
	
	subtraction := 2 - 2
	fmt.Print("2 - 2 : ", subtraction, "\n")

	multiplication := 2 * 2
	fmt.Printf("2 * 2 : %d\n", multiplication)

	division := 2 / 2
	fmt.Print("2 / 2 : ", division, "\n")

	modulus := 2 % 2
	fmt.Println("2 % 2 :", modulus)

	// comparison operators
	section("comparison operators")
	isEquals := 2 == 2
	fmt.Printf("2 == 2 \t: %t\n", isEquals)

	isNotEquals := 2 != 2
	fmt.Print("2 != 2 \t: ", isNotEquals, "\n")

	isLowerThan := 2 < 2
	fmt.Println("2 <  2 \t:", isLowerThan)

	isLowerThanEquals := 2 <= 2
	fmt.Printf("2 <= 2 \t: %t\n", isLowerThanEquals)

	isGreaterThan := 2 > 2
	fmt.Print("2  > 2 \t: ", isGreaterThan, "\n")

	isGreaterThanEquals := 2 >= 2
	fmt.Println("2 >= 2 \t:", isGreaterThanEquals)

	// logic operators 
	section("logic operators")
	isTrue, isFalse := true, false
	
	and := isTrue && isFalse
	fmt.Printf("%t && %t : %t\n", isTrue, isFalse, and)
}

func section(text string) {
	fmt.Println("\n")
	line()
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	line()
}

func line() {
	fmt.Println(strings.Repeat(lineChar, int(lineLength)))
}
