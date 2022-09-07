package stack

import (
	ll "data-structures/linked_list"
)

type Stack struct {
	elements ll.SinglyLinkedList
}

func (s *Stack) New() {
	s.elements = ll.SinglyLinkedList{}
}

func (s *Stack) Push(number int) {
	s.elements.InsertAt(s.elements.Length, number)
}

func (s *Stack) Pop() int {
	top := s.Top()
	s.elements.DeleteAt(s.elements.Length)
	return top
}

func (s *Stack) Top() int {
	return s.elements.ElementAt(s.elements.Length + 1)
}

func (s *Stack) IsEmpty() bool {
	return s.elements.IsEmpty()
}

func (s *Stack) Size() int {
	return s.elements.Length
}
