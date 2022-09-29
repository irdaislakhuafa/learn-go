package main

import (
	"fmt"

	. "github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/dotprefix"
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
}
