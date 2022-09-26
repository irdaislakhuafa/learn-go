package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		line("with variadic function")
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

		line("with slice")
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

	{
		line("input variadic parameters with slice")
		calculate := func(values ...int) int64 {
			total := 0
			for _, v := range values {
				total += v
			}

			return int64(total)
		}

		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		total := calculate(slice...)
		fmt.Printf("total: %v\n", total)
	}

	{
		line("general parameter and variadic parameter")
		bio := func(name string, hobbies ...string) {
			fmt.Printf("name: %v\n", name)
			fmt.Printf("hobbies: %v\n", strings.Join(hobbies, ", "))
		}

		bio("irda", "ngoding", "makan", "tidur")

		line2("input with slice")
		hobbies := []string{"mengoding", "makan", "tidur"}
		bio("irda", hobbies...)
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
