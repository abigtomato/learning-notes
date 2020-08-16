package queue

// An FIFO Queue.
type Queue []int

// Pushes the element into the queue.
func (q *Queue) Push(elem int) {
	*q = append(*q, elem)
}

// Pops element from head.
func (q *Queue) Pop() int {
	elem := (*q)[0]
	*q = (*q)[1:]
	return elem
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}