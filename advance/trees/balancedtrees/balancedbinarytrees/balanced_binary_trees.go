package balancedbinarytrees

import (
	"math"

	"github.com/juanmanuel0963/go_language/v1/advance/datastructures/trees"
)

// IsBalanced checks if a binary tree is balanced
func IsBalanced[T any](node *trees.BinaryNode[T]) bool {
	_, balanced := isBalanced(node)
	return balanced
}

func isBalanced[T any](node *trees.BinaryNode[T]) (int, bool) {
	if node == nil {
		return 0, true
	}

	left, balanced := isBalanced(node.Left)
	if !balanced {
		return 0, false
	}

	right, balanced := isBalanced(node.Right)
	if !balanced {
		return 0, false
	}

	if math.Abs(float64(left-right)) > 1 {
		return 0, false
	}

	if left > right {
		return 1 + left, true
	}
	return 1 + right, true
}
