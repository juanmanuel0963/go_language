package detectcycles

import "github.com/juanmanuel0963/go_language/v1/advance/datastructures/linkedlists/singlylinkedlists"

// ContainsCycle checks if the list contains a cycle
func ContainsCycle[T comparable](head *singlylinkedlists.SLLNode[T]) bool {
	fastPointer := head
	slowPointer := head
	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
