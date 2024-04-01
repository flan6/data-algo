package sorting_test

import (
	"fmt"
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func TestSelectionSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		tests := []sorting.Cases[int]{
			sorting.SortedCases([]int{1, 6, 5, 3, 9, 10}),
		}

		sorting.HelperSortOrdered[[]int](t, sorting.SelectionSort, tests)
	})

	t.Run("strings", func(t *testing.T) {
		tests := []sorting.Cases[string]{
			sorting.SortedCases([]string{"z", "x", "c", "v", "b", "n"}),
			sorting.SortedCases([]string{"zebra", "abobora", "urso", "abelha", "limao", "batata"}),
		}

		sorting.HelperSortOrdered[[]string](t, sorting.SelectionSort, tests)
	})

	t.Run("floats", func(t *testing.T) {
		tests := []sorting.Cases[float32]{
			sorting.SortedCases([]float32{1.123, 1.321, 2.0, 0.3}),
		}

		sorting.HelperSortOrdered[[]float32](t, sorting.SelectionSort, tests)
	})
}

func ExampleSelectionSort() {
	strs := []string{"abelha", "urso", "abobora", "limao", "zebra"}
	ints := []int{9, 8, 6, 5, 3, 7}
	floats := []float32{1.123, 1.321, 2.0, 0.3}

	sorting.SelectionSort(strs)
	fmt.Println(strs)

	sorting.SelectionSort(ints)
	fmt.Println(ints)

	sorting.SelectionSort(floats)
	fmt.Println(floats)

	// Output:
	// [abelha abobora limao urso zebra]
	// [3 5 6 7 8 9]
	// [0.3 1.123 1.321 2]
}
