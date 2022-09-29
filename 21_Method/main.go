package main

import (
	"fmt"
	"strings"

	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	/*
	   method is a function attached to a struct
	*/

	{
		utils.Line("method implementation")
		s1 := &student{
			name: "me",
			age:  20,
		}
		fmt.Printf("s1: %+v\n", s1)

		s1.sayHello()

		fmt.Printf("nick name: %v\n", s1.getNameAt(2))
	}
}

type student struct {
	name string
	age  uint8
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s student) getNameAt(index int) string {
	if index > len(strings.Split(s.name, " ")) {
		return s.name
	}
	return strings.Split(s.name, " ")[(index - 1)]
}
