package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tree := NewBinaryTree()
	assert.Equal(t, true, BreadthFirstSearch(tree, 45))
	assert.Equal(t, true, BreadthFirstSearch(tree, 7))
	assert.Equal(t, false, BreadthFirstSearch(tree, 69))
}
