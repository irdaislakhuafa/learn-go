package my_imagination

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	count = 0
)

func PrintValues(v any) {
	printValue(v, "=")
}

func printValue(v any, sep string) {
	const line = 20
	typeV := reflect.TypeOf(v)
	valueV := reflect.ValueOf(v)

	if sep != "" {
		fmt.Println(strings.Repeat("\t", count) + strings.Repeat(sep, line))
	}

	for i := 0; i < typeV.NumField(); i++ {

		switch typeV.Field(i).Type.String() {
		case "unsafe.Pointer":
			fallthrough
		case "reflect.flag":
			fallthrough
		case "*reflect.rtype":
			continue
		}

		fmt.Printf("%s%+v\t: ", strings.Repeat("\t", count), typeV.Field(i).Name)
		fmt.Printf("%+v\n", func() any {
			switch valueV.Field(i).Kind() {
			case reflect.Int:
				return valueV.Field(i).Int()
			case reflect.Bool:
				return valueV.Field(i).Bool()
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				return valueV.Field(i).Float()
			case reflect.String:
				return valueV.Field(i).String()
			case reflect.Struct:
				count++
				fmt.Println("")
				printValue(valueV.Field(i).Interface(), "-")
				count--
				return ""
			default:
				return ""
			}
		}())
	}

	if sep != "" {
		fmt.Println(strings.Repeat("\t", count) + strings.Repeat(sep, line))
	}
}
