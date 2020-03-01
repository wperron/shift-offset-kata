package offset

import (
	"math"
)

func IntOffset(list []int) int {
	if len(list) == 2 {
		if list[0] > list [1] {
			return 1
		}
	}

	half := int(math.Floor(float64(len(list))) / 2)
	left := list[0:half]
	right := list[half:]

	// case where shift is exactly half of length
	if left[len(left) - 1] > right[0] {
		return half
	}

	// check which side of the slice contains the shift
	if left[0] > left[len(left) - 1] {
		return IntOffset(left)
	} else {
		return len(left) + IntOffset(right)
	}
}