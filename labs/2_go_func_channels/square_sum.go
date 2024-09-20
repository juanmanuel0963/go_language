package main

import (
	"math"
	"sync"
)

//Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

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
