package reverses

import "github.com/juanmanuel0963/go_language/v1/datastructures/linkedlists/singlylinkedlists"

// Reverse reverse a singly linked list
func Reverse[T any](node *singlylinkedlists.SLLNode[T]) *singlylinkedlists.SLLNode[T] {
	current := node
	var previous *singlylinkedlists.SLLNode[T]
	var next *singlylinkedlists.SLLNode[T]
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
