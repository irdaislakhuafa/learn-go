package main

import (
	"fmt"

	. "github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	Line("pointer implementation")
	// create pointer variables
	var number *int
	var name *string

	fmt.Printf("number: %v\n", number) // nil
	fmt.Printf("name: %v\n", name)     // nil

	Line2("implementation")
	var numberA int = 2
	var numberB *int = &numberA
	printValueAndAddress(&numberA, &numberB)

	Line("pointer value change effect")
	numberA = 10
	printValueAndAddress(&numberA, &numberB)
}

func printValueAndAddress(a *int, b **int) {
	fmt.Printf("numberA (Value)\t\t: %v\n", *a)
	fmt.Printf("numberA (Address)\t: %v\n", a)
	fmt.Printf("numberB (Value)\t\t: %v\n", **b)
	fmt.Printf("numberB (Address)\t: %v\n", *b)

}
