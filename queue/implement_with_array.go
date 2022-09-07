package stack

type QueueList struct {
	elements []int
}

func (s *QueueList) New() {
	s.elements = []int{}
}

func (s *QueueList) Enqueue(number int) {
	s.elements = append(s.elements, number)
}

func (s *QueueList) Dequeue() int {
	first := s.elements[0]
	s.elements = s.elements[1:len(s.elements)]
	return first
}

func (s *QueueList) Peek() int {
	return s.elements[0]
}

func (s *QueueList) IsEmpty() bool {
	return len(s.elements) == 0
}
