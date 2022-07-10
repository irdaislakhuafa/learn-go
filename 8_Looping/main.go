package main

import (
	"fmt"
	"strings"
)

func main() {
	// looping with `for` example 1
	line(`general "for" style`)
	for i := 0; i < 5; i++ {
		fmt.Println("Number :", i)
	}

	// for with `while` style
	line(`"for" with "while" style`)
	i := 0
	for i < 5 {
		fmt.Println("Number :", i)
		i++
	}
}

func line(text string) {
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
