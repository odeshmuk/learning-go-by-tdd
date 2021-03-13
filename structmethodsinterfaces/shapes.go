package structmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}
func (circle Circle) Perimeter() (perimeter float64) {
	return 2 * math.Pi * circle.Radius
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.Height * rectangle.Width
}

func (circle Circle) Area() (area float64) {
	return circle.Radius * circle.Radius * math.Pi
}
