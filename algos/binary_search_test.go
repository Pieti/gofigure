package main

import (
	"testing"
)

type binarySearchTestCase struct {
	target   int
	expected int
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	tests := []binarySearchTestCase{
		{target: 69, expected: 3},
		{target: 1336, expected: -1},
		{target: 1, expected: 0},
		{target: 0, expected: -1},
		{target: 69420, expected: 10},
	}

	for _, tt := range tests {
		actual := BinarySearch(arr, tt.target)
		if actual != tt.expected {
			t.Errorf("BinarySearch(%d): expected %v, actual %v", tt.target, tt.expected, actual)
		}
	}
}

func TestBinarySearchEmptyArray(t *testing.T) {
	arr := []int{}

	actual := BinarySearch(arr, 10)
	if actual != -1 {
		t.Errorf("BinarySearch(%d): expected %v, actual %v", 10, -1, actual)
	}
}
