package initstruct

import "fmt"

var StudentInit = struct {
	Name  string
	Grade uint8
}{}

func init() {
	StudentInit.Name = "default name"
	StudentInit.Grade = 1

	fmt.Println("---> initstruct.go is imported")
}
