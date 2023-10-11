package main

import (
	"fmt"
)

func main() {
	// here i learn about how to declare variable in Go Language

	// 1st way
	var firstName string = "irda"
	fmt.Printf(`My Name : %+s `, firstName)

	// 2st way
	var middleName string
	middleName = "islakhu"
	fmt.Print(middleName, " ")

	// 3st way
	var lastName = "afa"
	fmt.Println(lastName)

	// 4st way
	meSkills, age := "Java and Go Language", 20
	fmt.Println("My Skills :", meSkills)
	fmt.Println("My Age :", age)

	// create pointer variable with new() keyword
	gf := new(string)
	*gf = "ana ardani"
	fmt.Println("GF :", *gf)

	// create variable with make() keyword
	// only can be used in (slice, array, map, channel)
	hobbies1 := make(map[any]any)
	hobbies2 := map[any]any{} // same like above
	var a map[any]any         // nil/unalocated value

	hobbies1["irda"] = "Mengoding"
	hobbies2["someone"] = "Menyanyi"
	// a["anything"] = "anything" //panic
	a = map[any]any{}
	a["anything"] = "anything" //success

	fmt.Printf("Hobbies 1: %+v\n", hobbies1)
	fmt.Printf("Hobbies 2: %+v\n", hobbies2)
	fmt.Printf("anything: %+v\n", a)
}
