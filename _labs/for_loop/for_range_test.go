package for_loop

import (
	"fmt"
	"testing"
)

func TestForRangeSlice(t *testing.T) {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

}

func TestForRangeArray(t *testing.T) {

	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 9 {
			fmt.Println("index:", i)
		}
	}

}

func TestForRangeMap(t *testing.T) {

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "golang" {
		fmt.Printf("%d -> %d\n", i, c)
	}
}

func TestForRangeChannel(t *testing.T) {

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
