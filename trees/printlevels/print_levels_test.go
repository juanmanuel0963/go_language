package printlevels

import (
	"testing"

	"github.com/juanmanuel0963/go_language/v1/datastructures/trees"
)

func TestPrintLevels(t *testing.T) {
	node4 := &trees.MultiNode[int]{Data: 5, Children: make([]*trees.MultiNode[int], 0)}
	node5 := &trees.MultiNode[int]{Data: 6, Children: make([]*trees.MultiNode[int], 0)}

	node6 := &trees.MultiNode[int]{Data: 7, Children: make([]*trees.MultiNode[int], 0)}
	node7 := &trees.MultiNode[int]{Data: 8, Children: make([]*trees.MultiNode[int], 0)}

	node1 := &trees.MultiNode[int]{Data: 2, Children: []*trees.MultiNode[int]{node4, node5}}
	node2 := &trees.MultiNode[int]{Data: 3, Children: []*trees.MultiNode[int]{node6, node7}}
	node3 := &trees.MultiNode[int]{Data: 4, Children: make([]*trees.MultiNode[int], 0)}

	node := &trees.MultiNode[int]{Data: 1, Children: []*trees.MultiNode[int]{node1, node2, node3}}
	PrintLevels(node)
}
