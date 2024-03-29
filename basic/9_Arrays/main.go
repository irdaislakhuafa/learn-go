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

	// array multidimension
	line(`array multidimension`)
	numbers1 := [2][3]int{[3]int{1, 2, 3}, [...]int{4, 5, 6}} // example 1
	numbers2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}               // example 2
	fmt.Printf("Numbers 1 \t: %d\n", numbers1)
	fmt.Println("Numbers 2 \t:", numbers2)
	// they are equals

	// looping array with `for` keyword
	line(`looping array with "for" keyword`)
	for i := 0; i < len(fruits); i++ {
		fmt.Printf(`Element %d : "%s" %s`, i, fruits[i], "\n")
	}

	// looping array with `for - range` keyword
	line(`with "for - range"`)
	for i, fruit := range fruits {
		fmt.Printf("Element %d : \"%s\"\n", i, fruit)
	}

	// use `_` variable in `for`
	line(`use "_" variable in "for"`)
	for _, fruit := range fruits {
		fmt.Printf("Element : \"%s\"\n", fruit)
	}

	// use index only
	line(`use index only`)
	for i := range fruits { // or for i := range fruits
		fmt.Println("Index :", i)
	}

	// create array with keyword `make`
	line(`create array with keyword "make"`)
	genders := make([]string, 2)
	genders[0] = "male"
	genders[1] = "female"

	fmt.Println("Genders:", genders)
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
