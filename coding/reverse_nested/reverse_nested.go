package main

import "fmt"

type TreeNode struct {
	name string
	sons []*TreeNode
}

func main() {
	//solution("(bar)")
	//solution("foo(bar)")
	//solution("foo(bar)baz")
	//solution("foo(bar)baz(blim)")
	//solution("foo(bar(baz(aeiou)))blim")
	solution("foo(bar(baz))blim")
}

func solution(inputString string) {

	// Create the root node
	root := NewTreeNode("Root")

	RecursiveParser(inputString, root)

	PrintTree(root, 0)

}

//"foo(bar(baz(aeiou)))blim"
func RecursiveParser(inputString string, parentNode *TreeNode) {

	newNodeName := ""

	for i := 0; i < len(inputString); i++ {

		if string(inputString[i]) != "(" && string(inputString[i]) != ")" {

			newNodeName = newNodeName + string(inputString[i])

			if i == len(inputString)-1 {

				//If a new node is found
				if len(newNodeName) > 0 {

					//Create a new Son
					sonNode := NewTreeNode(string(newNodeName))

					//Append the son to the parent node
					parentNode.AddChild(sonNode)

					newNodeName = ""
				}

			}

		} else if string(inputString[i]) == "(" {

			//If a new node is found
			if len(newNodeName) > 0 {

				//Create a new Son
				sonNode := NewTreeNode(string(newNodeName))

				//Append the son to the parent node
				parentNode.AddChild(sonNode)

				newSubString := FindSubstring(inputString[i+1:], ")")

				if len(newSubString) == 0 {
					newSubString = inputString[i+1:]
				}

				RecursiveParser(newSubString, sonNode)

				i = i + len(newSubString)

				newNodeName = ""
			}

		} else if string(inputString[i]) == ")" {

			//If a new node is found
			if len(newNodeName) > 0 {

				//Create a new Son
				sonNode := NewTreeNode(string(newNodeName))

				//Append the son to the parent node
				parentNode.AddChild(sonNode)

				newNodeName = ""

			}
		}
	}
}

func FindSubstring(inputString string, charToFind string) string {

	returnString := ""

	for i := 0; i < len(inputString); i++ {

		if string(inputString[i]) == string(charToFind) {

			returnString = inputString[:i]
			break
		}
	}

	return returnString
}

/*
func ReverseParser(inputString string, parentNode *TreeNode) int {

	newNodeName := ""

	index := 0

	for i := len(inputString) - 1; i >= 0; i-- {

		if string(inputString[i]) != "(" && string(inputString[i]) != ")" {

			newNodeName = newNodeName + string(inputString[i])

		} else if string(inputString[i]) == ")" {

			//If a new node is found
			if len(newNodeName) > 0 {

				//Create a new Son
				sonNode := NewTreeNode(ReverseString(string(newNodeName))) //"foo(bar(baz(aeiou)))blim"

				//Append the son to the parent node
				parentNode.AddChild(sonNode)

				newNodeName = ""

			}

			index = i
			break
		}

	}

	return index
}
func ReverseString(inputString string) string {

	outputString := ""

	for i := len(inputString) - 1; i >= 0; i-- {
		outputString = outputString + string(inputString[i])
	}

	return outputString
}
*/
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

/*
func printNode(nodeParent *TreeNode, level int) {

	fmt.Printf("Level %v, Parent %s, Sons %v \n", level, nodeParent.name, len(nodeParent.sons))

	for i := 0; i < len(nodeParent.sons); i++ {
		level++
		printNode(nodeParent.sons[i], level)
	}
}
*/

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
		indent += "-"
	}
	return indent
}
