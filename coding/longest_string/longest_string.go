package main

import "fmt"

func solution(inputArray []string) []string {

	longest := 0
	longestElements := make([]string, 0)

	//Go throught the slice and get each element
	for _, element := range inputArray {
		//Get element lenght
		lenght := len(element)

		//Check if this is the longest
		if lenght > longest {
			longest = lenght
		}
	}

	fmt.Println(longest)

	//Go throught the slice and get each element
	for _, element := range inputArray {
		//Get element if lenght is equal to the longest
		if len(element) == longest {
			//Save it in an array of longests elements
			longestElements = append(longestElements, element)
		}
	}

	fmt.Println(longestElements)
	//return the slice

	return longestElements
}

func main() {
	got := solution([]string{"aba", "aa", "ad", "vcd", "aba"})
	fmt.Println(got)
}
