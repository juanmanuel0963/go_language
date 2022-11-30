package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenghtCapacity(t *testing.T) {

	// ----------Creating an array------------
	array := [6]string{"This", "is", "a", "Go", "interview", "question"}

	// Print array
	fmt.Println("Original Array:", array)

	// Length of Original Array using len()
	arrayLen := len(array)

	// Length of Original Array using len()
	sFormat := fmt.Sprintf("Orignal Array Len: %v", arrayLen)
	fmt.Println(sFormat)

	assert.Equal(t, 6, arrayLen)

	// ----------Creating a slice-------------
	slicedArray := array[1:4] //Pointer to position 1 until position 4-1 = 3

	// Display slice
	fmt.Println("Sliced Array:", slicedArray)

	// Length of slice calculated using len()--------- Positions 1, 2, 3
	sliceArrayLen := len(slicedArray)

	// Length of slice calculated using len()
	sFormat = fmt.Sprintf("Length of the slice: %v", sliceArrayLen)
	fmt.Println(sFormat)
	assert.Equal(t, 3, sliceArrayLen)

	// Capacity of slice calculated using cap()------- Positions 1, 2, 3, 4, 5
	sliceArrayCap := cap(slicedArray)

	// Capacity of slice calculated using cap()
	sFormat = fmt.Sprintf("Capacity of the slice: %v", sliceArrayCap)
	fmt.Println(sFormat)
	assert.Equal(t, 5, sliceArrayCap)

}
