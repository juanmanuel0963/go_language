package functions

import "strings"

// Variadic function to join strings and separate them with hyphen
func JoinStringVariadicFunction(element ...string) string {
	return strings.Join(element, "-")
}
