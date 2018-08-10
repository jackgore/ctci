package arrays

import (
	"testing"
)

var rotateTests = []struct {
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
			{1, 2},
			{3, 4},
		},

		[][]int{
			{2, 4},
			{1, 3},
		},
	},
}

func compareMatrix(a [][]int, b [][]int) bool {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestRotate(t *testing.T) {
	for _, test := range rotateTests {
		/*		if rotate := rotate(test.a); !compareMatrix(rotate, test.result) {

				t.Errorf("for string %v expected %v but received %v", test.a, test.result, rotate)
			}*/

		if rotate := rotateSwap(test.a); !compareMatrix(rotate, test.result) {
			t.Errorf("for string %v expected %v but received %v", test.a, test.result, rotate)
		}
	}
}
