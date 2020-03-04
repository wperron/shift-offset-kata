package offset

import (
	"testing"
)

func TestOffset(t *testing.T) {
	var res int
	params := []struct{
		input []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{2}, 0},
		{[]int{3, 1}, 1},
		{[]int{1, 2, 3}, 0},
		{[]int{5, 7, 9, 1, 3, 4}, 3},
		{[]int{5, 7, 9, 17, 25, 1, 3, 4}, 5},
		{[]int{5, 7, 9, 17, 25, 31, 1, 3, 4}, 6},
		{[]int{5, 7, 9, 17, 25, 31, 43, 1, 3, 4}, 7},
	}

	for _, param := range params {
		res = IntOffset(param.input)
		if res != param.expected {
			t.Errorf("Invalid offset, given %v, expected %d got %d\n", param.input, param.expected, res)
		}
	}
}
