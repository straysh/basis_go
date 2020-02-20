package datastructure

type stackInt struct {
	data []int
	size int
}

func NewStackInt(size int) *stackInt {
	instance := &stackInt{}
	instance.data = make([]int, size)

	return instance
}

func (s *stackInt) Pop() int {
	item := s.data[s.size-1]
	s.data = s.data[0 : s.size-1]
	s.size--

	return item
}

func (s *stackInt) Push(item int) {
	s.data[s.size] = item
	s.size++
}

func (s *stackInt) Size() int {
	return s.size
}

type stackBT struct {
	data []*BTNode
	size int
}

func NewStackBT(size int) *stackBT {
	instance := &stackBT{}
	instance.data = make([]*BTNode, size)

	return instance
}

func (s *stackBT) Pop() *BTNode {
	item := s.data[s.size-1]
	s.data = s.data[0 : s.size-1]
	s.size--

	return item
}

func (s *stackBT) Push(item *BTNode) {
	s.data[s.size] = item
	s.size++
}

func (s *stackBT) Size() int {
	return s.size
}
