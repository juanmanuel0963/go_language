package main

import (
	"fmt"
)

// TreeNode struct definition
type TreeNode struct {
	name string
	sons []*TreeNode
}

// Function to create a new TreeNode
func NewTreeNode(name string) *TreeNode {
	return &TreeNode{
		name: name,
		sons: []*TreeNode{},
	}
}

// Function to add a child to a TreeNode
func (node *TreeNode) AddChild(child *TreeNode) {
	node.sons = append(node.sons, child)
}

// Function to recursively create a tree with a given depth
func CreateTree(node *TreeNode, depth int) {
	if depth <= 0 {
		return
	}

	for i := 1; i <= 2; i++ {
		childName := fmt.Sprintf("Node %d at depth %d", i, 4-depth)
		childNode := NewTreeNode(childName)
		node.AddChild(childNode)
		CreateTree(childNode, depth-1)
	}
}

// Function to recursively print the tree
func PrintTree(node *TreeNode, level int) {
	fmt.Printf("%s%s\n", getIndent(level), node.name)
	for _, child := range node.sons {
		PrintTree(child, level+1)
	}
}

// Helper function to get the appropriate indentation
func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	return indent
}

func main() {
	// Create the root node
	root := NewTreeNode("Root")

	// Recursively create a 5-level tree
	CreateTree(root, 3)

	// Print the tree
	PrintTree(root, 0)
}
