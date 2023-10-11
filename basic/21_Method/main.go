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

	{
		utils.Line("method with pointer")
		/*
		   - method without pointer cat change value but only in current method and not change value of reference
		   - method with reference can change of reference
		*/
		s1 := &student{
			name: "me",
			age:  20,
		}

		fmt.Printf("s1: %+v\n", s1)
		s1.setNameWithoutPointer("you") // me
		fmt.Printf("s1: %+v\n", s1)

		s1.setNameWithPointer("you") // you
		fmt.Printf("s1: %+v\n", s1)
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

func (s student) setNameWithoutPointer(name string) {
	v := strings.Repeat(">", 3)
	fmt.Println(v + " change \"" + s.name + "\" " + v + " \"" + name + "\"")
	s.name = name
}

func (s *student) setNameWithPointer(name string) {
	v := strings.Repeat(">", 3)
	fmt.Println(v + " change \"" + s.name + "\" " + v + " \"" + name + "\"")
	s.name = name
}
