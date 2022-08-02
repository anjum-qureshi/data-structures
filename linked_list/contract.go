package linkedlist

type SNode struct {
	Value int
	Next  *SNode
}

type DNode struct {
	Value int
	Next  *DNode
	Prev  *DNode
}
