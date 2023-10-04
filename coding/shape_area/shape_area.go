package main

import "fmt"

//Your task is to find the area of a polygon for a given n.

//A 1-interesting polygon is just a square with a side of length 1.

//An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim, side by side.

func solution(n int) int {

	current := n * n
	last := (n - 1) * (n - 1)

	return current + last
}

func main() {
	got := solution(2)
	fmt.Println(got)
}
