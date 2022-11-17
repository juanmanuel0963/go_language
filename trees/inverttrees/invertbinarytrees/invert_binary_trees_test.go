package invertbinarytrees

import (
	"testing"

	"github.com/juanmanuel0963/go_language/v1/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestInvertTrees(t *testing.T) {
	node2 := &trees.BinaryNode[int]{Data: 2}
	node3 := &trees.BinaryNode[int]{Data: 3}
	node := &trees.BinaryNode[int]{Data: 1, Left: node2, Right: node3}

	node = InvertTree(node)

	assert.Equal(t, node.Left.Data, 3)
	assert.Equal(t, node.Right.Data, 2)
}
