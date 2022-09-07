package stack

type StackList struct {
	elements []int
}

func (s *StackList) New() {
	s.elements = []int{}
}

func (s *StackList) Push(number int) int {
	top := s.Top()
	s.elements = append(s.elements, number)
	return top
}

func (s *StackList) Pop() {
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *StackList) Top() int {
	return s.elements[len(s.elements)-1]
}

func (s *StackList) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackList) Size() int {
	return len(s.elements)
}
