package main

import (
	"fmt"
	"strconv"
)

func solution(n int) bool {

	//Convert the integer to string
	inputString := strconv.Itoa(n)
	//fmt.Println(inputString)

	inputSlice := make([]string, 0)

	for i := 0; i < len(inputString); i++ {
		//Save each char in a slice of strings
		inputSlice = append(inputSlice, string(inputString[i]))
	}

	//fmt.Println(inputSlice)

	sumFirstHalf := 0
	sumSecondHalf := 0

	//iterate the slice from 0-middle and sum the values
	for i := 0; i < len(inputSlice)/2; i++ {
		val, _ := strconv.Atoi(inputSlice[i])
		//fmt.Println(val)
		sumFirstHalf += val
	}

	fmt.Printf("Sum first half %v: \n", sumFirstHalf)

	//iterate the slice from middle-end and sum the values
	for i := (len(inputSlice) / 2); i < len(inputSlice); i++ {
		val, _ := strconv.Atoi(inputSlice[i])
		//fmt.Println(val)
		sumSecondHalf += val
	}

	fmt.Printf("Sum second half %v: \n", sumSecondHalf)

	//Compare the 2 results

	return sumFirstHalf == sumSecondHalf
}

func main() {
	got := solution(1230)
	fmt.Println(got)
}
