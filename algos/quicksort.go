package main

func qs(data *[]int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(data, left, right)
	qs(data, left, pivot-1)
	qs(data, pivot+1, right)

}

func partition(data *[]int, left, right int) int {
	pivot := (*data)[right]
	idx := left - 1
	for i := left; i < right; i++ {
		if (*data)[i] < pivot {
			idx++
			(*data)[i], (*data)[idx] = (*data)[idx], (*data)[i]
		}
	}
	idx++
	(*data)[right] = (*data)[idx]
	(*data)[idx] = pivot
	return idx
}

func QuickSort(data *[]int) {
	qs(data, 0, len(*data)-1)
}
