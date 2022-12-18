package types_version

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeConversion_v1(t *testing.T) {
	//Float To Int
	i := 56                //int
	j := 67.8              //float64
	float_to_int := int(j) //j is converted to int
	fmt.Println(float_to_int)
	sum := i + float_to_int
	fmt.Printf("Type of sum is %T", sum)
	fmt.Println()

	assert.Equal(t, 123, sum)

	//Int to String
	int_to_str := strconv.Itoa(55) //i is converted to string
	fmt.Println(int_to_str)
	fmt.Printf("Type of int_to_str is %T", int_to_str)
	fmt.Println()
	assert.Equal(t, "55", int_to_str)

	//String to Int

	string_to_int, err := strconv.Atoi("42")
	fmt.Println(string_to_int)

	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Printf("Type of string_to_int is %T", string_to_int)
	fmt.Println()
	assert.Equal(t, 42, string_to_int)
}

func TestTypeConversion_Float_to_Int(t *testing.T) {
	var floatValue float32 = 10.65
	// type conversion from float to int
	var intValue int = int(floatValue)
	fmt.Printf("Float Value: %g\n", floatValue)
	fmt.Printf("Integer Value: %d\n", intValue)
}

func TestTypeConversion_Int_To_Float(t *testing.T) {
	var intValue int = 10
	// type conversion from int to float
	var floatValue float32 = float32(intValue)
	fmt.Printf("Integer Value: %d\n", intValue)
	fmt.Printf("Float Value: %f\n", floatValue)
}

func TestTypeConversion_Add_Int_and_Float(t *testing.T) {
	var num1 int = 90
	var num2 float32 = 8.7
	var sum float32 = 0.0

	// addition of different data types
	sum = float32(num1) + num2

	fmt.Printf("Sum is %g\n", sum)
}
