package main

import "fmt"

//Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue having an non-negative integer size. Since he likes to make things perfect, he wants to arrange them from smallest to largest so that each statue will be bigger than the previous one exactly by 1. He may need some additional statues to be able to accomplish that. Help him figure out the minimum number of additional statues needed.
//Example
//For statues = [6, 2, 3, 8], the output should be
//solution(statues) = 3.
//Ratiorg needs statues of sizes 4, 5 and 7.

func solution(statues []int) int {

	var sorted []int

	for i := 0; i < len(statues); i++ {

		sorted = sortArray(statues)

	}

	fmt.Println(sorted)

	holes := 0

	for i := 0; i < len(sorted)-1; i++ {

		valCurrent := sorted[i]
		valNext := sorted[i+1]

		//fmt.Println(valCurrent)
		//fmt.Println(valNext)

		if (valNext - valCurrent) > 1 {
			holes += valNext - valCurrent - 1
			fmt.Println(holes)
		}

		//fmt.Println("---------")

	}

	return holes
}

func sortArray(inputArray []int) []int {

	for i := 0; i < len(inputArray)-1; i++ {

		if inputArray[i] > inputArray[i+1] {

			//Get the values
			tempMax := inputArray[i]
			tempMin := inputArray[i+1]

			//Swap values
			inputArray[i] = tempMin
			inputArray[i+1] = tempMax
		}
	}

	return inputArray
}

func main() {
	got := solution([]int{6, 2, 3, 8})
	fmt.Println(got)
}
