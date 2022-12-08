package slices

import (
	"fmt"
	"testing"
)

func TestSortSlice(t *testing.T) {

	sortData := []int{10, 1, 0, 5, 4, 9, 2, 3, 8, 7, 6}
	fmt.Println(sortData)

	sortedData := SortSlice(sortData)
	fmt.Println(sortedData)
}
