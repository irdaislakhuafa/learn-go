package examinterface

import "math"

type Circle struct{ Diameter float64 }

func (c *Circle) Fingers() float64 {
	return c.Diameter / 2
}

func (c *Circle) Large() float64 {
	return math.Pi * math.Pow(c.Fingers(), 2)
}

func (c *Circle) Around() float64 {
	return math.Pi * c.Diameter
}
