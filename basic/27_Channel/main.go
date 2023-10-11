package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c <- "error"
		// c <- "errors"
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(len(c))
	fmt.Println(<-c)
	fmt.Println(len(c))
	fmt.Println(<-c)
	fmt.Println("Hello World")
}
