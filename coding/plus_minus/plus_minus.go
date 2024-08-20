package main

import (
	"fmt"
)

func plusMinus(arr []int) {
	var pos, neg, zero float32

	for _, num := range arr {
		switch {
		case num > 0:
			pos++
		case num < 0:
			neg++
		default:
			zero++
		}
	}

	total := float32(len(arr))
	fmt.Printf("%.6f\n%.6f\n%.6f\n", pos/total, neg/total, zero/total)
}

func main() {

	var T int
	fmt.Scan(&T)

	// Create a slice to store the test cases
	arr := make([]int, T)

	// Read each test case and store it in the slice
	for i := 0; i < T; i++ {
		fmt.Scan(&arr[i])
	}

	plusMinus(arr)
}
