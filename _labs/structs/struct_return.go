package structs

type MyStruct struct {
	Val int
}

func StructReturn_v1() MyStruct {
	return MyStruct{Val: 1}
}

func StructReturn_v2() *MyStruct {
	return &MyStruct{Val: 2}
}

func StructReturn_v3(s *MyStruct) *MyStruct {
	s.Val = 3
	return s
}
