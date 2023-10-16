package main

import (
	"fmt"
	"time"
)

func Solution(A []int) int {

	minVal, maxVal := FindMinMaxVal(A)
	minMissing := 0

	AMap := ArrayToMap(A)

	//fmt.Println(AMap)

	for i := minVal; i <= maxVal; i++ {

		existVal := AMap[i]

		//existVal := FindVal(A, i)

		if !existVal {
			minMissing = i
			break
		}
	}

	if minMissing == 0 {
		minMissing = maxVal + 1
	}

	return minMissing
}

func ArrayToMap(A []int) map[int]bool {

	AMap := make(map[int]bool)

	for _, val := range A {

		AMap[val] = true
	}

	return AMap
}

/*
func FindVal(A []int, findVal int) bool {

		for _, val := range A {

			if findVal == val {
				return true
			}

		}

		return false
	}
*/
func FindMinMaxVal(A []int) (int, int) {

	var minVal = A[0]
	var maxVal = A[0]

	for _, val := range A {

		if val < minVal {
			minVal = val
		}

		if val > maxVal {
			maxVal = val
		}

	}

	if minVal < 0 {
		minVal = 1
	}

	if maxVal < 0 {
		maxVal = 1
	}

	return minVal, maxVal
}

func main() {

	start := time.Now()

	got := Solution([]int{1, 3, 6, 4, 1, 2})
	fmt.Println(got)
	/*
		got1 := Solution([]int{1, 2, 3})
		fmt.Println(got1)

		got2 := Solution([]int{-1, -3})
		fmt.Println(got2)
	*/
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}
