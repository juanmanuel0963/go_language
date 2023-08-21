package formating

import (
	"fmt"
	"testing"
)

func TestFormatting(t *testing.T) {

	s := fmt.Sprintf("Size: %d MB.", 50)
	fmt.Println(s)

	s = fmt.Sprintf("%d %s\n", 100, "MB.")
	fmt.Println(s)

	s = fmt.Sprintf("Orignal Array Len: %v", 20000)
	fmt.Println(s)

	const (
		resultFormat = "%d.%s"
	)

	quotient := 100 / 5
	fractionString := "0"

	s = fmt.Sprintf(resultFormat, quotient, fractionString)
	fmt.Println(s)
}
