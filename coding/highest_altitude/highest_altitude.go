package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func largestAltitude(gain []int) int {

	maxAltitude := 0
	currentAltitude := 0

	for i := 0; i < len(gain); i++ {
		currentAltitude = currentAltitude + gain[i]
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}

func main() {

	//gain := []int{-5, 1, 5, 0, -7}
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	got := largestAltitude(gain)

	fmt.Println(got)
}
