package palindromestrings

import "github.com/juanmanuel0963/go_language/v1/advance/numbers/reverses"

// IsPalindrome determines if the input is a palindrome
func IsPalindrome(number int) bool {
	return number == reverses.Reverse(number)
}
