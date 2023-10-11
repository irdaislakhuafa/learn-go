package main

import (
	"fmt"
	"runtime"

	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	utils.Line("goroutine implementation")
	{
		// maximum cpu used to run this code
		runtime.GOMAXPROCS(4)
		print := func(until int, message string) {
			for i := 0; i < until; i++ {
				fmt.Println((i + 1), message)
			}
		}

		// run print() method with goroutine (mini thread)
		go print(10, "hello")

		// run print() method with main thread
		print(10, "world")

		// to hold main thread close before user press enter, because maybe main thread is done before goroutine and goroutine will be stoped after main thread is done
		fmt.Scanln(new(string))
	}
}
