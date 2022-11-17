package printzigzag

import (
	"fmt"

	"github.com/juanmanuel0963/go_language/v1/datastructures/linkedlists/doublylinkedlists"
	"github.com/juanmanuel0963/go_language/v1/datastructures/queues"
	"github.com/juanmanuel0963/go_language/v1/datastructures/trees"
)

// PrintZigZag prints a tree zig zag
func PrintZigZag[T any](node *trees.MultiNode[T]) {
	if node == nil {
		return
	}

	queue := queues.New[*trees.MultiNode[T]]()
	queue.Enqueue(node)

	list := doublylinkedlists.New[*trees.MultiNode[T]]()

	isForwardDirection := true
	for !queue.IsEmpty() {
		for size := queue.Size(); size > 0; size-- {
			element, _ := queue.Dequeue()
			fmt.Print(element.Data, " ")

			for _, child := range element.Children {
				list.Add(child)
			}
		}

		if isForwardDirection {
			for _, value := range list.GetValues() {
				queue.Enqueue(value)
			}
		} else {
			for _, value := range list.GetReverseValues() {
				queue.Enqueue(value)
			}
		}

		list.Clear()
		isForwardDirection = !isForwardDirection
		fmt.Println()
	}
}
