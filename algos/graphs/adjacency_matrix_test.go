package graphs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdjacencyMatrix(t *testing.T) {
	assert.Equal(t, BreadthFirstSearch(matrix, 0, 6), []int{0, 1, 4, 5, 6})
}
