package main

import "fmt"

func main() {
	// constanta is immutable value
	const (
		myName string = "Irda Islakhu Afa"
	)

	// try to change "myName"
	// myName = "some one" // error

	fmt.Print("My name : ", myName+"\n") // not add space between params
	fmt.Println("My name :", myName)     // the are same
}
