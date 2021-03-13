package arrayslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Adding array of a set size", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Testing sum using slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t.Run("Testing sum using slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})
}

func BenchmarkSum(b *testing.B) {
	Sum([]int{1, 2, 3, 4, 5})
}

func BenchmarkArray(b *testing.B) {
	SumArray([5]int{1, 2, 3, 4, 5})
}
