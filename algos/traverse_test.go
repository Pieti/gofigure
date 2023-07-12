package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreOrderSearch(t *testing.T) {
	tree := NewBinaryTree()
	path := PreOrderSearch(tree)
	assert.Equal(t, []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}, *path)
}

func TestInOrderSearch(t *testing.T) {
	tree := NewBinaryTree()
	path := InOrderSearch(tree)
	assert.Equal(t, []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}, *path)
}

func TestPostOrderSearch(t *testing.T) {
	tree := NewBinaryTree()
	path := PostOrderSearch(tree)
	assert.Equal(t, []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}, *path)
}
