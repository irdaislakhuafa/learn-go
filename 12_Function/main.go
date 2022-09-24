package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		printMessage := func(text string, arr ...string) {
			fmt.Println(text + " " + strings.Join(arr, " "))
		}

		name := []string{"irda", "islakhu", "afa"}
		printMessage("hi!", name...)
	}
}
