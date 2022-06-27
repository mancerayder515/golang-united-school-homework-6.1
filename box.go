package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         []Shape{},
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("box is full")
	}
	
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	
	if len(b.shapes) <= i || b.shapes[i] == nil {
		return nil, errors.New("index out of range or nil")
	}
	
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	
	sh, ok := b.GetByIndex(i)
	if ok != nil {
		return sh, fmt.Errorf("cannot extract nonexistent: %w", ok)
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return sh, ok
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	sh, ok := b.GetByIndex(i)
	if ok != nil {
		return sh, fmt.Errorf("cannot replace nonexistent: %w", ok)
	}
	b.shapes[i] = shape
	return sh, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		if shape == nil {
			continue
		}
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		if shape == nil {
			continue
		}
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	found := false
	for i, shape := range b.shapes {
		if _, ok := shape.(Circle); ok {
			_, _ = b.ExtractByIndex(i)
			found = true
		}
	}
	if !found {
		return errors.New("no circles are found")
	}
	return nil
}
