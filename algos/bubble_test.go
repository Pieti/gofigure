package main

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	BubbleSort(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("BubbleSort() failed, %v", arr)
		}
	}
}
