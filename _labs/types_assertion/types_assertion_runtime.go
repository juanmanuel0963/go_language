package miscellaneous

import "fmt"

func VariableTypeAssertionRuntime(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Double %v is %v\n", v, v*2)
	case uint32:
		fmt.Println("uint32 type")
	case uint64:
		fmt.Println("uint64 type")
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))

		s := i.(string)
		fmt.Println(s)

		s, ok := i.(string)
		fmt.Println(s, ok)

		f, ok := i.(float64)
		fmt.Println(f, ok)

	default:
		fmt.Printf("I don't know  type %T!\n", v)
	}

}
