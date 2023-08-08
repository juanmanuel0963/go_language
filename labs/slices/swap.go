package slices

import "fmt"

func SliceSwap_v0(mySlice []int) []int {

	mySlice[0], mySlice[1] = mySlice[1], mySlice[0]

	fmt.Println("SliceSwap_v0: ")
	fmt.Println(mySlice)

	return mySlice
}

func SliceSwap_v1(mySlice []int) []int {

	fmt.Println()
	fmt.Println("SliceSwap_v1: ")

	j := len(mySlice)
	for i := 0; i <= len(mySlice)-1; i = i + 1 {
		j = j - 1
		if i < j {
			mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
			to_print := fmt.Sprintf("S[%v]: %v", i, mySlice[i])
			fmt.Println(to_print)
		}
	}
	fmt.Println(mySlice)
	return mySlice
}

func SliceSwap_v2(mySlice []int) []int {

	fmt.Println()
	fmt.Println("SliceSwap_v2: ")

	j := len(mySlice)
	for i := 0; i < j; i = i + 1 {
		j = j - 1
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
		to_print := fmt.Sprintf("S[%v]: %v", i, mySlice[i])
		fmt.Println(to_print)
	}
	fmt.Println(mySlice)
	return mySlice
}

func SliceSwap_v3(mySlice []int) []int {

	fmt.Println()
	fmt.Println("SliceSwap_v3: ")

	for i, j := 0, len(mySlice)-1; i < j; i, j = i+1, j-1 {
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
		to_print := fmt.Sprintf("S[%v]: %v", i, mySlice[i])
		fmt.Println(to_print)
	}
	fmt.Println(mySlice)
	return mySlice
}
