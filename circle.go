package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	return c.Radius * math.Pi * 2
}

func (c Circle) CalcArea() float64 {
	return c.Radius * c.Radius * math.Pi
}
