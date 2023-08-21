package main

import (
	"fmt"
	"strconv"
)

func sumUntilEqual(S1, S2 int) int {
	sumS1 := S1
	sumS2 := S2

	for {

		newSumS1 := sumS1
		newSumS2 := sumS2

		for _, digitChar := range strconv.Itoa(newSumS1) {
			digit, _ := strconv.Atoi(string(digitChar))
			//fmt.Println(digit)
			newSumS1 += digit
		}

		fmt.Println(newSumS1)

		for _, digitChar := range strconv.Itoa(newSumS2) {
			digit, _ := strconv.Atoi(string(digitChar))
			//fmt.Println(digit)
			newSumS2 += digit
		}

		fmt.Println(newSumS2)

		fmt.Println("")

		sumS1 = newSumS1
		sumS2 = newSumS2

		if newSumS1 == newSumS2 {
			break
		}

		if newSumS2 > 1000 {
			break
		}
	}

	return sumS2
}
