package programs

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})
}
