package data_structures_test

import (
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestLinkedList(t *testing.T) {
	// Setup
	list := ds.NewLinkedListEval[int]()

	val := []int{0, -10, 69, 420}
	for i := range val {
		list.Append(val[i])
	}

	// Length
	expectedSize := len(val)
	listSize := list.Length()
	if expectedSize != int(listSize) {
		t.Fatalf("Length: wrong size for the list, expected %d got %d", expectedSize, listSize)
	}

	// Get and Append from setup
	for i := range val {
		res, err := list.Get(uint(i))
		if res != val[i] || err != nil {
			t.Fatalf("Append: got %d expected %d at index %d", res, val[i], i)
		}
	}

	// Prepend
	v := 69420
	list.Prepend(v)
	if res, err := list.Get(0); res != v || err != nil {
		t.Fatalf("Prepend: got %d but expected %d at head", res, v)
	}

	// IndexOf
	idx := list.IndexOf(v)
	if idx != 0 {
		t.Fatalf("IndexOf: got index %d but expected 0 for value %d", idx, v)
	}

	// Set
	v = 123
	err := list.Set(0, v)
	if err != nil {
		t.Fatalf("Set: coult not set %d at index 0, got err %s", v, err)
	}
	if res, err := list.Get(0); res != v || err != nil {
		t.Fatalf("Set: got %d but expected %d at head", res, v)
	}
}

func FuzzLinkedListEval_RemoveValue(f *testing.F) {
	values := []string{"test", "fuzz", "random", "tree", "list"}
	for i := range values {
		f.Add(values[i])
	}

	f.Fuzz(func(t *testing.T, a string) {
		list := ds.NewLinkedListEval[string]()
		list.Append(a)
		res, err := list.Get(0)
		if err != nil || res != a {
			t.Errorf("%q %v", res, err)
		}

		list.RemoveValue(a)
		res, err = list.Get(0)
		if err != nil || res != "" {
			t.Errorf("%q %v", res, err)
		}
	})
}
