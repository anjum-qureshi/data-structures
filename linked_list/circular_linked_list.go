package linkedlist

// Traversal - access each element of the linked list
// Insertion - adds a new element to the linked list
// Deletion - removes the existing elements
// Search - find a node in the linked list
// Sort - sort the nodes of the linked list
// Insert At - insert at a position
// Delete At - delete at a position

type CircularLinkedList struct {
	Tail   *SNode
	Length int
}

func (l *CircularLinkedList) isEmpty() bool {
	return l.Tail == nil
}

func (l *CircularLinkedList) increment() {
	l.Length += 1
}

func (l *CircularLinkedList) InsertAt(position, value int) error {
	new := &SNode{Value: value, Next: l.Tail}
	defer l.increment()
	if l.isEmpty() {
		l.Tail = new
		l.Tail.Next = new
	}
	curr := l.Tail
	for index := 1; index < position; index++ {
			curr = curr.Next
	}
	curr.Next = new
	return nil
}

func (l *CircularLinkedList) DeleteAt(position int) error {
	new := &SNode{Next: l.Tail}
	defer l.increment()
	if l.isEmpty() {
		l.Tail = new
		l.Tail.Next = new
	}
	curr := l.Tail
	for index := 1; index < position; index++ {
		if curr != nil {
			curr = curr.Next
		}
	}
	new.Next = curr.Next
	curr.Next = new
	return nil
}
