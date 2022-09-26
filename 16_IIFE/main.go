package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		line("Immediately-Invoked Function Expression (IIFE)")
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		minimumValues := func(min int) (res []int) {
			for _, value := range values {
				if value <= min {
					res = append(res, value)
				} else {
					continue
				}
			}
			return
		}(3)

		fmt.Printf("values: %v\n", values)
		fmt.Printf("minimumValues: %v\n", minimumValues)
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
