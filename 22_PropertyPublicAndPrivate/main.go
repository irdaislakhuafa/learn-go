package main

import (
	"github.com/irdaislakhuafa/learn-go/22_PropertyPublicAndPrivate/library"
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
}
