package main

import (
	"fmt"
	"strconv"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func compress(chars []byte) int {

	//fmt.Println(chars)

	counter := 1
	compressed := make([]string, 0)
	compressed = append(compressed, string(chars[0]))

	for i := 1; i < len(chars); i++ {

		if chars[i-1] == chars[i] {
			counter += 1
		} else {
			compressed = append(compressed, strconv.Itoa(counter))
			counter = 1

			compressed = append(compressed, string(chars[i]))
		}
	}

	if counter > 1 {
		compressed = append(compressed, strconv.Itoa(counter))
	}

	fmt.Println(compressed)

	return len(compressed)
}

func main() {

	//chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	got := compress(chars)

	fmt.Println(got)
}
