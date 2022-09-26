package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		line("function as parameter of function")
		filter := func(values []string, callback func(string) bool) (res []string) {
			for _, v := range values {
				if status := callback(v); status {
					res = append(res, v)
				}
			}
			return
		}

		names := []string{"me", "you", "they", "we"}
		fmt.Printf("names\t\t\t: %v\n", strings.Join(names, ", "))

		dataContainsE := filter(names, func(s string) bool { return strings.Contains(strings.ToLower(s), "e") })
		fmt.Printf("names containe \"e\"\t: %v\n", strings.Join(dataContainsE, ", "))

		dataLength2 := filter(names, func(s string) bool { return len(s) <= 2 })
		fmt.Printf("names with length 2\t: %v\n", strings.Join(dataLength2, ", "))
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
