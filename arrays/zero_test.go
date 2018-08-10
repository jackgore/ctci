package arrays

import (
	"testing"
)

var zeroTests = []struct {
	a      [][]int
	result [][]int
}{
	{
		[][]int{
			{1},
		},

		[][]int{
			{1},
		},
	},
	{
		[][]int{
			{0, 2},
			{3, 4},
		},

		[][]int{
			{0, 0},
			{0, 4},
		},
	},
}

func TestZero(t *testing.T) {
	for _, test := range zeroTests {
		if zero := zero(test.a); !compareMatrix(zero, test.result) {
			t.Errorf("for string %v expected %v but received %v", test.a, test.result, zero)
		}
	}
}
