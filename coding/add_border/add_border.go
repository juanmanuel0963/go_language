package main

import "fmt"

func solution(picture []string) []string {

	var output []string

	stars := make([]rune, 0)

	for i := 0; i < len(picture); i++ {

		if i == 0 {
			output = append(output, "")
		}

		output = append(output, "*"+picture[i]+"*")

	}

	output = append(output, "")

	cols := len(picture[0]) + 2

	for j := 0; j < cols; j++ {
		stars = append(stars, '*')
	}

	output[0] = string(stars)
	output[len(output)-1] = string(stars)

	return output
}

func main() {
	got := solution([]string{"abc", "def"})
	fmt.Printf("got: %s \n", got)
}
