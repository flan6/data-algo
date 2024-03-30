package data_structures_test

import (
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestStack(t *testing.T) {
	stack := ds.NewStack[string]()

	vals := []string{"first", "second", "last"}
	for _, val := range vals {
		stack.Push(val)
	}

	for i := len(vals) - 1; i >= 0; i-- {
		res := stack.Pop()
		val := vals[i]

		if res != val {
			t.Errorf("assert failed, expected: %v got: %v", val, res)
		}
	}
}
