package slice2

import "fmt"

func Slice2Reverse_v1(my_slice []int) []int {

	j := len(my_slice) - 1

	for i := 0; i < j; i = i + 1 {

		my_slice[i], my_slice[j] = my_slice[j], my_slice[i]

		fmt.Println(my_slice[i])

		j = j - 1
	}

	return my_slice
}

func Slice2Reverse_v2(my_slice []int) []int {

	for i, j := 0, len(my_slice)-1; i < j; i, j = i+1, j-1 {

		my_slice[i], my_slice[j] = my_slice[j], my_slice[i]

		fmt.Println(my_slice[i])

	}

	return my_slice
}
