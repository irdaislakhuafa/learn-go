package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		// closure: function stored in variable
		line("create closure function")
		getMinMax := func(values ...int) (min, max int) {
			for i, v := range values {
				switch {
				case i == 0:
					min, max = v, v
				case v > max:
					max = v
				case v < min:
					min = v
				}
			}
			return
		}

		values := []int{1, 2, 3, 4, 5, 6, 7}
		min, max := getMinMax(values...)
		fmt.Printf("values: %v\n", values)
		fmt.Printf("min: %v\n", min)
		fmt.Printf("max: %v\n", max)
	}
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}

func line2(v any) {
	s := strings.Repeat("-", 35)
	fmt.Println(s)
	fmt.Println(v)
	fmt.Println(s)
}
