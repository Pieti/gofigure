package search

func BinarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack)

	for low < high {
		mid := (low + (high-low)/2)
		v := haystack[mid]
		if v == needle {
			return mid
		} else if v > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1

}
