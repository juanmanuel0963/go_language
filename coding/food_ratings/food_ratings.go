package main

import (
	"fmt"
	"strconv"
)

func solution(N int, ratings [][]int) int {

	sum := make(map[int]int)
	frecuency := make(map[int]int)
	average := make(map[int]int)
	highestAverage := 0
	highestAverageID := 0

	//for _, row := range ratings {
	for i := 0; i < N; i++ {
		sum[ratings[i][0]] += ratings[i][1]
		frecuency[ratings[i][0]] += 1
		average[ratings[i][0]] = sum[ratings[i][0]] / frecuency[ratings[i][0]]
	}

	for ID, row := range average {
		if highestAverageID == 0 {
			highestAverage = row
			highestAverageID = ID
		} else if row >= highestAverage && ID < highestAverageID {
			highestAverage = row
			highestAverageID = ID
		}
	}

	return highestAverageID
}

func main() {

	var T int
	fmt.Scan(&T)

	// Create a slice to store the test cases
	testCases := make([]string, T*2)

	// Read each test case and store it in the slice
	for i := 0; i < T*2; i++ {
		fmt.Scan(&testCases[i])
	}

	//fmt.Println(testCases)

	ratings := make([][]int, 0)

	for i := 0; i < T*2; i++ {
		ID, _ := strconv.Atoi(testCases[i])
		rating, _ := strconv.Atoi(testCases[i+1])

		row := []int{ID, rating}
		ratings = append(ratings, row)
		i++
	}

	fmt.Println(solution(T, ratings))
}
