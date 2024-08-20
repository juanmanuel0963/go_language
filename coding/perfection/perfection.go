package main

import (
	"fmt"
)

func perfection(arr []int) {

	for iArray := 0; iArray < len(arr); iArray++ {

		suma := 0

		for i := 1; i < arr[iArray]; i++ {
			if arr[iArray]%i == 0 {
				suma += i
			}
		}

		if arr[iArray] == suma {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func main() {

	var T int
	fmt.Scan(&T)

	// Create a slice to store the test cases
	testCases := make([]int, T)

	// Read each test case and store it in the slice
	for i := 0; i < T; i++ {
		fmt.Scan(&testCases[i])
	}

	perfection(testCases)
}
