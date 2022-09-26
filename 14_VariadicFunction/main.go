package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		// with variadic function
		average := func(values ...int) float64 {
			total := 0
			for _, v := range values {
				total += v
			}
			return float64(total) / float64(len(values))
		}

		avg := average(1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Printf("avg: %.2f\n", avg)

		// with slice
		averageWithSlice := func(values []int) float64 {
			total := 0
			for _, v := range values {
				total += v
			}

			return float64(total) / float64(len(values))
		}

		avg = averageWithSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		fmt.Printf("avg: %.2f\n", avg)
	}
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
