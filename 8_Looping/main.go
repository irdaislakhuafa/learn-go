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

	// for without arguments
	line(`"for" without arguments`)
	i = 0
	for {
		if i == 5 {
			break
		}
		fmt.Println("Number :", i)
		i++
	}

	// use `break` and `continue` keyword
	line(`use "break" and "continue" keyword`)
	for i := 0; i <= 10; i++ {
		if (i % 2) == 1 {
			continue // go to next loop with ignore code under this code
		}

		if i > 8 {
			break // out from loop
		}

		fmt.Printf("Number : %+d\n", i)
	}

	// create nested looping
	line("create nested loop")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Printf("%+d", j)
		}
		fmt.Printf("\n")
	}
}

func line(text string) {
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
