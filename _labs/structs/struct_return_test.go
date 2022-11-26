package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructReturn_v1(t *testing.T) {

	the_struct := StructReturn_v1()
	fmt.Println(the_struct.Val)

	assert.Equal(t, 1, the_struct.Val)
}

func TestStructReturn_v2(t *testing.T) {

	the_struct := StructReturn_v2()
	fmt.Println(the_struct.Val)

	assert.Equal(t, 2, the_struct.Val)
}

func TestStructReturn_v3(t *testing.T) {

	the_struct := &MyStruct{}
	the_struct_return := StructReturn_v3(the_struct)

	fmt.Println(the_struct_return.Val)

	assert.Equal(t, 3, the_struct_return.Val)
}
