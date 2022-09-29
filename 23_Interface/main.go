package main

import (
	"fmt"

	"github.com/irdaislakhuafa/learn-go/23_Interface/examinterface"
	"github.com/irdaislakhuafa/learn-go/utils"
)

func main() {
	{
		utils.Line("interface implementation")
		var twoDimentionalFigure examinterface.Count

		twoDimentionalFigure = &examinterface.Rectangle{Side: 10.0}
		utils.Line2(fmt.Sprintf("rectangle: %+v", twoDimentionalFigure))
		fmt.Println("rectangle large\t\t:", twoDimentionalFigure.Large())
		fmt.Println("rectangle around\t:", twoDimentionalFigure.Around())

		twoDimentionalFigure = &examinterface.Circle{Diameter: 14.0}
		utils.Line2(fmt.Sprintf("circle: %+v", twoDimentionalFigure))
		fmt.Println("circle large\t\t:", twoDimentionalFigure.Large())
		fmt.Println("circle around\t\t:", twoDimentionalFigure.Around())
		fmt.Println("circle fingers\t\t:", twoDimentionalFigure.(*examinterface.Circle).Fingers())
	}

	{

	}
}
