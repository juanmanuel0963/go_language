package main

import "fmt"

func solution(year int) int {
	//Divide year/100

	mod := year % 100
	century := year / 100

	fmt.Println(mod)

	if mod > 0 {
		century = century + 1
	}

	fmt.Println(century)

	return century
}

func main() {

	got := solution(2020)
	fmt.Println(got)
}
