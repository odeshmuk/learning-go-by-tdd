package structmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Testing Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkPerimeter(t, rectangle, 36)
	})

	t.Run("Testing Circle Perimeter", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)

	})

}

func BenchmarkPerimeter(b *testing.B) {

	rectangle := Rectangle{12.0, 6.0}
	rectangle.Perimeter()

	circle := Circle{10.0}
	circle.Perimeter()

	benchmarksTable := []struct {
		shape Shape
	}{
		{Rectangle{12.0, 6.0}},
		{Circle{10.0}},
	}
	for _, shapeInstance := range benchmarksTable {
		shapeInstance.shape.Perimeter()
		shapeInstance.shape.Area()
	}

}
