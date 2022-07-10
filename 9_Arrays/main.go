package main

import (
	"fmt"
	"strings"
)

func main() {
	// create array of `string`
	line(`create array of "string"`) // index start from 0
	var names [3]string
	names[0] = "irda"
	names[1] = "ana"
	names[2] = "kita"
	// names[3] = "someone" // error because slot of array is 3

	fmt.Printf("name 1: %s, name 2 : %s, name 3 = %s\n", names[0], names[1], names[2])

	// create array with assign values (horizontal style)
	line(`create array with assign values`)
	fruits := [4]string{"banana", "apple", "manggo", "lecy"}
	fmt.Println("Total Array \t:", len(fruits))
	fmt.Printf("All Array \t: %v\n", fruits)

	// initialze array with vertical style
	fruits = [4]string{
		"banana",
		"apple",
		"manggo",
		"lecy",
	}

	// create array without telling maximum value
	line(`create array without telling maximum value`)
	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Total Array \t:", len(numbers))
	fmt.Println("All Arrays\t:", numbers)
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
