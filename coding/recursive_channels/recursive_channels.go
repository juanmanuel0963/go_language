package main

import (
	"fmt"
)

// TreeNode struct definition
type TreeNode struct {
	name string
	sons []*TreeNode
}

func main() {
	// Create the root node
	root := NewTreeNode("Root")

	// Create a channel to signal the completion of the tree creation
	chParentDone := make(chan struct{})

	// Recursively create a 3-level tree using channels
	go func() {
		CreateTree(root, 3, chParentDone)
		close(chParentDone)
	}()

	// Wait for the tree creation to finish
	<-chParentDone

	// Print the tree
	PrintTree(root, 0)
}

// Function to recursively create a tree with a given depth using channels
func CreateTree(node *TreeNode, depth int, chParentDone chan struct{}) {
	if depth <= 0 {
		chParentDone <- struct{}{}
		return
	}

	chSonDone := make(chan struct{}, 2)

	for i := 1; i <= 2; i++ {
		go func(i int) {
			defer func() { chSonDone <- struct{}{} }()
			sonName := fmt.Sprintf("Node %d at depth %d", i, 4-depth)
			sonNode := NewTreeNode(sonName)
			node.AddSon(sonNode)
			CreateTree(sonNode, depth-1, chSonDone)
		}(i)
	}

	// Wait for all child goroutines to finish
	for i := 0; i < 2; i++ {
		<-chSonDone
	}

	chParentDone <- struct{}{}
}

// Function to create a new TreeNode
func NewTreeNode(name string) *TreeNode {
	return &TreeNode{
		name: name,
		sons: []*TreeNode{},
	}
}

// Function to add a child to a TreeNode
func (node *TreeNode) AddSon(child *TreeNode) {
	node.sons = append(node.sons, child)
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
