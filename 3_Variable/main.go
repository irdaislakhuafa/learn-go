package main

import "fmt"

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
}
