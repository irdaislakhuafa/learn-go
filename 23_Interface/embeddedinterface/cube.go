package embeddedinterface

import "math"

type Cube struct{ Side float64 }

func (c *Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (c *Cube) Large() float64 {
	return math.Pow(c.Side, 2)
}

func (c *Cube) Around() float64 {
	return c.Side * 12
}
