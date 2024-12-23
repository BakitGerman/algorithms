package binarysearch

func binarysearchInt(s []int, target int) int {
	high := len(s) - 1
	low := 0
	for low <= high {
		mid := int(uint(low+high) >> 1)
		guess := s[mid]
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearch(n int, f func(i int) bool) int {
	low, high := 0, n
	for low < high {
		mid := int(uint(low+high) >> 1)
		if !f(mid) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
