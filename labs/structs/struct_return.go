package structs

type MyStruct struct {
	Val int
}

// A. Since the function has a return type of the struct, the function returns a copy of the
// struct by setting the value as 1.
func StructReturn_v1() MyStruct {
	return MyStruct{Val: 1}
}

// B. Since the function returns *MyStruct , which is a pointer to the struct, it returns
// a pointer to the struct value created within the function.
func StructReturn_v2() *MyStruct {
	return &MyStruct{Val: 2}
}

// C. Since the function expects the existing struct object as a parameter and in the
// function, we are setting the value of its attribute, at the end of execution the value of
// Val variable of the struct object is set to 3.
func StructReturn_v3(s *MyStruct) {
	s.Val = 3
}
