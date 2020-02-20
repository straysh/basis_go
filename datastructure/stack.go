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
