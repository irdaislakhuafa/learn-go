package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	{
		line("hello world")
		diameter := 15.0
		// without prefined return variable
		calculate := func(v float64) (float64, float64) {
			// hitung luas
			area := math.Pi * math.Pow(v/2, 2)

			// hitung keliling
			circumference := math.Pi * v

			// return all value
			return area, circumference
		}
		calculate(diameter)

		// with predefined return variable
		calculate = func(v float64) (area float64, circumference float64) {
			area = math.Pi * math.Pow(v/2, 2)
			circumference = math.Pi * v
			return area, circumference
		}
		// and they are same

		area, circumference := calculate(diameter)
		fmt.Printf("Luas Lingkaran\t\t: %.2f\n", area)
		fmt.Printf("Keliling Lingkaran\t: %.2f\n", circumference)
	}
}
func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
