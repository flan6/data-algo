package sorting_test

import (
	"math/rand"
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func TestMergeSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		ints := make([]int, 21)
		for i := range ints {
			ints[i] = rand.Intn(99)
		}

		tests := []sorting.Cases[int]{
			sorting.SortedCases(ints),
			sorting.SortedCases([]int{9, 0, 10, 99, 1, 88, 6, 5, 3, int(1 << 62)}),
			sorting.SortedCases([]int{1, 2, 0, 22, 6, 5, 3, 9, 10}),
			sorting.SortedCases([]int{111, 11, 1, 2, 1, 1111, 1}),
		}

		sorting.HelperSortOrdered[[]int](t, sorting.MergeSort, tests)
	})

	t.Run("strings", func(t *testing.T) {
		tests := []sorting.Cases[string]{
			sorting.SortedCases([]string{"z", "x", "c", "v", "b", "n"}),
			sorting.SortedCases([]string{"zebra", "abobora", "urso", "abelha", "limao", "batata"}),
		}

		sorting.HelperSortOrdered[[]string](t, sorting.MergeSort, tests)
	})

	t.Run("floats", func(t *testing.T) {
		tests := []sorting.Cases[float32]{
			sorting.SortedCases([]float32{1.123, 1.321, 2.0, 0.3, 0}),
		}

		sorting.HelperSortOrdered[[]float32](t, sorting.MergeSort, tests)
	})
}
