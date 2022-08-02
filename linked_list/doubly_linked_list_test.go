package linkedlist_test

import (
	ll "data-structures/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestDoublyLinkedList struct {
	suite.Suite
}

func TestDoublyLinkedList_InsertAt_AtStartWhenListIsEmpty(t *testing.T) {
	list := &ll.DoublyLinkedList{}
	list.InsertAt(1, 10)
	assert.Exactly(t, createDoublyLinkedList(10), list)
}

func TestDoublyLinkedList_InsertAt_AtStartWhenListHasOneElement(t *testing.T) {
	list := createDoublyLinkedList(10)
	list.InsertAt(1, 20)
	assert.Exactly(t, createDoublyLinkedList(20, 10), list)
}

func TestDoublyLinkedList_InsertAt_AtStartWhenListHasTwoElement(t *testing.T) {
	list := createDoublyLinkedList(10, 30)
	list.InsertAt(1, 20)
	expected := createDoublyLinkedList(20, 10, 30)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_InsertAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createDoublyLinkedList(20, 10)
	list.InsertAt(3, 30)
	expected := createDoublyLinkedList(20, 10, 30)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_InsertAt_AtLengthWhenPositionEqualToLength(t *testing.T) {
	list := createDoublyLinkedList(20, 10)
	list.InsertAt(2, 30)
	expected := createDoublyLinkedList(20, 30, 10)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_InsertAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createDoublyLinkedList(10, 20, 30, 40, 50, 60)
	list.InsertAt(4, 16)
	expected := createDoublyLinkedList(10, 20, 30, 16, 40, 50, 60)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_InsertAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createDoublyLinkedList(10, 20, 30)
	err := list.InsertAt(-1, 16)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.InsertAt(0, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createDoublyLinkedList(10, 20, 30)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_InsertAt_ErrorWhenPositionIsGreaterThanLengthPlus1(t *testing.T) {
	list := createDoublyLinkedList(10, 20)
	err := list.InsertAt(5, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createDoublyLinkedList(10, 20)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_DeleteAt_WhenListIsEmpty(t *testing.T) {
	list := &ll.DoublyLinkedList{}
	err := list.DeleteAt(1)
	assert.Equal(t, err.Error(), "invalid position")
	assert.Exactly(t, list, &ll.DoublyLinkedList{})
}

func TestDoublyLinkedList_DeleteAt_StartWhenListHasOneElement(t *testing.T) {
	list := createDoublyLinkedList(10)
	list.DeleteAt(1)
	assert.Exactly(t, &ll.DoublyLinkedList{}, list)
}

func TestDoublyLinkedList_DeleteAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createDoublyLinkedList(10, 20)
	list.DeleteAt(2)
	expected := createDoublyLinkedList(10)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_DeleteAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createDoublyLinkedList(1,2,3)
	list.DeleteAt(2)
	expected := createDoublyLinkedList(1,3)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_DeleteAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createDoublyLinkedList(1,2,3)
	err := list.DeleteAt(-1)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.InsertAt(0, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createDoublyLinkedList(1,2,3)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_DeleteAt_ErrorWhenPositionIsGreaterThanLength(t *testing.T) {
	list := createDoublyLinkedList(1,2,3)
	err := list.DeleteAt(4)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createDoublyLinkedList(1,2,3)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_Sort_ShouldSortLinkedListInAscendingOrder(t *testing.T) {
	list := createDoublyLinkedList(5, 10, -20, 30, 0, 5, 3)
	list.Sort()
	expected := createDoublyLinkedList(-20, 0, 3, 5, 5, 10, 30)
	assert.Exactly(t, expected, list)
}

func TestDoublyLinkedList_Reverse_ShouldReversedLinkedList(t *testing.T) {
	list := createDoublyLinkedList(5, 10, -20, 30, 0, 5, 3).Reverse()
	expected := createDoublyLinkedList(3, 5, 0, 30, -20, 10, 5)
	assert.Exactly(t, expected, list)
}


func createDoublyLinkedList(values ...int) *ll.DoublyLinkedList {
	list := ll.DoublyLinkedList{Length: len(values)}
	if len(values) == 0 {
		return &list
	}
	list.Head = &ll.DNode{Value: values[0]}
	curr := list.Head
	for _, value := range values[1:] {
		curr.Next = &ll.DNode{Value: value, Prev: curr}
		curr = curr.Next
	}
	return &list
}
