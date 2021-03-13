package structmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Testing Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {

		}
	})

	t.Run("Testing Circle Perimeter", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

}

func BenchmarkPerimeter(b *testing.B) {

	rectangle := Rectangle{12.0, 6.0}
	rectangle.Perimeter()

	circle := Circle{10.0}
	circle.Perimeter()
}

func BenchmarkArea(b *testing.B) {

	rectangle := Rectangle{12.0, 6.0}
	rectangle.Area()

	circle := Circle{10.0}
	circle.Area()
}
