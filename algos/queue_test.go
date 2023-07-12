package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)

	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, q.size, 2)

	q.Enqueue(11)
	assert.Equal(t, q.Dequeue(), 7)
	assert.Equal(t, q.Dequeue(), 9)
	assert.Equal(t, q.Peek(), 11)
	assert.Equal(t, q.Dequeue(), 11)
	assert.Equal(t, q.Dequeue(), nil)
	assert.Equal(t, q.size, 0)

}
