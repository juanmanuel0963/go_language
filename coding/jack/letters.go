package main

import (
	"fmt"
	"strconv"
)

func main() {

	result1 := encode("aaaabbbccazzzz")
	fmt.Println(result1)

	result2 := decode("a4b3c2a1z4")
	fmt.Println(result2)
}

// a4b3c2a1 -> aaaabbbcca
func decode(input string) string {

	//iterate input, every 2 jumps count
	//append string n times the letter found

	var result string

	for i := 0; i < len(input); i = i + 2 {

		letter := string(input[i])
		count, _ := strconv.Atoi(string(input[i+1]))

		//fmt.Println(letter)
		//fmt.Println(count)

		for j := 0; j < count; j++ {
			result = result + letter
		}
	}

	return result
}

// aaaabbbccazzzz -> a4b3c2a1z4

func encode(input string) string {

	letters := make([]string, 0)
	count := make([]int, 0)
	counter := 0

	// fmt.Println(len(input))

	letters = append(letters, string(input[0]))
	counter++

	//Iterate input string
	for i := 1; i < len(input); i++ {

		//fmt.Println(string(input[i]))
		if string(input[i]) != letters[len(letters)-1] {
			count = append(count, counter)
			letters = append(letters, string(input[i]))
			counter = 1
		} else {
			counter++
		}

	}

	count = append(count, counter)
	//iterate slice to return expected string
	var result string

	for i := 0; i < len(letters); i++ {
		result = result + string(letters[i]) + strconv.Itoa(count[i])
	}
	//fmt.Println(result)
	return result
}
