package main

import (
	"fmt"
	"strings"
)

func solution(inputString string) string {

	if inputString == "()" {
		return ""
	}

	outputString := inputString

	counterNestedParentheses := findNestedParentheses(inputString)

	fmt.Println(counterNestedParentheses)

	for i := 0; i <= counterNestedParentheses; i++ {

		outputString = replaceParentheses(outputString)
		//fmt.Println(outputString)

	}

	return outputString
}

func replaceParentheses(inputString string) string {

	outputString := ""

	for i := len(inputString) - 1; i >= 0; i-- {

		//"Capturing )"
		if string(inputString[i]) == ")" && string(inputString[i-1]) != ")" && string(inputString[i-1]) != "(" {

			outputString = outputString + string(inputString[i])
		}

		if strings.Contains(outputString, ")") {
			//"Capturing characters"
			if string(inputString[i]) != "(" && string(inputString[i]) != ")" {

				outputString = outputString + string(inputString[i])
			}
		}

		//"Capturing ("
		if string(inputString[i]) == "(" && string(inputString[i+1]) != "(" && string(inputString[i+1]) != ")" {

			outputString = outputString + string(inputString[i])
			break
		}
	}

	oldString := reverseString(outputString)

	newString := reverseStringWithParentheses(oldString)

	outputString = strings.Replace(inputString, oldString, newString, 1)

	return outputString
}

func reverseString(inputString string) string {

	outputString := ""

	for i := len(inputString) - 1; i >= 0; i-- {

		//if string(inputString[i]) != "(" && string(inputString[i]) != ")" {
		outputString = outputString + string(inputString[i])
		//}

	}

	return outputString
}

func reverseStringWithParentheses(inputString string) string {

	outputString := ""

	for i := len(inputString) - 1; i >= 0; i-- {

		if string(inputString[i]) != "(" && string(inputString[i]) != ")" {
			outputString = outputString + string(inputString[i])
		}

	}

	return outputString
}

func findNestedParentheses(inputString string) int {

	counter := 0

	for i := len(inputString) - 1; i >= 0; i-- {

		if string(inputString[i]) == ")" {
			counter += 1
		}
	}

	return counter

}

func main() {
	got := solution("()") //mybar
	//got := solution("my(bar)")//mybar
	//got := solution("foo(bar)baz")//foorabbaz
	//got := solution("foo(bar)baz(blim)") //foorabbazmilb
	//got := solution("foo(bar(baz))blim") //foobazrabblim
	//got := solution("foo(bar(baz))blim") //foobazrabblim

	fmt.Println(got)
}
