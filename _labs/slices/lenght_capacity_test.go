package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenghtCapacity_v1(t *testing.T) {

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

func TestLenghtCapacity_v2(t *testing.T) {

	//Slice capacity = len(array)
	n1 := []int{10, 20, 30, 40, 50}
	fmt.Println("n1: ", n1)
	fmt.Println("len: ", len(n1))
	fmt.Println("cap: ", cap(n1))
	assert.Equal(t, 5, len(n1))
	assert.Equal(t, 5, cap(n1))
	fmt.Println("--------------")

	//If len(appended) < len(array), capacity = 2 x len(array)
	n2 := []int{10, 20, 30, 40, 50}
	n2 = append(n2, 100)
	n2 = append(n2, 200)
	fmt.Println("n2: ", n2)
	fmt.Println("len: ", len(n2))
	fmt.Println("cap: ", cap(n2))
	assert.Equal(t, 7, len(n2))
	assert.Equal(t, 10, cap(n2))
	fmt.Println("--------------")

	//If len(appended) < len(array), capacity = 2 x len(array)
	n3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	n3 = append(n3, 100)
	n3 = append(n3, 200)
	fmt.Println("n3: ", n3)
	fmt.Println("len: ", len(n3))
	fmt.Println("cap: ", cap(n3))
	assert.Equal(t, 12, len(n3))
	assert.Equal(t, 20, cap(n3))
	fmt.Println("--------------")

	//If len(appended) > len(array), capacity = 2 x (last capacity)
	n4 := []int{10, 20, 30, 40, 50}
	n4 = append(n4, 100)
	n4 = append(n4, 200)
	n4 = append(n4, 300)
	n4 = append(n4, 400)
	n4 = append(n4, 500)
	n4 = append(n4, 600)
	n4 = append(n4, 700)

	fmt.Println("n4: ", n4)
	fmt.Println("len: ", len(n4))
	fmt.Println("cap: ", cap(n4))
	assert.Equal(t, 12, len(n4))
	assert.Equal(t, 20, cap(n4))
	fmt.Println("--------------")

	//If len(appended) > len(array), capacity = 2 x (last capacity)
	n5 := []int{10, 20, 30, 40, 50}
	n5 = append(n5, 100)
	n5 = append(n5, 200)
	n5 = append(n5, 300)
	n5 = append(n5, 400)
	n5 = append(n5, 500)
	n5 = append(n5, 600)
	n5 = append(n5, 700)
	n5 = append(n5, 800)
	n5 = append(n5, 900)
	n5 = append(n5, 1000)
	n5 = append(n5, 1100)
	n5 = append(n5, 1200)
	n5 = append(n5, 1300)
	n5 = append(n5, 1400)
	n5 = append(n5, 1500)
	n5 = append(n5, 1600)
	n5 = append(n5, 1700)
	fmt.Println("n5: ", n5)
	fmt.Println("len: ", len(n5))
	fmt.Println("cap: ", cap(n5))
	assert.Equal(t, 22, len(n5))
	assert.Equal(t, 40, cap(n5))
	fmt.Println("--------------")
}
