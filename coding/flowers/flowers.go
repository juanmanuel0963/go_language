package main

import (
	"fmt"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {

		if i == 0 { //start
			//fmt.Println("istart = " + strconv.Itoa(i) + ", value = " + strconv.Itoa(flowerbed[i]))

			if flowerbed[i] == 0 && len(flowerbed) == 1 {
				flowerbed[i] = 1
				n = n - 1
			} else if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n = n - 1
			}
		} else if i == len(flowerbed)-1 { //end
			//fmt.Println("iend = " + strconv.Itoa(i) + ", value = " + strconv.Itoa(flowerbed[i]))

			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n = n - 1
			}
		} else {
			//fmt.Println("inormal = " + strconv.Itoa(i) + ", value = " + strconv.Itoa(flowerbed[i]))

			if (flowerbed[i-1] == 0) && (flowerbed[i] == 0 && flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n = n - 1
			}
		}

		if n == 0 {
			return true
		}

	}

	return false
}

func main() {
	//got := solution(0)
	//fmt.Println(got)

	//var flowerbed = []int{1, 0, 0, 0, 1}
	//var n int = 1

	//var flowerbed = []int{0}
	//var n int = 1

	//var flowerbed = []int{0, 0, 0, 0, 0, 1, 0, 0}
	//var n int = 0

	var flowerbed = []int{0}
	var n int = 1

	got := canPlaceFlowers(flowerbed, n)
	fmt.Println(got)
}
