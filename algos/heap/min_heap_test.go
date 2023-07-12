package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()

	h.Push(5)
	h.Push(3)
	h.Push(69)
	h.Push(420)
	h.Push(4)
	h.Push(1)
	h.Push(8)
	h.Push(7)

	assert.Equal(t, 8, h.length)
	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 5, h.Pop())
	assert.Equal(t, 4, h.length)
	assert.Equal(t, 7, h.Pop())
	assert.Equal(t, 8, h.Pop())
	assert.Equal(t, 69, h.Pop())
	assert.Equal(t, 420, h.Pop())
	assert.Equal(t, 0, h.length)

}
