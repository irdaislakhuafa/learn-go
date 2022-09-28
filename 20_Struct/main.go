package main

import (
	"fmt"

	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	{
		/*
			struct like class in OOP language
		*/
		utils.Line("struct declaration")
		// create a struct
		type student struct {
			name  string
			grade int
		}

		utils.Line("struct implementation")
		student1 := student{
			name: "me",
		}
		student1.grade = 1

		fmt.Printf("student1.name\t: %v\n", student1.name)
		fmt.Printf("student1.grade\t: %v\n", student1.grade)
		fmt.Printf("student{}\t: %+v\n", student{})
	}
}
