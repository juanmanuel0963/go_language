package queues

type Queue []int

func (q Queue) Empty() bool {

	return len(q) == 0
}

func (q *Queue) Enqueue(v int) {

	(*q) = append((*q), v)
}

func (q *Queue) Dequeue() int {

	v := (*q)[0]

	(*q) = (*q)[1:len(*q)]

	return v
}

func (q Queue) GetValues() []int {

	return q
}
