package slice

import "fmt"

func SliceReverse_v1(mySlice []int) []int {

	fmt.Println("SliceReverse_v1")

	j := len(mySlice)
	for i := 0; i <= len(mySlice)-1; i = i + 1 {
		j = j - 1
		if i < j {
			mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
			fmt.Println("Si:", mySlice[i])
		}
	}
	return mySlice
}

func SliceReverse_v2(mySlice []int) []int {

	fmt.Println("SliceReverse_v2")

	j := len(mySlice)
	for i := 0; i < j; i = i + 1 {
		j = j - 1
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
		fmt.Println("Si:", mySlice[i])

	}
	return mySlice
}

func SliceReverse_v3(mySlice []int) []int {

	fmt.Println("SliceReverse_v3")

	for i, j := 0, len(mySlice)-1; i < j; i, j = i+1, j-1 {
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
		fmt.Println("Si:", mySlice[i])
	}
	return mySlice
}
