package linkedlist_test

import (
	ll "data-structures/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSinglyLinkedList struct {
	suite.Suite
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListIsEmpty(t *testing.T) {
	list := ll.SinglyLinkedList{}
	list.InsertAt(1, 10)
	expected := createSinglyLinkedList(10)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListHasOneElement(t *testing.T) {
	list := createSinglyLinkedList(10)
	expected := createSinglyLinkedList(20, 10)
	list.InsertAt(1, 20)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListHasTwoElement(t *testing.T) {
	list := createSinglyLinkedList(10, 30)
	expected := createSinglyLinkedList(20, 10, 30)
	list.InsertAt(1, 20)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createSinglyLinkedList(20, 10)
	expected := createSinglyLinkedList(20, 10, 30)
	list.InsertAt(3, 30)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtLengthWhenPositionEqualToLength(t *testing.T) {
	list := createSinglyLinkedList(20, 10)
	expected := createSinglyLinkedList(20, 30, 10)
	list.InsertAt(2, 30)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	expected := createSinglyLinkedList(10, 16, 20, 30)
	list.InsertAt(2, 16)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	expected := createSinglyLinkedList(10, 20, 30)
	list.InsertAt(-1, 16)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_ErrorWhenPositionIsGreaterThanLengthPlus1(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	expected := createSinglyLinkedList(10, 20, 30)
	list.InsertAt(5, 16)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_WhenListIsEmpty(t *testing.T) {
	list := ll.SinglyLinkedList{}
	err := list.DeleteAt(1)
	assert.Equal(t, err.Error(), "invalid position")
	assert.Exactly(t, list, ll.SinglyLinkedList{})
}

func TestSinglyLinkedList_DeleteAt_StartWhenListHasOneElement(t *testing.T) {
	list := createSinglyLinkedList(10)
	list.DeleteAt(1)
	assert.Exactly(t, ll.SinglyLinkedList{}, list)
}

func TestSinglyLinkedList_DeleteAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createSinglyLinkedList(20, 10)
	list.DeleteAt(2)
	expected := createSinglyLinkedList(20)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	list.DeleteAt(2)
	expected := createSinglyLinkedList(10, 30)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	err := list.DeleteAt(-1)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.DeleteAt(0)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createSinglyLinkedList(10, 20, 30)
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_ErrorWhenPositionIsGreaterThanLength(t *testing.T) {
	list := createSinglyLinkedList(10, 20, 30)
	err := list.DeleteAt(4)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createSinglyLinkedList(10, 20, 30)
	assert.Exactly(t, expected, list)
}

func createSinglyLinkedList(values ...int) ll.SinglyLinkedList {
	list := ll.SinglyLinkedList{Length: len(values)}
	if len(values) == 0 {
		return list
	}
	list.Head = &ll.SNode{Value: values[0]}
	curr := list.Head
	for _, value := range values[1:] {
		curr.Next = &ll.SNode{Value: value}
		curr = curr.Next
	}
	return list
}
