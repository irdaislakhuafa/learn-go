package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/irdaislakhuafa/learn-go/25_Reflect/my_imagination"
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
		utils.Line("learn reflect")
		type Num struct {
			No int
		}
		number := Num{
			No: 20,
		}
		reflectValue := reflect.ValueOf(number)

		fmt.Printf("number: %v | type: \"%v\"\n", number, reflectValue.Type())

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

		// just try
		type details struct {
			Hobbies string
			Eye     int
			Address struct {
				Regency string
			}
		}
		type human struct {
			Name    string
			Age     int
			Details details
		}

		h1 := human{
			Name: "me",
			Age:  20,
			Details: details{
				Hobbies: "oioi",
				Eye:     2,
				Address: struct{ Regency string }{
					Regency: "Tuban",
				},
			},
		}
		my_imagination.PrintValues(h1)

		// example again
		getMyAge := func() (int, time.Time, error) {
			layout := "02 January 2006"
			value := "01 January 2002"
			born, err := time.Parse(layout, value)
			if err != nil {
				return 0, time.Time{}, err
			}

			now := time.Now()
			age := now.Year() - born.Year()
			return age, born, nil
		}

		age, born, err := getMyAge()
		if err != nil {
			panic(err)
		}
		h := Human{Name: "Irda Islakhu Afa", Age: age, BornAt: born}
		h.printPropertiesInfo()

		// access method with reflect
		fmt.Println("")
		hRef := reflect.ValueOf(&h)
		hMeth := hRef.MethodByName("SetName")
		hMeth.Call([]reflect.Value{reflect.ValueOf("me")})
		h.printPropertiesInfo()
	}
}

type Human struct {
	Name   string
	Age    int
	BornAt time.Time
}

func (h *Human) printPropertiesInfo() {
	v := reflect.ValueOf(h)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("Name: %+v\n", t.Field(i).Name)
		fmt.Printf("Data Type: %+v\n", t.Field(i).Type)
		fmt.Printf("Value: %+v\n", v.Field(i).Interface())
		fmt.Println("")
	}
}

func (h *Human) SetName(name string) {
	h.Name = name
}
