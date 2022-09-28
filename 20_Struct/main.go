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

		utils.Line("create 3 struct variable with different way")
		// 1
		v1 := student{}
		v1.name = "me"
		v1.grade = 1

		// 2
		v2 := student{"you", 2}

		// 3
		v3 := student{
			name:  "we",
			grade: 3,
		}

		fmt.Printf("v1: %+v\n", v1)
		fmt.Printf("v2: %+v\n", v2)
		fmt.Printf("v3: %+v\n", v3)
	}
}
