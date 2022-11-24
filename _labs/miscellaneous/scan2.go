package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text []byte
		n    int
	)
	for scanner.Scan() {
		n++
		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
	}

	os.Stdout.Write(text)
	//fmt.Print(text)
	// 1. One
	// 2. Two
	// 3. Three
	// 4. Four
}
