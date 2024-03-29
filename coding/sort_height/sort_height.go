package main

import "fmt"

//Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees. People can be very tall!

//Example

//For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
//solution(a) = [-1, 150, 160, 170, -1, -1, 180, 190].

func sortSlice(toSort []int) []int {

	for i := 0; i < len(toSort)-1; i++ {

		if toSort[i] != -1 && toSort[i+1] != -1 {

			current := toSort[i]
			next := toSort[i+1]

			if current > next {
				toSort[i] = next
				toSort[i+1] = current
			}
		}
	}

	return toSort
}

func solution(a []int) []int {

	unsortedPersons := make([]int, 0)
	sortedPersons := make([]int, 0)

	//Iterate the input slice
	for i := 0; i < len(a); i++ {
		//fmt.Println(a[i])
		//Get persons
		if a[i] > -1 {
			unsortedPersons = append(unsortedPersons, a[i])
		}
	}

	//Iterate the unsorted persons
	for i := 0; i < len(unsortedPersons); i++ {
		//Get persons sorted
		sortedPersons = sortSlice(unsortedPersons)
	}

	persons := 0
	//Iterate the input slice
	for i := 0; i < len(a); i++ {

		//if it's a person, replace the height
		if a[i] != -1 {
			a[i] = sortedPersons[persons]
			persons += 1
		}
	}

	//fmt.Println(sortedPersons)
	//fmt.Println(a)

	return a

}

func main() {
	got := solution([]int{-1, 150, 190, 170, -1, -1, 160, 180})
	fmt.Println(got)
}
