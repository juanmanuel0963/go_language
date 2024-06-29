package pointers

import (
	"fmt"
	"testing"
)

// What is the output of the following code snippet?
func TestMCQ_q1(t *testing.T) {
	x := 3
	y := &x

	fmt.Println(*y) //3

	*y = 4 //4

	fmt.Println(x) //4
}

func TestPointers_v1(t *testing.T) {

	x := 5
	p := &x //p=ce3z
	q := *p // q=5

	fmt.Printf("p = %d", *p) //p=5
	fmt.Println()
	fmt.Printf("q = %d", q) //q=5
	fmt.Println()
}

func TestPointers_v2(t *testing.T) {

	//Go provides the & (ampersand) operator also known as the address of operator.
	lang := "Golang"
	fmt.Println("lang address:", &lang) // -> 0xc00010a040

	var x int = 2                                                                   // -> int value
	ptr := &x                                                                       // -> pointer to int
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, *ptr, &ptr) // -> the pointer is saved at an address as well

}
