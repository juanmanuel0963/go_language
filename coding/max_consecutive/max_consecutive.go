package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func longestOnes(nums []int, k int) int {

	//fmt.Println(k)

	var consecutiveCount = 0
	var maxConsecutiveCount = 0
	var replaceCount = 0

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {

			if nums[j] == 1 {
				//Continúa la cuenta
				consecutiveCount += 1
				//Se verifica si es la máxima cuenta históricamente
				if consecutiveCount > maxConsecutiveCount {
					maxConsecutiveCount = consecutiveCount
				}
			} else if nums[j] == 0 {

				if k == 0 {
					//Se resetea la cuenta
					replaceCount = 0
					consecutiveCount = 0
				} else if replaceCount < k {
					//Continúa la cuenta
					replaceCount += 1
					consecutiveCount += 1
				} else {
					//fmt.Println(j)
					//fmt.Println(consecutiveCount)
					//fmt.Println("---")

					//Se resetea la cuenta
					replaceCount = 0
					consecutiveCount = 0

					//Inicia la cuenta
					replaceCount += 1
					consecutiveCount += 1
				}
				//Se verifica si es la máxima cuenta históricamente
				if consecutiveCount > maxConsecutiveCount {
					maxConsecutiveCount = consecutiveCount
				}
			}
		}

		//Se resetea la cuenta
		replaceCount = 0
		consecutiveCount = 0
	}

	//fmt.Println(consecutiveCount)
	//fmt.Println("---")

	//fmt.Println("-end-")
	return maxConsecutiveCount
}

func main() {
	//got := solution(0)

	var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	var k = 2

	//var nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	//var k = 3

	got := longestOnes(nums, k)

	fmt.Println(got)
}
