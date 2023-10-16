package main

import "fmt"

func solution(a []int) []int {

	weights := make([]int, 0)

	sum1 := 0
	sum2 := 0

	for i := 0; i < len(a); i++ {

		if i%2 == 0 {
			sum1 = sum1 + a[i]
		} else {
			sum2 = sum2 + a[i]
		}
	}

	weights = append(weights, sum1)
	weights = append(weights, sum2)

	return weights
}

func main() {
	got := solution([]int{50, 60, 60, 45, 70})
	fmt.Println(got)
}
