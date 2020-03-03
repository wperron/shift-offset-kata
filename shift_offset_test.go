package offset

import (
	"testing"
)

func TestOffset(t *testing.T) {
	var res int

	res = IntOffset([]int{})
	if res != 0 {
		t.Errorf("Invalid offset, expected 0 got %d\n", res)
	}

	res = IntOffset([]int{2})
	if res != 0 {
		t.Errorf("Invalid offset, expected 0 got %d\n", res)
	}

	res = IntOffset([]int{3, 1})
	if res != 1 {
		t.Errorf("Invalid offset, expected 1 got %d\n", res)
	}

	res = IntOffset([]int{5, 7, 9, 1, 3, 4})
	if res != 3 {
		t.Errorf("Invalid offset, expected 3 got %d\n", res)
	}

	res = IntOffset([]int{5, 7, 9, 17, 25, 1, 3, 4})
	if res != 5 {
		t.Errorf("Invalid offset, expected 5 got %d\n", res)
	}

	res = IntOffset([]int{5, 7, 9, 17, 25, 31, 1, 3, 4})
	if res != 6 {
		t.Errorf("Invalid offset, expected 6 got %d\n", res)
	}

	res = IntOffset([]int{5, 7, 9, 17, 25, 31, 43, 1, 3, 4})
	if res != 7 {
		t.Errorf("Invalid offset, expected 7 got %d\n", res)
	}
}
