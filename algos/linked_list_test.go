package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := &LinkedList{}
	testList(l, t)

}

func testList(l List, t *testing.T) {
	l.Append(5)
	l.Append(7)
	l.Append(9)

	res, ok := l.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 9, res)

	res, ok = l.RemoveAt(1)
	assert.True(t, ok)
	assert.Equal(t, 7, res)

	length := l.Length()
	assert.Equal(t, 2, length)

	l.Append(11)
	res, ok = l.RemoveAt(1)
	assert.Equal(t, 9, res)

	res, ok = l.Remove(9)
	assert.False(t, ok)
	assert.Equal(t, nil, res)

	res, ok = l.RemoveAt(0)
	assert.True(t, ok)
	assert.Equal(t, 5, res)

	res, ok = l.RemoveAt(0)
	assert.True(t, ok)
	assert.Equal(t, 11, res)

	length = l.Length()
	assert.Equal(t, 0, length)

	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)

	res, ok = l.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 5, res)

	res, ok = l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 9, res)

	res, ok = l.Remove(9)
	assert.True(t, ok)
	assert.Equal(t, 9, res)

	length = l.Length()
	assert.Equal(t, 2, length)

	res, ok = l.Get(0)
	assert.True(t, ok)
	assert.Equal(t, 7, res)
}
