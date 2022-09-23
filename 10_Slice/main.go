package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		line(`initialize slice and array`)
		sliceFruits := []string{"banana", "apple"}    // slice
		arrayFruits := [...]string{"banana", "apple"} // array

		// try use "append()"
		// sliceFruits = append(sliceFruits, "lecy") // can use "append()" in slice
		// arrayFruits = append(arrayFruits, "lecy") // can't use "append()" in array

		fmt.Println("Slice Fruits:", sliceFruits)
		fmt.Println("Array Fruits:", arrayFruits)
		fmt.Println("Slice Fruits[0]:", sliceFruits[0])
		fmt.Println("Array Fruits[0]:", arrayFruits[0])
		fmt.Println("Slice Fruits[1]:", sliceFruits[1])
		fmt.Println("Array Fruits[1]:", arrayFruits[1])
		// fmt.Println("Slice Fruits[2]:", sliceFruits[2]) // don't have error notif but panic while code is running
		// fmt.Println("Array Fruits[0]:", arrayFruits[2]) // have error notif
	}

	{
		line(`slice and array relationship and slice operation`)
		arrayNames := [...]string{"me", "you"}
		sliceOfArrayNames := arrayNames[0:2] // 2 index technique return slice

		fmt.Printf("arrayNames (Array): %v\n", arrayNames)
		fmt.Printf("sliceOfArrayNames (Slice): %v\n\n", sliceOfArrayNames)

		fmt.Printf("sliceOfArrayNames[0:1]: %v\n", sliceOfArrayNames[0:1]) // ["me"]
		fmt.Printf("sliceOfArrayNames[0:2]: %v\n", sliceOfArrayNames[0:2]) // ["me", "you"]
		fmt.Printf("sliceOfArrayNames[0:0]: %v\n", sliceOfArrayNames[0:0]) // []
		fmt.Printf("sliceOfArrayNames[1:1]: %v\n", sliceOfArrayNames[1:1]) // []
		// fmt.Printf("sliceOfArrayNames[1:0]: %v\n", sliceOfArrayNames[1:0]) // first index must be lower than second index
		fmt.Printf("sliceOfArrayNames[:]: %v\n", sliceOfArrayNames[:])   // ["me", "you"] / all elements
		fmt.Printf("sliceOfArrayNames[0:]: %v\n", sliceOfArrayNames[0:]) // ["me", "you"] / all elements start from index 0 to end index
		fmt.Printf("sliceOfArrayNames[:1]: %v\n", sliceOfArrayNames[:1]) // ["me"] / all elements from index 0 until before index 1
	}

	{
		line(`slice is reference data type (pointers of array)`)
		peoples := []string{"me", "you", "they", "we"}

		aPeoples := peoples[0:3] // ["me", "you", "they"]
		bPeoples := peoples[1:4] // ["you", "they", "we"]

		aaPeoples := peoples[1:2] // ["you"]
		bbPeoples := peoples[0:1] // ["me"]

		fmt.Println(`before change some value in slice`)
		fmt.Println(strings.Repeat("-", 35))

		fmt.Printf("aPeoples: %v\n", aPeoples)
		fmt.Printf("bPeoples: %v\n", bPeoples)

		fmt.Printf("aaPeoples: %v\n", aaPeoples)
		fmt.Printf("bbPeoples: %v\n", bbPeoples)

		// change some value
		aaPeoples[0] = "someone"

		fmt.Println(`after change some value in slice`)
		fmt.Println(strings.Repeat("-", 35))

		fmt.Printf("aPeoples: %v\n", aPeoples)
		fmt.Printf("bPeoples: %v\n", bPeoples)

		fmt.Printf("aaPeoples: %v\n", aaPeoples)
		fmt.Printf("bbPeoples: %v\n", bbPeoples)

		/*
			all new slice from old slice will have same effect while we change some value from old slice of new slice because slice is reference data type (ponter) of array
		*/
	}

	{
		line(`len() function`) // to get length of array/slice
		names := []string{"me", "you"}
		fmt.Println("names:", names)
		fmt.Println("Length of names:", len(names))
	}

	{
		line(`cap() function`) // get total capcity of slice
		names := []string{"me", "you", "they", "we"}

		fmt.Println("len names:", len(names), "expected: 4") // 4
		fmt.Println("cap names:", cap(names), "expected: 4") // 4

		tempNames := names[0:3] // ["me", "you", "they"]

		fmt.Println("len tempNames:", len(tempNames), "expected: 3") // 3
		fmt.Println("cap tempNames:", cap(tempNames), "expected: 4") // 4

		tempNames = names[1:4]

		fmt.Println("len tempNames:", len(tempNames), "expected: 3") // 3
		fmt.Println("cap tempNames:", cap(tempNames), "expected: 3") // 4

		/*
			code		|	value						| len()	| cap()
			names[0:4]	|	["me", "you", "they", "we"]	|	4	|	4
			names[0:3]	|	["me", "you", "they", --]	|	3	|	4
			names[1:4]	|	--- ["you", "they", "we"]	|	3	|	3
		*/
	}
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}
