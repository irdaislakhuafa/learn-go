package main

import (
	"fmt"
	"strings"

	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	/*
	   empty interface is a data type in go, like int, string, bool, etc.
	   empty interface can be used to store any data type like int, bool, string, struct and etc.
	*/

	{
		utils.Line("empty interface usage")
		var anyValue interface{}

		// store string value
		anyValue = "irda islakhu afa"
		fmt.Printf("anyValue (string)\t: %v\n", anyValue)

		// store array if string
		anyValue = []string{"never", "give", "up"}
		fmt.Printf("anyValue ([]string)\t: %v\n", strings.Join(anyValue.([]string), ", "))

		// store float value
		anyValue = 20.20
		fmt.Printf("anyValue (anyValue)\t: %f\n", anyValue)

		values := map[string]interface{}{
			"name":    "irda islakhu afa",
			"age":     20.10,
			"hobbies": []string{"ngoding", "makan", "tidur"},
		}

		utils.Line2("")
		fmt.Printf("values: %v\n", values)
	}

	{
		/*
			for Go version 18.1 or above we can use `any` keyword instead of interface{}`
		*/
		utils.Line(`use "any" instead of "interface{}"`)
		data := map[string]any{
			"name": "irda islakhu afa",
			"age":  20.10,
			"hobbies": []string{
				"ngoding",
				"makan",
				"tidur",
			},
		}
		fmt.Printf("data: %v\n", data)
	}
}
