package linkedlist

import (
	"fmt"
)

// Traversal - access each element of the linked list
// Insertion - adds a new element to the linked list
// Deletion - removes the existing elements
// Search - find a node in the linked list
// Sort - sort the nodes of the linked list
// Insert At - insert at a position
// Delete At - delete at a position

type SinglyLinkedList struct {
	Head   *SNode
	Length int
}

func (l *SinglyLinkedList) InsertAt(position, value int) error {
	if l.isInvalidInsert(position) {
		return fmt.Errorf("invalid position")
	}
	defer l.increment()
	new := &SNode{Value: value, Next: l.Head}
	if l.addAtHead(position) {
		l.Head = new
		return nil
	}

	curr := l.Head
	for index := 2; index < position; index++ {
		if curr.Next != nil {
			curr = curr.Next
		}
	}
	new.Next = curr.Next
	curr.Next = new
	return nil
}

func (l *SinglyLinkedList) DeleteAt(position int) error {
	if l.isInvalidDelete(position) {
		return fmt.Errorf("invalid position")
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

func (l *SinglyLinkedList) PrintAll() {
	currNode := l.Head
	if l.isEmpty() {
		fmt.Println("empty linked list")
		return
	}
	for currNode != nil {
		fmt.Printf("%d ", currNode.Value)
		currNode = currNode.Next
	}
	fmt.Println("")
}

func (l *SinglyLinkedList) Exists(nodeValue int) (bool, int) {
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

func (l *SinglyLinkedList) Sort() {
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

func (l *SinglyLinkedList) Reverse() *SinglyLinkedList {
	reversed := &SinglyLinkedList{}
	curr := l.Head
	for curr != nil {
		reversed.Head = &SNode{Value: curr.Value, Next: reversed.Head}
		curr = curr.Next
	}
	return reversed
}

func (l *SinglyLinkedList) isEmpty() bool {
	return l.Head == nil
}

func (l *SinglyLinkedList) addAtHead(position int) bool {
	return position == 1 || l.isEmpty()
}

func (l *SinglyLinkedList) isNegativeIndex(position int) bool {
	return position <= 0
}

func (l *SinglyLinkedList) isGreaterThanLength(position int) bool {
	return position > l.Length
}

func (l *SinglyLinkedList) isInvalidInsert(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position-1)
}

func (l *SinglyLinkedList) isInvalidDelete(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position)
}

func (l *SinglyLinkedList) increment() {
	l.Length += 1
}

func (l *SinglyLinkedList) decrement() {
	l.Length -= 1
}
