package linkedlist

import "fmt"

// Traversal - access each element of the linked list
// Insertion - adds a new element to the linked list
// Deletion - removes the existing elements
// Search - find a node in the linked list
// Sort - sort the nodes of the linked list
// Insert At - insert at a position
// Delete At - delete at a position

type DoublyLinkedList struct {
	Head   *DNode
	Length int
}

func (l *DoublyLinkedList) InsertAt(position, value int) error {
	if l.isInvalidInsert(position) {
		return fmt.Errorf("invalid position")
	}
	defer l.increment()
	curr := l.Head
	for index := 2; index < position; index++ {
		if curr != nil {
			curr = curr.Next
		}
	}
	new := &DNode{Value: value}
	if position == 1 {
		l.Head = new
		new.Next = curr
		if curr != nil {
			curr.Prev = new
		}
		return nil
	}
	new.Prev = curr
	new.Next = curr.Next
	if curr.Next != nil {
		curr.Next.Prev = new
	}
	curr.Next = new

	return nil
}

func (l *DoublyLinkedList) Print() {
	curr1 := l.Head
	for curr1 != nil {
		var p, c, n string
		if curr1.Prev != nil {
			p = fmt.Sprintf("%d", curr1.Prev.Value)
		}
		c = fmt.Sprintf("%d", curr1.Value)
		if curr1.Next != nil {
			n = fmt.Sprintf("%d", curr1.Next.Value)
		}
		fmt.Printf("%s <--- %s ---> %s\n", p, c, n)
		curr1 = curr1.Next
	}
}

func (l *DoublyLinkedList) DeleteAt(position int) error {
	if l.isInvalidDelete(position) {
		return fmt.Errorf("invalid position")
	}

	defer l.decrement()
	curr := l.Head
	for index := 1; index < position; index++ {
		if curr.Next != nil {
			curr = curr.Next
		}
	}
	if curr.Prev == nil {
		l.Head = nil
		return nil
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}
	curr.Prev.Next = curr.Next
	return nil
}

func (l *DoublyLinkedList) PrintAll() {
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

func (l *DoublyLinkedList) Exists(nodeValue int) (bool, int) {
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

func (l *DoublyLinkedList) Sort() {
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

func (l *DoublyLinkedList) Reverse() *DoublyLinkedList {
	reversed := &DoublyLinkedList{}
	curr := l.Head
	for curr != nil {
		new := &DNode{Value: curr.Value, Next: reversed.Head}
		if reversed.Head != nil {
			reversed.Head.Prev = new
		}
		reversed.Head = new
		curr = curr.Next
	}
	reversed.Length = l.Length
	return reversed
}

func (l *DoublyLinkedList) isEmpty() bool {
	return l.Head == nil
}

func (l *DoublyLinkedList) addAtHead(position int) bool {
	return position == 1 || l.isEmpty()
}

func (l *DoublyLinkedList) isNegativeIndex(position int) bool {
	return position <= 0
}

func (l *DoublyLinkedList) isGreaterThanLength(position int) bool {
	return position > l.Length
}

func (l *DoublyLinkedList) isInvalidInsert(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position-1)
}

func (l *DoublyLinkedList) isInvalidDelete(position int) bool {
	return l.isNegativeIndex(position) || l.isGreaterThanLength(position)
}

func (l *DoublyLinkedList) increment() {
	l.Length += 1
}

func (l *DoublyLinkedList) decrement() {
	l.Length -= 1
}
