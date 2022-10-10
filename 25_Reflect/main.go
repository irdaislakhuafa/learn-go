package main

import (
	"fmt"
	"reflect"

	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	/*
		reflect is variable inspection technique, used to get information from a variable and manipulate a variable
		reflect have many function but 2 more important function
		- reflect.ValueOf() : will return object with type reflect.Value, contains all information about value of variable
		- reflect.TypeOf() : will return object with type reflect.Type, contains all information about type of variable
	*/
	{
		utils.Line(" ")
		number := struct {
			No int
		}{
			No: 20,
		}
		reflectValue := reflect.ValueOf(number)

		fmt.Printf("number: %v | type: \"%v\"\n", number, reflectValue.Type().Name())

		if reflectValue.Kind() == reflect.Int {
			fmt.Println("value is:", reflectValue.Int())
		}

		if reflectValue.Kind() == reflect.String {
			fmt.Println("value is:", reflectValue.String())
		}

		if reflectValue.Kind() == reflect.Bool {
			fmt.Println("value is:", reflectValue.Bool())
		}

		if reflectValue.Kind() == reflect.Struct {
			fmt.Printf("value is: %+v\n", reflectValue)
		}
	}
}
