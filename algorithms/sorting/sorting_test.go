package sorting

import "testing"

func TestSwap(t *testing.T) {
	a, b := 2, 1

	swap(&a, &b)

	if a != 1 || b != 2 {
		t.Errorf("failed to swap elements, 'a' expected: 1 got: %d, 'b' expected: 2 got: %d", a, b)
	}
}
