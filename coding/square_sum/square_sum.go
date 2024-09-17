package SquareSum

import (
	"math"
	"sync"
)

//Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

func judgeSquareSum(c int) bool {
	maxC := math.Sqrt(float64(c))
	sqrtMap := make(map[int]struct{})
	for a := 0; a <= int(maxC); a++ {
		sqrtMap[a*a] = struct{}{}
	}

	for key := range sqrtMap {
		target := c - key
		if _, ok := sqrtMap[target]; ok {
			return true
		}
	}
	return false
}

func judgeSquareSumVer2(c int) bool {
	start := 0
	end := int(math.Sqrt(float64(c)))

	for start <= end {
		sum := start*start + end*end
		if sum == c {
			return true
		} else if sum < c {
			start++
		} else {
			end--
		}
	}

	return false
}

func judgeSquareSumVer3(c int) bool {
	if c < 0 {
		return false
	}

	sqrtC := int(math.Sqrt(float64(c)))
	resultChan := make(chan bool)
	var wg sync.WaitGroup

	for a := 0; a <= sqrtC; a++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			b := int(math.Sqrt(float64(c - a*a)))
			if a*a+b*b == c {
				resultChan <- true
			}
		}(a)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result {
			return true
		}
	}
	return false
}
