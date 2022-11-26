package miscellaneous

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringConcat(t *testing.T) {
	// ZERO-VALUE:
	//
	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var sb strings.Builder

	for i := 0; i < 1000; i++ {
		sb.WriteString(strconv.Itoa(i) + " ")
	}

	fmt.Println(sb.String())
}
