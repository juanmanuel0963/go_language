package main

import (
	"strconv"
)

func lexicographically_smaller(s string, t string) int {

	//Go through s and check each char, if it's a number supress the char and save the new string in sList
	sList := []string{}

	for i := 0; i < len(s); i++ {

		//If s[i] is a number
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			sText := s[:i] + s[i+1:]
			sList = append(sList, sText)
		}
	}

	//fmt.Println(sList)

	//Go through t and check each char, if it's a number supress the char and sabe the new string in tList
	tList := []string{}

	for i := 0; i < len(t); i++ {

		//If s[i] is a number
		if _, err := strconv.Atoi(string(t[i])); err == nil {
			sText := t[:i] + t[i+1:]
			tList = append(tList, sText)
		}
	}

	//fmt.Println(tList)
	count := 0
	//Compare sList vs t
	for _, valueS := range sList {
		//fmt.Println(valueS)

		if valueS < t {
			//fmt.Println(valueS + " < " + t)
			count++
		}
	}

	//Compare s vs tList
	for _, valueT := range tList {

		//fmt.Println(valueT)

		if s < valueT {
			//fmt.Println(s + " < " + valueT)
			count++
		}
	}
	return count
}
