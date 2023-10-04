package main

import "fmt"

//For sequence = [1, 3, 2, 1], the output should be
//solution(sequence) = false.

//There is no one element in this array that can be removed in order to get a strictly increasing sequence.

//For sequence = [1, 3, 2], the output should be
//solution(sequence) = true.

//You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].
func solution(sequence []int) bool {

	if len(sequence) == 2 {
		return true
	}

	countOne := 0
	countTwo := 0

	for i := 0; i < len(sequence)-1; i++ {

		if sequence[i] >= sequence[i+1] {
			countOne++
		}

		if i != 0 {
			if sequence[i-1] >= sequence[i+1] {
				countTwo++
			}
		}
	}

	if countOne == 1 && countTwo <= 1 {
		return true
	} else {
		return false
	}
}

func main() {
	got := solution([]int{1, 3, 2, 1})
	fmt.Println(got)
}
