package for_loop

import (
	"fmt"
	"testing"
)

func TestForBreak_v1(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking out of loop")
			break // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}

func TestForBreak_v2(t *testing.T) {

	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("Breaking out of outer loop")
			break // break here
		}
		fmt.Println("The value of outer is", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("Breaking out of inner loop")
				break // break here
			}
			fmt.Println("The value of inner is", inner)
		}
	}
	fmt.Println("Exiting program")
}
