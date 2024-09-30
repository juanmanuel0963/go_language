package main

func GroupWordsByLengthWithCount(words []string) (map[int][]string, map[int]int) {
	grouped := make(map[int][]string)
	counts := make(map[int]int)
	for _, word := range words {
		length := len(word)
		grouped[length] = append(grouped[length], word)
		counts[length]++
	}
	return grouped, counts
}
