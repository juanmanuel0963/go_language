package main

import (
	"fmt"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func maxArea(height []int) int {

	//fmt.Println(height)
	maxArea := 0

	for i := 0; i < len(height)-1; i++ {

		for j := i + 1; j < len(height); j++ {
			//fmt.Println("i: " + strconv.Itoa(height[i]))
			//fmt.Println("j: " + strconv.Itoa(height[j]))

			minHeight := height[i]
			if height[j] < height[i] {
				minHeight = height[j]
			}
			distance := j - i
			area := minHeight * distance
			if area > maxArea {
				maxArea = area
			}

			//fmt.Println("minHeight: " + strconv.Itoa(minHeight))
			//fmt.Println("distance: " + strconv.Itoa(distance))
			//fmt.Println("area: " + strconv.Itoa(area))
		}
		//fmt.Println("----------")
	}

	//fmt.Println("maxArea: " + strconv.Itoa(maxArea))

	return maxArea
}

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	got := maxArea(height)
	fmt.Println(got)
}
