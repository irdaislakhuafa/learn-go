package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		line("closure as return")
		findMax := func(numbers []int, maximum int) (int, func() []int) {
			res := []int{}
			for _, v := range numbers {
				if v <= maximum {
					res = append(res, v)
				} else {
					continue
				}
			}

			return len(res), func() []int { return res }
		}

		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		totalElement, getNumbers := findMax(values, 4)

		fmt.Printf("values: %v\n", values)
		fmt.Printf("Maximum is 4\n")
		fmt.Printf("totalElement: %v\n", totalElement)
		fmt.Printf("getNumbers(): %v\n", getNumbers())
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
