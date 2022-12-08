package slices

import "fmt"

func SortSlice(sliceToSort []int) []int {

	lastIndex := len(sliceToSort) - 1

	for i := 0; i < len(sliceToSort)-1; i++ {
		fmt.Println(lastIndex)
		for j := 0; j < lastIndex; j++ {
			if sliceToSort[j] > sliceToSort[j+1] {
				sliceToSort[j], sliceToSort[j+1] = sliceToSort[j+1], sliceToSort[j]
			}
		}
		lastIndex--
	}

	return sliceToSort
}
