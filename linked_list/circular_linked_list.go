package linkedlist

import "fmt"

// Traversal - access each element of the linked list
// Insertion - adds a new element to the linked list
// Deletion - removes the existing elements
// Search - find a node in the linked list
// Sort - sort the nodes of the linked list
// Insert At - insert at a position
// Delete At - delete at a position

type CircularLinkedList struct {
	Head   *SNode
	Length int
}

func (l *CircularLinkedList) InsertAt(position, value int) error {
	if l.isInvalidInsert(position) {
		return fmt.Errorf("invalid position")
	}
	new := &SNode{Value: value}
	defer l.increment()
	if l.isEmpty() {
		l.Head = new
	}

	if l.Length == position {
		new.Next = l.Head
	}

	curr := l.Head
	for index := 1; index < position-1; index++ {
		curr = curr.Next
	}
	new.Next = curr.Next
	curr.Next = new
	return nil
}

func (l *CircularLinkedList) DeleteAt(position int) error {
	if l.isInvalidDelete(position) {
		return fmt.Errorf("invalid delete")
	}

	defer l.decrement()
	var prev *SNode
	curr := l.Head
	for index := 1; index < position; index++ {
		if curr.Next != nil {
			prev = curr
			curr = curr.Next
		}
	}

	if prev == nil {
		l.Head = nil
		return nil
	}

	prev.Next = curr.Next
	return nil
}

func (l *CircularLinkedList) Exists(nodeValue int) (bool, int) {
	currNode := l.Head
	position := 0
	for currNode != nil {
		if currNode.Value == nodeValue {
			return true, position
		}
		currNode = currNode.Next
		position++
	}
	return false, -1
}

func (l *CircularLinkedList) Sort() {
	current := l.Head
	for current != nil {
		index := current.Next
		for index != nil {
			if current.Value > index.Value {
				temp := current.Value
				current.Value = index.Value
				index.Value = temp
			} else {
				index = index.Next
			}
		}
		current = current.Next
	}
}

func (l *CircularLinkedList) Reverse() *CircularLinkedList {
	reversed := &CircularLinkedList{}
	curr := l.Head
	for curr != nil {
		reversed.Head = &SNode{Value: curr.Value, Next: reversed.Head}
		curr = curr.Next
	}
	l.Head.Next = curr
	return reversed
}

func (l *CircularLinkedList) isEmpty() bool {
	return l.Head == nil
}

func (l *CircularLinkedList) increment() {
	l.Length += 1
}

func (l *CircularLinkedList) isInvalidDelete(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position)
}

func (l *CircularLinkedList) decrement() {
	l.Length = l.Length - 1
}

func (l *CircularLinkedList) isNegativeIndex(position int) bool {
	return position <= 0
}

func (l *CircularLinkedList) isGreaterThanLength(position int) bool {
	return position > l.Length
}

func (l *CircularLinkedList) isInvalidInsert(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position-1)
}
