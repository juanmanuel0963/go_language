package stacks

type Stack []int

func (s Stack) Empty() bool {

	return len(s) == 0

}

func (s Stack) Push(v int) Stack {

	s = append(s, v)
	return s
}

func (s Stack) Pop() (Stack, int) {

	v := (s)[len(s)-1]

	(s) = (s)[:len(s)-1]

	return s, v
}

func (s Stack) GetValues() []int {
	return s
}
