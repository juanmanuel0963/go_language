package main

import "fmt"

func findBlocksRecursive(inputString string, words []string) (string, []string) {

	fmt.Println("input: '" + inputString + "'")
	//Iterate on the string and find "("" "content" ")

	content := ""

	for i := 0; i < len(inputString); i++ {
		//Find left "("
		if string(inputString[i]) == "(" || string(inputString[i]) == ")" {

			//Find "content"
			content = inputString[i+1:]
			fmt.Println("content: '" + content + "'")

			newWord := inputString[:i]

			if len(newWord) > 0 {
				words = append(words, "*"+newWord+"*")
			}

			break
		}
	}

	if len(content) > 0 {
		content, words = findBlocksRecursive(content, words)
	} else {
		if len(inputString) > 0 {
			words = append(words, "'"+inputString+"'")
		}
	}

	return content, words
}

func solution(inputString string) string {

	words := make([]string, 0)
	newString, words := findBlocksRecursive(inputString, words)

	fmt.Println(newString)
	fmt.Println(words)

	return newString
}

func main() {
	got := solution("foo(bar(baz(aeiou)))blim")
	fmt.Println(got)
}
