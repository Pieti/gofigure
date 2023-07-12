package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 1, 1, 0, 2, 8, 5, 1, 6}
	QuickSort(&arr)
	assert.Equal(t, []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr)
}
