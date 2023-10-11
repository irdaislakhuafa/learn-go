package examinterface

import "math"

type Rectangle struct{ Side float64 }

func (r *Rectangle) Large() float64 {
	return math.Pow(r.Side, 2)
}

func (r *Rectangle) Around() float64 {
	return r.Side * 4
}
