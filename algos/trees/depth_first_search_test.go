package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepthFirstSear(t *testing.T) {
	tree := NewBinaryTree()
	assert.Equal(t, DepthFirstSearch(tree, 45), true)
	assert.Equal(t, DepthFirstSearch(tree, 7), true)
	assert.Equal(t, DepthFirstSearch(tree, 69), false)
}
