package stack

import (
	ll "data-structures/linked_list"
)

type Queue struct {
	elements ll.SinglyLinkedList
}

func (s *Queue) New() {
	s.elements = ll.SinglyLinkedList{}
}

func (s *Queue) Enqueue(number int) {
	s.elements.InsertAt(number,s.elements.Length+1)
}

func (s *Queue) Dequeue() int {
	first := s.elements.ElementAt(1)
	s.elements.DeleteAt(1)
	return first
}

func (s *Queue) Peek() int {
	return s.elements.ElementAt(1)
}

func (s *Queue) IsEmpty() bool {
	return s.elements.IsEmpty()
}

func (s *Queue) Size() int {
	return s.elements.Length
}
