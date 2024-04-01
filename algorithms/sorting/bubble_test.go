package sorting_test

import (
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func TestBubbleSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		tests := []sorting.Cases[int]{
			sorting.SortedCases([]int{1, 6, 5, 3, 9, 10}),
		}

		sorting.HelperSortOrdered[[]int](t, sorting.BubbleSort, tests)
	})

	t.Run("strings", func(t *testing.T) {
		tests := []sorting.Cases[string]{
			sorting.SortedCases([]string{"z", "x", "c", "v", "b", "n"}),
			sorting.SortedCases([]string{"zebra", "abobora", "urso", "abelha", "limao", "batata"}),
		}

		sorting.HelperSortOrdered[[]string](t, sorting.BubbleSort, tests)
	})

	t.Run("floats", func(t *testing.T) {
		tests := []sorting.Cases[float32]{
			sorting.SortedCases([]float32{1.123, 1.321, 2.0, 0.3}),
		}

		sorting.HelperSortOrdered[[]float32](t, sorting.BubbleSort, tests)
	})
}
