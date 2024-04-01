package sorting_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

type cases[T cmp.Ordered] struct {
	input    []T
	expected []T
}

func TestSelectionSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		tests := []cases[int]{
			sortedCases([]int{1, 6, 5, 3, 9, 10}),
		}

		for name, test := range tests {
			t.Run(fmt.Sprintf("%s case: %d", t.Name(), name), func(t *testing.T) {
				t.Parallel()
				sorting.SelectionSort(test.input)

				for i := range test.input {
					if test.input[i] != test.expected[i] {
						t.Fatalf("input not sorted, got: %v expected: %v", test.input, test.expected)
					}
				}
			})
		}
	})

	t.Run("strings", func(t *testing.T) {
		tests := []cases[string]{
			sortedCases([]string{"z", "x", "c", "v", "b", "n"}),
			sortedCases([]string{"zebra", "abobora", "urso", "abelha", "limao", "batata"}),
		}

		for name, test := range tests {
			t.Run(fmt.Sprintf("%s case: %d", t.Name(), name), func(t *testing.T) {
				t.Parallel()
				sorting.SelectionSort(test.input)

				for i := range test.input {
					if test.input[i] != test.expected[i] {
						t.Fatalf("input not sorted, got: %v expected: %v", test.input, test.expected)
					}
				}
			})
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []cases[float32]{
			sortedCases([]float32{1.123, 1.321, 2.0, 0.3}),
		}

		for name, test := range tests {
			t.Run(fmt.Sprintf("%s case: %d", t.Name(), name), func(t *testing.T) {
				t.Parallel()
				sorting.SelectionSort(test.input)

				for i := range test.input {
					if test.input[i] != test.expected[i] {
						t.Fatalf("input not sorted, got: %v expected: %v", test.input, test.expected)
					}
				}
			})
		}
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

func sortedCases[E []T, T cmp.Ordered](input E) cases[T] {
	expec := make(E, len(input))
	copy(expec, input)
	slices.Sort(expec)
	return cases[T]{input: input, expected: expec}
}
