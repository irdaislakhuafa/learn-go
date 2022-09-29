package main

import (
	"fmt"

	meIsAlias "github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/aliases"
	. "github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/dotprefix"
	"github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/initstruct"
	"github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/library"
	"github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/structs"
	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	{
		utils.Line("call private/unexported and public/exported function in Go")
		library.SayHello()
		// library.introduce("me") // cannot be call because this is private function

		utils.Line("call unexported/private method from exported/public method")
		library.SayHelloWithIntroduce("me")
	}

	{
		utils.Line("use exported/public and unexported/private access")
		s1 := &structs.Student{
			Name:  "me",
			Grade: 1,
		}
		fmt.Printf("s1: %+v\n", s1)
	}

	{
		utils.Line("import with dot(.) prefix")
		SaySomething()
	}

	{
		utils.Line("use alias when import package")
		meIsAlias.MyAlias()
	}

	{
		/*
			when run go build or go run all files with package main must be call in the argument like this "go run main.go partial.go"
		*/
		utils.Line("access property in file with same package")
		sayHelloFromPartial()
	}

	{
		/*
			init() function will be call first time when app running, is this function attached to main package, init() method will be call before main() method
		*/
		utils.Line("init() function")
		s1 := initstruct.StudentInit

		fmt.Printf("s1.Name: %v\n", s1.Name)
		fmt.Printf("s1.Grade: %v\n", s1.Grade)
	}
}
