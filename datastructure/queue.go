package datastructure

type queueBT struct {
	data []*BTNode
	size int
}

func NewQueueBT(size int) *queueBT {
	instance := &queueBT{}
	instance.data = make([]*BTNode, size)

	return instance
}

func (q *queueBT) Size() int {
	return q.size
}

func (q *queueBT) Enqueue(node *BTNode) {
	q.data[q.size] = node
	q.size++
}

func (q *queueBT) Dequeue() *BTNode {
	node := q.data[0]

	temp := q.data[1:q.size]
	copy(q.data, temp)
	q.data[q.size] = nil
	q.size--

	return node
}
