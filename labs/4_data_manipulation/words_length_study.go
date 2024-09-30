package main

import "fmt"

func GroupWordsByLengthWithCountStudy(input []string) (map[int][]string, map[int]int, error) {

	//fmt.Println(input)

	grouped := make(map[int][]string)
	count := make(map[int]int)

	for _, word := range input {

		len := len(word)
		//fmt.Println(index)

		grouped[len] = append(grouped[len], word)
		count[len]++
	}

	return grouped, count, nil
}

func main() {

	words := make([]string, 0)

	words = append(words, "apple")
	words = append(words, "bannana")
	words = append(words, "pineapple")
	words = append(words, "bannana")

	gotGrouped, gotCount, err := GroupWordsByLengthWithCountStudy(words)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gotGrouped)
	fmt.Println(gotCount)
}
