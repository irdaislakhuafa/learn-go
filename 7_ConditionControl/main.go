package main

import (
	"fmt"
)

func main() {
	/*
	   Go Lang have 2 ways to control code condidions, they are `if else` an `switch`
	*/

	// use if else
	fmt.Print("Your Grade : ")
	if point := 20; point == 8 {
		fmt.Println("E")
	} else if point == 10 {
		fmt.Println("D")
	} else if point == 12 {
		fmt.Println("C")
	} else if point == 14 {
		fmt.Println("B")
	} else if point >= 16 {
		fmt.Println("A")
	}
}
