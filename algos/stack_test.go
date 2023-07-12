package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(5)
	s.Push(7)
	s.Push(9)

	assert.Equal(t, s.Pop(), 9)
	assert.Equal(t, s.size, 2)

	s.Push(11)
	assert.Equal(t, s.Pop(), 11)
	assert.Equal(t, s.Pop(), 7)
	assert.Equal(t, s.Peek(), 5)
	assert.Equal(t, s.Pop(), 5)
	assert.Equal(t, s.Pop(), nil)
	assert.Equal(t, s.size, 0)

}
