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
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
