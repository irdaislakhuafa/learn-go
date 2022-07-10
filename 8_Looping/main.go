package main

import (
	"fmt"
	"strings"
)

func main() {
	// looping with `for` example 1
	for i := 0; i < 5; i++ {
		fmt.Println("Number :", i)
	}
	line()

	// for with `while` style
	i := 0
	for i < 5 {
		fmt.Println("Number :", i)
	}
}

func line() {
	fmt.Println(strings.Repeat("=", 20))
}
