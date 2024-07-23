package main

import (
	"fmt"
)

func solution(input int) int {

	fmt.Println(input)
	return 0
}

func isSubsequence(s string, t string) bool {

	//fmt.Println(s)
	//fmt.Println(t)
	counterFound := 0
	start := 0

	for i := 0; i < len(s); i++ {
		//fmt.Println("Starting at " + strconv.Itoa(start))
		for j := start; j < len(t); j++ {
			if s[i] == t[j] {
				//fmt.Println("Found " + string(s[i]))
				counterFound += 1
				start = j + 1
				break
			}
		}
	}
	//fmt.Println(counterFound)
	return counterFound == len(s)
}

func main() {
	s := "aaaaaa"
	t := "bbaaaa"
	got := isSubsequence(s, t)
	fmt.Println(got)
}
