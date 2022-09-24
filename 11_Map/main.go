package main

import (
	"fmt"
	"strings"
)

func main() {
	/* map is key-value data type like slice/array but we can define data type for key and value */
	{
		line("initialize map data type")
		line2("before initialized")
		var chicken map[string]int // the key data type is string and value data type is int
		fmt.Printf("chicken: %v\n", chicken)
		// chicken["one"] = 1 // error because nil value of map is nil
		// fmt.Printf("chicken: %v\n", chicken)

		line2("after initialized")
		chicken = map[string]int{}
		chicken["one"] = 1
		fmt.Printf("chicken[\"one\"]: %v\n", chicken["one"])

		line2("initialize map in first defition variable")
		ageOfPeoples := map[string]byte{"me": 20, "you": 21}
		fmt.Printf("ageOfPeoples: %v\n", ageOfPeoples)

		line2("initialize map without first value")
		mapWithoutFirstValue := map[string]byte{}
		fmt.Printf("mapWithoutFirstValue: %v\n", mapWithoutFirstValue)
	}

	{
		line("iteration with for-range in map")
		months := map[any]any{
			"januari":  1,
			"february": 2,
			"maret":    3,
			"april":    4,
			"mei":      5,
		}
		fmt.Printf("months: %v\n", months)
		for key, value := range months {
			fmt.Println("Key:", key, "-> value:", value)
		}
	}

	{
		line("delete key-value from map")
		ageOfPeoples := map[any]any{
			"me":  20,
			"you": 21,
		}
		fmt.Printf("ageOfPeoples: %v\n", ageOfPeoples)
		delete(ageOfPeoples, "me") // only canbe use in map data type
		fmt.Printf("ageOfPeoples: %v\n", ageOfPeoples)
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
