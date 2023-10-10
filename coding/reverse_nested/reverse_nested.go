package main

import "fmt"

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

//"foo(bar(baz(aeiou)))blim"
func findBlocksRecursive(inputString string, words []string, nodeParent *TreeNode) (string, []string) {

	//Iterate on the string and find "("" "content" ")

	subContent := ""

	for i := 0; i < len(inputString); i++ {
		//Find left "("
		if string(inputString[i]) == "(" || string(inputString[i]) == ")" {

			newWord := inputString[:i]

			if len(newWord) > 0 {

				words = append(words, newWord)

				//Create a new Son
				nodeSon := NewTreeNode(newWord)

				//Append the son to the parent node
				nodeParent.AddChild(nodeSon)

				//Find "subContent"
				subContent = inputString[i+1:]

				if len(subContent) > 0 {
					subContent, words = findBlocksRecursive(subContent, words, nodeSon)
				}
			}

			break
		}
	}
	/*
		if len(subContent) > 0 {
			subContent, words = findBlocksRecursive(subContent, words, &nodeSon)
		} else {
	*/

	if len(inputString) > 0 {
		fmt.Println("input: '" + inputString + "'")
		fmt.Println("nodeParent: '" + nodeParent.name + "'")
		fmt.Println("-->", inputString)

		/*
			//Create a new Son
			nodeSon := NewTreeNode(inputString)

			//Append the son to the parent node
			nodeParent.AddChild(nodeSon)
		*/
	}

	return subContent, words
}

func printNode(nodeParent *TreeNode, level int) {

	fmt.Printf("Level %v, Parent %s, Sons %v \n", level, nodeParent.name, len(nodeParent.sons))

	for i := 0; i < len(nodeParent.sons); i++ {
		level++
		printNode(nodeParent.sons[i], level)
	}
}

func solution(inputString string) string {

	words := make([]string, 0)
	// Create the root node
	root := NewTreeNode("Root")

	newString, words := findBlocksRecursive(inputString, words, root)

	fmt.Println("----")
	fmt.Println(newString)
	fmt.Println(words)
	printNode(root, 0)

	return newString
}

func main() {
	got := solution("foo(bar(baz(aeiou)))blim")
	fmt.Println(got)
}
