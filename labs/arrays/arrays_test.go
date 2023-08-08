package arrays

import (
	"fmt"
	"testing"
)

//Explain ARRAY and SLICE types and the differences between them.
//---------------------------------------------------------------

//- Golang has two data structures to handle lists of records: ARRAYS and SLICES.

//- An ARRAY is a composite, indexable type that stores a collection of elements.
//- An ARRAY has a fixed length.
//- We specify how many items are in the ARRAY when we declare it.
//- The ARRAY length is part of its type, so arrays cannot be resized.

// - This is in contrast to a SLICE that has a dynamic length (it can shrink or grow at runtime).
// - Every element in an array or slice must be of the same type.
// - Slices are a key data type in Golang and are everywhere.
func TestArrays(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Here’s an example of declaring and initializing an array of type [4] string using the short declaration syntax.
	friends := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Println(friends)

}
