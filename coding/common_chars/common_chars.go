package main

import "fmt"

func solution(s1 string, s2 string) int {

	s2Slice := make([]string, 0)
	matchs := 0

	//Create an string slice to save s2 chars
	for i := 0; i < len(s2); i++ {
		s2Slice = append(s2Slice, string(s2[i]))
	}

	fmt.Println(s2Slice)

	//Iterate s1, get each char, and find it in s2 slice
	for i := 0; i < len(s1); i++ {

		for j := 0; j < len(s2Slice); j++ {

			//If it's a match, increase counter, delete char in the s2 slice
			if string(s1[i]) == s2Slice[j] {
				fmt.Println(string(s1[i]))
				matchs++
				s2Slice[j] = ""
				break
			}
		}
	}

	return matchs
}

func main() {
	got := solution("aabcc", "adcaa")
	fmt.Println(got)
}
