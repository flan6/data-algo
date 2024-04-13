package algorithms_test

import (
	"testing"

	"github.com/flan6/data-algo/algorithms"
)

type testCase struct {
	genInput      func() []bool
	expectedIndex int
}

func TestTwoCrystalBalls(t *testing.T) {
	tests := map[string]testCase{
		"big": {genInput: func() []bool {
			input := make([]bool, 10000)

			for i := 5000; i < 10000; i++ {
				input[i] = true
			}

			return input
		}, expectedIndex: 5000},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := algorithms.TwoCrystalBalls(test.genInput())
			if res != test.expectedIndex {
				t.Errorf("unexpected result, got: %d expected: %d", res, test.expectedIndex)
			}
		})
	}
}
