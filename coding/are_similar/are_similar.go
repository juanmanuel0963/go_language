package main

import "fmt"

func solution(a []int, b []int) bool {

	if compareSlices(a, b) {
		return true
	}

	swap := make([]int, len(a))

	for i := 0; i < len(a); i++ {

		for j := 0; j < len(a); j++ {

			if i != j {

				copy(swap, a)

				swap[i] = a[j]
				swap[j] = a[i]

				if compareSlices(swap, b) {
					return true
				}
			}
		}
	}

	return false
}

func compareSlices(a []int, b []int) bool {

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {

	got1 := solution([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(got1)

	got2 := solution([]int{1, 2, 3}, []int{2, 1, 3})
	fmt.Println(got2)

	got3 := solution([]int{1, 2, 2}, []int{2, 1, 1})
	fmt.Println(got3)
}
