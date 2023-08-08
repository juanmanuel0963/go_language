package slices

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {

	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))

}
