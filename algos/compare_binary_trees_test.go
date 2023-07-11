package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareBinaryTrees(t *testing.T) {
	t1 := NewBinaryTree()
	t2 := NewBinaryTreeTwo()

	assert.Equal(t, CompareBinaryTrees(t1, t1), true)
	assert.Equal(t, CompareBinaryTrees(t1, t2), false)

}
