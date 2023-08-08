package binarytreeheights

import (
	"testing"

	"github.com/juanmanuel0963/go_language/v1/advance/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestHeight_EmptyNode(t *testing.T) {
	assert.Equal(t, Height[int](nil), 0)
}

func TestHeight_RootNode(t *testing.T) {
	node := &trees.BinaryNode[int]{Data: 1}

	assert.Equal(t, Height(node), 1)
}

func TestHeight_HalfTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node := &trees.BinaryNode[int]{Data: 1, Left: node2}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_FullTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node3 := &trees.BinaryNode[int]{Data: 3}
	node := &trees.BinaryNode[int]{Data: 1, Left: node2, Right: node3}

	assert.Equal(t, Height(node), 2)
}

func TestHeight_LongerTree(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node3 := &trees.BinaryNode[int]{Data: 3, Left: node2}
	node := &trees.BinaryNode[int]{Data: 1, Right: node3}

	assert.Equal(t, Height(node), 3)
}
