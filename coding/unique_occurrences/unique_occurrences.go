package main

import "fmt"

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func uniqueOccurrences(arr []int) bool {

	//fmt.Println(arr)
	counterNumbersMap := make(map[int]int, 0)
	counterOcurrencesMap := make(map[int]int, 0)

	for _, val := range arr {
		//fmt.Println(ind)
		//fmt.Println(val)
		counterNumbersMap[val] = counterNumbersMap[val] + 1
	}

	for _, val := range counterNumbersMap {
		//fmt.Println(val)
		counterOcurrencesMap[val] = counterOcurrencesMap[val] + 1

		if counterOcurrencesMap[val] > 1 {
			return false
		}
	}

	//fmt.Println(counterNumbersMap)
	//fmt.Println(counterOcurrencesMap)

	return true
}

func main() {

	arr := []int{1, 2, 2, 1, 1, 3}
	got := uniqueOccurrences(arr)
	fmt.Println(got)
}
