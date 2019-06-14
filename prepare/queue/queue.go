package queue

type Queue []int

//入队
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

//出队
func (q *Queue) Pop() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func IsEmpty(q Queue) bool {
	return len(q) == 0
}
