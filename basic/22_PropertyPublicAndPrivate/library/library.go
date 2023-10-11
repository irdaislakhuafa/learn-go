package library

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayHelloWithIntroduce(name string) {
	SayHello()
	introduce(name)
}

func introduce(name string) {
	fmt.Println("my name:", name)
}
