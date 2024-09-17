package heightChecker

import (
	"sort"
	"sync"
)

// A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line
//in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i]
//is the expected height of the ith student in line.
//
//You are given an integer array heights representing the current order that the students are standing in. Each heights[i]
//is the height of the ith student in line (0-indexed).
//
//Return the number of indices where heights[i] != expected[i].

func heightChecker(heights []int) int {
	firstArray := make([]int, len(heights))
	copy(firstArray, heights)

	sort.Ints(heights)

	rightArray := heights

	res := 0
	for i := 0; i < len(heights); i++ {
		if firstArray[i] != rightArray[i] {
			res++
		}
	}

	return res
}

func heightCheckerVer2(heights []int) int {
	count := [10]int{}

	for _, height := range heights {
		count[height]++
	}
	res := 0
	j := 0
	for _, height := range heights {

		for count[j] == 0 {
			j++
		}
		if height != j {
			res++
		}
		count[j]--
	}

	return res
}

func heightCheckerVer3(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	var wg sync.WaitGroup
	resChan := make(chan bool, len(heights))

	for i := 0; i < len(heights); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			if heights[idx] != expected[idx] {
				resChan <- true
			} else {
				resChan <- false
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	res := 0
	for result := range resChan {
		if result {
			res++
		}
	}

	return res
}
