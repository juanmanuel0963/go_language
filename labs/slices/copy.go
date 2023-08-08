package slices

import "fmt"

func SliceCopy(from_slice []int, to_slice []int) []int {

	to_control := to_slice

	fmt.Println("Before Copy To control: ")
	fmt.Println(to_control)

	fmt.Println("From slice: ")
	fmt.Println(from_slice)

	fmt.Println("Before Copy To slice: ")
	fmt.Println(to_slice)

	copy(to_slice, from_slice)

	fmt.Println("After Copy To slice: ")
	fmt.Println(to_slice)

	fmt.Println("After Copy To control: ")
	fmt.Println(to_control)

	return to_slice
}
