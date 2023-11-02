package main

import "fmt"

func solution(inputIntervals [][]int, inputNewInterval []int) [][]int {

	output := make([][]int, 0)

	var outputMin = 0
	var outputMax = 0

	for i := 0; i < len(inputIntervals); i++ {

		//fmt.Println(interval)

		if inputIntervals[i][0] <= inputNewInterval[0] && inputNewInterval[0] <= inputIntervals[i][1] && outputMax == 0 {

			outputMin = inputIntervals[i][0]
			fmt.Printf("[ min:%v ]  \n", outputMin)

		} else if i-1 >= 0 && inputIntervals[i-1][1] < inputNewInterval[0] && inputNewInterval[0] < inputIntervals[i][0] && outputMax == 0 {

			outputMin = inputIntervals[i][0]
			fmt.Printf("[] min:%v []  \n", outputMin)

		} else if i-1 >= 0 && inputIntervals[i-1][1] < inputNewInterval[1] && inputNewInterval[1] < inputIntervals[i][0] {

			outputMax = inputNewInterval[1]
			fmt.Printf("[] max:%v []  \n", outputMax)

			if outputMin > 0 && outputMax > 0 {
				output = append(output, []int{outputMin, outputMax})
				fmt.Printf("new add [%v,%v] \n", outputMin, outputMax)
				outputMin = 0
				outputMax = 0

				output = append(output, inputIntervals[i])
				fmt.Printf("no touched %v \n", inputIntervals[i])
			}

		} else if inputIntervals[i][0] <= inputNewInterval[1] && inputNewInterval[1] <= inputIntervals[i][1] {

			outputMax = inputIntervals[i][1]
			fmt.Printf("[ max:%v ]  \n", outputMax)

		}

		if outputMin > 0 && outputMax > 0 {
			output = append(output, []int{outputMin, outputMax})
			fmt.Printf("new add [%v,%v] \n", outputMin, outputMax)
			outputMin = 0
			outputMax = 0
		}
	}

	fmt.Println(output)
	//modify := [][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}, {10, 11}}
	return output
}

func main() {

	current := [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}
	new := []int{3, 4}

	got := solution(current, new)
	fmt.Println(got)
}
