package scan

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Wrong "text string" variable
func TestScan_v1(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text string
		n    int
	)

	for scanner.Scan() {
		n++
		text += fmt.Sprintf("%d. %s\n", n, scanner.Text())
		//fmt.Print(text)
	}

	fmt.Print(text)

	// Output:
	// 1. One
	// 2. Two
	// 3. Three
	// 4. Four
}

// Right "text []byte" variable
func TestScan_v2(t *testing.T) {
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
