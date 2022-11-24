package slices

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSlice_v1(t *testing.T) {

	fruits := map[string]int{
		"oranges": 100,
		"apples":  200,
		"bananas": 300,
	}

	// Put the keys in a slice and sort it.
	var keys []string
	for key := range fruits {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Display keys according to the sorted slice.
	for _, key := range keys {
		fmt.Printf("%s:%v\n", key, fruits[key])
	}

}

func TestSortSlice_v2(t *testing.T) {

	fruits := map[string]int{
		"orange": 100, "banana": 200, "mango": 300, "apple": 400,
	}

	//keys := []string{} - do allocate memory
	var keys []string //doesnt't allocate memory

	for key := range fruits {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	//fmt.Print(keys)

	for _, key := range keys {
		fmt.Printf("%s: $%v \n", key, fruits[key])
	}
}
