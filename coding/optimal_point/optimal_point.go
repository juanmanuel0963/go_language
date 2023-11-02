package main

import (
	"fmt"
	"math"
)

/*
Given
- A rectangular board of size N × M
- 3 or more travelers are placed in different cells of the board
- The travelers can move vertically and horizontally only to adjacent cells
- Moving one traveler to an adjacent cell costs one effort point
Problem
- Write a program that outputs the coordinates of the cell which requires the minimum number of moves of all travelers
- There can be several optimal solutions

Example
Input
1 - 0 - 0 - 0 - 1
|   |   |   |   |
0 - 0 - 0 - 0 - 0
|   |   |   |   |
0 - 0 - 1 - 0 - 0

1 0 X 0 1
0 0 0 0 0
0 0 1 0 0
Output
0, 2
*/

func solution(rows int, cols int, participants [][]int) []int {

	var minDistance float64
	var minRow int
	var minCol int

	//Iterate the input matrix and for each position calculate steps to participants' positions
	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {

			//fmt.Printf("row:%v col:%v \n", i, j)

			var totalDistance float64

			for k := 0; k < len(participants); k++ {
				rowDistance := math.Abs(float64(participants[k][0]) - float64(i))
				colDistance := math.Abs(float64(participants[k][1]) - float64(j))
				totalDistance = totalDistance + rowDistance + colDistance
				//fmt.Printf("rowDistance:%v colDistance:%v \n", rowDistance, colDistance)
			}

			if i == 0 && j == 0 {
				minDistance = totalDistance
				minRow = i
				minCol = j

			} else if totalDistance < minDistance {
				minDistance = totalDistance
				minRow = i
				minCol = j
			}

			//fmt.Printf("totalDistance:%v \n", totalDistance)
		}
	}

	fmt.Printf("\nminDistance:%v \n", minDistance)
	//fmt.Printf("minRow:%v \n", minRow)
	//fmt.Printf("minCol:%v \n", minCol)

	return []int{minRow, minCol}
}

func main() {

	got := solution(3, 5, [][]int{{0, 0}, {0, 4}, {2, 2}})

	fmt.Printf("Optimal point:%v \n", got)
	//fmt.Println(got)
}
