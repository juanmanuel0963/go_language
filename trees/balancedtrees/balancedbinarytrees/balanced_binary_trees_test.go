package balancedbinarytrees

import (
	"testing"

	"github.com/juanmanuel0963/go_language/v1/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced_EmptyNode(t *testing.T) {
	assert.True(t, IsBalanced[int](nil))
}

func TestIsBalanced_RootNode(t *testing.T) {
	node := &trees.BinaryNode[int]{Data: 1}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_HalfTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node := &trees.BinaryNode[int]{Data: 1, Left: node2}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_FullTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node3 := &trees.BinaryNode[int]{Data: 3}
	node := &trees.BinaryNode[int]{Data: 1, Left: node2, Right: node3}

	assert.True(t, IsBalanced(node))
}

func TestIsBalanced_NotBalancedTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node3 := &trees.BinaryNode[int]{Data: 3, Left: node2}
	node := &trees.BinaryNode[int]{Data: 1, Right: node3}

	assert.False(t, IsBalanced(node))
}
