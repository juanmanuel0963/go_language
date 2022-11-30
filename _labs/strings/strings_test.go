package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {

	var Raw_String_Literals = `uninterrupted string`
	fmt.Println(Raw_String_Literals)

	var Interpreted_String_Literals = "interrupted\nstring"
	fmt.Println(Interpreted_String_Literals)
}

func TestStringsConcat(t *testing.T) {
	// ZERO-VALUE:
	//
	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var sb strings.Builder

	for i := 0; i < 100; i++ {
		sb.WriteString(strconv.Itoa(i) + " ")
	}

	fmt.Println(sb.String())

}
