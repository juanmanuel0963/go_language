package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringsConcat_v1(t *testing.T) {
	// ZERO-VALUE:
	//
	// It's ready to use from the get-go.
	// You don't need to initialize it.
	// There is an efficient and recommended way to concatenate string values and that is to use the strings.Builder type, which was introduced in Go 1.10.
	var sb strings.Builder

	for i := 0; i < 100; i++ {
		sb.WriteString(strconv.Itoa(i) + " ")
	}

	fmt.Println(sb.String())

}

func TestStringsConcat_v2(t *testing.T) {

	// Creating and initializing strings
	// using var keyword
	var str1 = "Hello "

	var str2 = "Reader!"

	// Concatenating strings
	// Using + operator
	fmt.Println("New string 1: ", str1+str2)

	// Creating and initializing strings
	// Using shorthand declaration
	str3 := "Welcome"
	str4 := "the Team"

	// Concatenating strings
	// Using + operator
	result := str3 + " to " + str4

	fmt.Println("New string 2: ", result)
}
