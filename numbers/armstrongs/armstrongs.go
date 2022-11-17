package armstrongs

import (
	"math"

	"github.com/juanmanuel0963/go_language/v1/numbers/countdigits"
)

// IsArmstrong determines if a number is an armstrong number
func IsArmstrong(number int) bool {
	digits := float64(countdigits.CountDigits(number))

	current := number
	sum := 0
	for current != 0 {
		sum += int(math.Pow(float64(current%10), digits))
		current /= 10
	}
	return sum == number
}
