package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	//Declare a bool array of the same lenght as candies
	var greatest []bool
	var max int

	//Iterate the candies to find the max value
	for i := 0; i < len(candies); i++ {
		if (candies[i]) > max {
			max = candies[i]
		}
	}

	//Iterate the candies to add the extraCandies and evaluate if is greates than the max value
	for i := 0; i < len(candies); i++ {

		if (candies[i])+extraCandies >= max {
			greatest = append(greatest, true)
		} else {
			greatest = append(greatest, false)
		}
	}

	return greatest
}

func main() {

	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	got := kidsWithCandies(candies, extraCandies)
	fmt.Println(got)
}
