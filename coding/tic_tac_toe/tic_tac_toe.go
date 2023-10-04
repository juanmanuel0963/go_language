package main

import "fmt"

var matrix = [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}

func checkIsWinner(player string) bool {

	count := 0

	//Checking horizontally
	for i := 0; i < len(matrix[0]); i++ {

		for j := 0; j < len(matrix); j++ {

			if matrix[i][j] == player {
				count += 1
			}

		}

		if count == 3 {
			return true
		} else {
			count = 0
		}

	}

	if count > 0 {
		count = 0
	}

	//Checking vertically
	for i := 0; i < len(matrix[0]); i++ {

		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == player {
				count += 1
			}
		}

		if count == 3 {
			return true
		} else {
			count = 0
		}

	}

	if count > 0 {
		count = 0
	}

	//Checking diagonal
	for i := 0; i < len(matrix[0]); i++ {

		for j := 0; j < len(matrix); j++ {

			if i == j { //(1,1) / (2,2) / (3,3)
				if matrix[j][i] == player {
					count += 1
				}
			}

			//(2,0) //Bottom, left
			if i == len(matrix) && j == 0 {
				if matrix[j][i] == player {
					count += 1
				}
			}

			// (0,2) //Top, right
			if i == 0 && j == len(matrix[0]) {

				if matrix[j][i] == player {
					count += 1
				}
			}

			if count == 3 {
				return true
			} else {
				count = 0
			}
		}
	}

	return false
}

func solution(row int, col int, player string) bool {

	//Check if the position is empty
	current := matrix[row][col]
	//valid := false

	if current == "" {
		//Mark the position with the player name (x/y)
		matrix[row][col] = player
		//valid = true
	}

	fmt.Println(matrix)

	IsWinner := checkIsWinner(player)

	return IsWinner
}

func main() {

	fmt.Println(matrix)

	got := solution(0, 0, "X")
	fmt.Println(got)

	got = solution(0, 1, "Y")
	fmt.Println(got)

	got = solution(1, 0, "X")
	fmt.Println(got)

	got = solution(1, 2, "Y")
	fmt.Println(got)

	got = solution(2, 0, "X")
	fmt.Println(got)
}

/*
solve(0,0, "X") returns false
solve(0,1, "Y") returns false
solve(1,0, "X") returns false

solve(1,2, "Y") return false
solve(2,0, "X") returns true
*/
