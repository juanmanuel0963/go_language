package main

import (
	"fmt"
	"strings"
)

func solution(inputString string) string {

	stack := make([]int, 0)
	buffer := []rune(inputString)
	//fmt.Printf("buffer:%v\n", string(buffer))

	for i, char := range inputString {

		if char == '(' {

			stack = append(stack, i)
			//fmt.Printf("stack with '(': %v\n", stack)

		} else if char == ')' {

			j := stack[len(stack)-1]
			//fmt.Printf("j:%v\n", j)

			stack = stack[:len(stack)-1]
			//fmt.Printf("stack with ')': %v\n", stack)

			reverse(buffer[j+1 : i])
			//fmt.Printf("buffer:%v\n", string(buffer))
		}
	}

	var result strings.Builder

	for _, char := range buffer {
		if char != '(' && char != ')' {

			result.WriteRune(char)
			//fmt.Println("result:\n", result.String())
		}
	}

	return result.String()
}

func reverse(slice []rune) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	inputString1 := "(bar)"
	fmt.Println(solution(inputString1)) // Output: "rab"
	/*
		inputString2 := "foo(bar)baz"
		fmt.Println(solution(inputString2)) // Output: "foorabbaz"
	*/
}
