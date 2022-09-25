package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	{
		line("function implementation")
		printMessage := func(text string, arr ...string) {
			fmt.Println(text + " " + strings.Join(arr, " "))
		}

		name := []string{"irda", "islakhu", "afa"}
		printMessage("hi!", name...)
	}

	{
		line("function with return value")
		// to make generated number with randome mode
		rand.Seed(time.Now().UnixMicro())

		// we can define data type 1x in function parameter if they are same
		genRandomNumberWithRange := func(min, max int64) int64 {
			return int64(rand.Int())%(max-min+1) + min
		}

		fmt.Println("Random Number With Range:", genRandomNumberWithRange(1, 3))

	}
}

func line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}

func line2(v any) {
	s := strings.Repeat("-", 35)
	fmt.Println(s)
	fmt.Println(v)
	fmt.Println(s)
}
