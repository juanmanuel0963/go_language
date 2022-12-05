package programs

import "fmt"

func ListSwap(mySlice []int) []int {

	for i, j := 0, len(mySlice)-1; i < j; i, j = i+1, j-1 {
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
		fmt.Println("Si:", mySlice[i])
	}

	return mySlice
}
