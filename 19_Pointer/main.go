package main

import (
	"fmt"

	. "github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	{
		printValueAndAddress := func(a *int, b **int) {
			fmt.Printf("numberA (Value)\t\t: %v\n", *a)
			fmt.Printf("numberA (Address)\t: %v\n", a)
			fmt.Printf("numberB (Value)\t\t: %v\n", **b)
			fmt.Printf("numberB (Address)\t: %v\n", *b)
		}

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

	{
		Line("pointer as parameter")
		change := func(current *int, nextValue int) {
			*current = nextValue
		}

		number := 4
		fmt.Printf("Before: %v\n", number)

		change(&number, 10)
		fmt.Printf("After: %v\n", number)
	}
}
