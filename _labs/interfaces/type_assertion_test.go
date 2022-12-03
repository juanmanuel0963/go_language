package interfaces

import (
	"fmt"
	"testing"
)

//type error_name interface {
//	Error() string
//}

func TestTypeAssertion_v1(t *testing.T) {

	var param interface{} = "hello"

	s := param.(string)
	fmt.Println(s)

	s, ok := param.(string)
	fmt.Println(s, ok)

	f, ok := param.(float64)
	fmt.Println(f, ok)
}
func TestTypeAssertion_v2(t *testing.T) {

	var param interface{} = 100

	switch v := param.(type) {
	case uint32:
		fmt.Println("Integer type")
	case uint64:
		fmt.Println("Integer type")
	case string:
		fmt.Println("String type")
	default:
		fmt.Printf("Unexpected type %T", v)
	}

}
