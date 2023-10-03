package main

import "fmt"

func IsPalindrome(inputWord string) bool {

	var theArray []string
	var theIndex int = 0
	var theString string

	for _, theChar := range inputWord {
		theArray = append(theArray, string(theChar))
		theIndex++
	}

	for i := len(theArray) - 1; i >= 0; i-- {
		theString = theString + theArray[i]
	}

	fmt.Println(theString)

	return (inputWord == theString)
}

func solution(inputString string) bool {

	outputString := ""
	//Go reverse throught the inputString
	for i := len(inputString) - 1; i >= 0; i-- {

		//Save the chars in a new string
		outputString += string(inputString[i])
	}

	fmt.Println(outputString)

	return inputString == outputString
}

func main() {

	var isPalindrome = IsPalindrome("JUAN")

	fmt.Println(isPalindrome)
}
