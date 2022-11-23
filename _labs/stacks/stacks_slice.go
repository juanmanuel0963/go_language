package stacks

type Stack []int

func (s Stack) Empty() bool { return len(s) == 0 }

func (s *Stack) Push(v int) { (*s) = append((*s), v) }

func (s *Stack) Pop() int {

	v := (*s)[len(*s)-1]

	(*s) = (*s)[:len(*s)-1]

	return v
}

func (s Stack) GetValues() []int {
	return s
}
