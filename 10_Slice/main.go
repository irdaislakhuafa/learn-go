package main

import (
	"fmt"
	"strings"
)

func main() {
	line(`initialize slice and array`)
	sliceFruits := []string{"banana", "apple"}
	arrayFruits := [...]string{"banana", "apple"}

	// try use "append()"
	sliceFruits = append(sliceFruits, "lecy")
	// arrayFruits = append(arrayFruits, "lecy") // can't use "append()" in array

	fmt.Println("Slice Fruits:", sliceFruits)
	fmt.Println("Array Fruits:", arrayFruits)
	fmt.Println("Slice Fruits[0]:", sliceFruits[0])
	fmt.Println("Array Fruits[0]:", arrayFruits[0])
	fmt.Println("Slice Fruits[1]:", sliceFruits[1])
	fmt.Println("Array Fruits[1]:", arrayFruits[1])
	fmt.Println("Slice Fruits[2]:", sliceFruits[2]) // don't have error notif but panic while code is running
	// fmt.Println("Array Fruits[0]:", arrayFruits[2]) // have error notif
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
