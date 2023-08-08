package printlevels

import (
	"fmt"

	"github.com/juanmanuel0963/go_language/v1/advance/datastructures/queues"
	"github.com/juanmanuel0963/go_language/v1/advance/datastructures/trees"
)

// PrintLevels prints a tree level by level
func PrintLevels[T any](node *trees.MultiNode[T]) {
	if node == nil {
		return
	}

	queue := queues.New[*trees.MultiNode[T]]()
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		for size := queue.Size(); size > 0; size-- {
			element, _ := queue.Dequeue()
			fmt.Print(element.Data, " ")

			for _, child := range element.Children {
				queue.Enqueue(child)
			}
		}
		fmt.Println()
	}
}
