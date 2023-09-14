package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ComputeClosestToZero(ts []int) int {
	if len(ts) == 0 {
		// Return 0 if the slice is empty
		return 0
	}

	closest := ts[0] // Assume the first element is the closest to zero

	for _, num := range ts {
		if abs(num) < abs(closest) || (abs(num) == abs(closest) && num > 0) {
			// If the current number's absolute value is smaller than the current closest,
			// or if their absolute values are equal but the current number is positive,
			// update the closest number
			closest = num
		}
	}

	fmt.Println(closest)

	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	ts := make([]int, n)

	for i := 0; i < n; i++ {
		ts[i], _ = strconv.Atoi(inputs[i])
	}

	fmt.Println(ts)

	stdoutWriter := os.Stdout
	os.Stdout = os.Stderr
	solution := ComputeClosestToZero(ts)
	os.Stdout = stdoutWriter
	fmt.Println(solution)
}
