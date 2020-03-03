package offset

import (
	"math"
)

func IntOffset(list []int) int {
	if len(list) <= 1 {
		return 0
	}

	mid := int(math.Floor(float64(len(list))) / 2) - 1
	low := 0
	high := len(list) - 1
	for low <= mid {
		// check which side of the slice contains the shift
		if list[low] > list[mid] {
			high = mid
			mid = low + int(math.Floor(float64(mid - low)) / 2)
		} else if list[mid + 1] > list[high] {
			low = mid + 1
			mid = int(mid + ((high - mid) / 2))
		} else {
			return mid + 1
		}
	}
	return mid
}
