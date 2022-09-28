package utils

import (
	"fmt"
	"strings"
)

func Line(text string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 35))
	fmt.Printf("\t%s\n", strings.ToUpper(text))
	fmt.Println(strings.Repeat("=", 35))
}

func Line2(v any) {
	s := strings.Repeat("-", 35)
	fmt.Println(s)
	fmt.Println(v)
	fmt.Println(s)
}
