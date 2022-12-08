package stacks

type StackPointer []int

func (s StackPointer) EmptyPointer() bool {

	return len(s) == 0

}

func (s *StackPointer) PushPointer(v int) {

	(*s) = append((*s), v)

}

func (s *StackPointer) PopPointer() int {

	v := (*s)[len(*s)-1]

	(*s) = (*s)[:len(*s)-1]

	return v
}

func (s StackPointer) GetValuesPointer() []int {
	return s
}
