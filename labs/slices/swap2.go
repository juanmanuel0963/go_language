package slices

import "fmt"

func SliceSwap2_v1(my_slice []int) []int {

	j := len(my_slice) - 1

	for i := 0; i < j; i = i + 1 {

		my_slice[i], my_slice[j] = my_slice[j], my_slice[i]

		fmt.Println(my_slice[i])

		j = j - 1
	}

	fmt.Println(my_slice)

	return my_slice
}

func SliceSwap2_v2(my_slice []int) []int {

	for i, j := 0, len(my_slice)-1; i < j; i, j = i+1, j-1 {

		my_slice[i], my_slice[j] = my_slice[j], my_slice[i]

		fmt.Println(my_slice[i])

	}

	fmt.Println(my_slice)

	return my_slice
}
