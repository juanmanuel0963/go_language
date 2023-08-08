package structs

import (
	"fmt"
	"testing"
)

func TestString_v1(t *testing.T) {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)

	fmt.Println(orange)
}

func TestString_v2(t *testing.T) {
	orange := &Orange{}
	orange.Increase(10)
	orange.Decrease(5)

	fmt.Println(orange)

}
