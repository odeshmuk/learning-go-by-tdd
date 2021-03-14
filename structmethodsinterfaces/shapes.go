package structmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

//interface resolution is implicit
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}
func (circle Circle) Perimeter() (perimeter float64) {
	return 2 * math.Pi * circle.Radius
}

func (triangle Triangle) Perimeter() (perimeter float64) {
	return triangle.Base + triangle.Height + math.Sqrt(math.Pow(triangle.Base, 2)+math.Pow(triangle.Height, 2))
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Area() (area float64) {
	return circle.Radius * circle.Radius * math.Pi
}

func (triangle Triangle) Area() (perimeter float64) {
	return .5 * triangle.Base * triangle.Height
}
