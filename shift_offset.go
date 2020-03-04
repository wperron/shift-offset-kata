package offset

import (
	"math"
)

func IntOffset(list []int) int {
	if len(list) <= 1 {
		return 0
	}

	if list[0] < list[len(list) - 1] {
		return 0
	}

	mid := int(len(list) / 2)
	low := 0
	high := len(list) - 1
	for low <= mid {
		// check which side of the slice contains the shift
		if list[low] > list[mid - 1] {
			high = mid
			mid = low + int(math.Floor(float64(mid - low)) / 2)
		} else if list[mid] > list[high] {
			low = mid
			mid = int(mid + ((high - mid) / 2))
		} else {
			return mid
		}
	}
	return mid
}
