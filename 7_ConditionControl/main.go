package main

import (
	"fmt"
)

func main() {
	/*
	   Go Lang have 2 ways to control code condidions, they are `if else` an `switch`
	*/

	// use if else (without temporary variable)
	fmt.Print("Your Grade : ")
	point := 20
	if point == 8 {
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

	// if else (with temporari variable)
	if tempVariableAge := 20; tempVariableAge == 20 {
		fmt.Printf("Your Age Is : %d\n", tempVariableAge)
	} else {
		fmt.Printf("Yout Age Is Not : %d\n", tempVariableAge)
	}

	// switch case (in golang, switch case is without break)
	point = 4
	switch point {
	case 2:
		fmt.Println("Your number is Two")
	case 4:
		fmt.Println("Your number is For")
	case 8:
		fmt.Println("Your number is Eight or Six")
	case 10:
		fmt.Println("Your number is Ten")
	default:
		fmt.Println("Not number between 1 - 10")
	}
}
