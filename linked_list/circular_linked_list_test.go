package linkedlist_test

import (
	ll "data-structures/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestCircularLinkedList struct {
	suite.Suite
}

func TestCircularLinkedList_InsertAt_AtStartWhenListIsEmpty(t *testing.T) {
	list := ll.CircularLinkedList{}
	list.InsertAt(1, 10)
	expected := createCircularLinkedList(10)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_AtStartWhenListHasOneElement(t *testing.T) {
	list := createCircularLinkedList(10)
	expected := createCircularLinkedList(10,20)
	list.InsertAt(1, 20)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_AtStartWhenListHasTwoElement(t *testing.T) {
	list := createCircularLinkedList(10, 30)
	expected := createCircularLinkedList(20, 10, 30)
	list.InsertAt(1, 20)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createCircularLinkedList(20, 10)
	expected := createCircularLinkedList(20, 10, 30)
	list.InsertAt(3, 30)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_AtLengthWhenPositionEqualToLength(t *testing.T) {
	list := createCircularLinkedList(20, 10)
	expected := createCircularLinkedList(20, 30, 10)
	list.InsertAt(2, 30)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	expected := createCircularLinkedList(10, 16, 20, 30)
	list.InsertAt(2, 16)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	expected := createCircularLinkedList(10, 20, 30)
	list.InsertAt(-1, 16)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_InsertAt_ErrorWhenPositionIsGreaterThanLengthPlus1(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	expected := createCircularLinkedList(10, 20, 30)
	list.InsertAt(5, 16)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_DeleteAt_WhenListIsEmpty(t *testing.T) {
	list := ll.CircularLinkedList{}
	err := list.DeleteAt(1)
	assert.Equal(t, err.Error(), "invalid position")
	assert.Exactly(t, list, ll.CircularLinkedList{})
}

func TestCircularLinkedList_DeleteAt_StartWhenListHasOneElement(t *testing.T) {
	list := createCircularLinkedList(10)
	list.DeleteAt(1)
	assert.Exactly(t, ll.CircularLinkedList{}, list)
}

func TestCircularLinkedList_DeleteAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := createCircularLinkedList(20, 10)
	list.DeleteAt(2)
	expected := createCircularLinkedList(20)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_DeleteAt_InBetweenWhenListHasElements(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	list.DeleteAt(2)
	expected := createCircularLinkedList(10, 30)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_DeleteAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	err := list.DeleteAt(-1)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.DeleteAt(0)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createCircularLinkedList(10, 20, 30)
	assert.Exactly(t, expected, list)
}

func TestCircularLinkedList_DeleteAt_ErrorWhenPositionIsGreaterThanLength(t *testing.T) {
	list := createCircularLinkedList(10, 20, 30)
	err := list.DeleteAt(4)
	assert.Equal(t, err.Error(), "invalid position")
	expected := createCircularLinkedList(10, 20, 30)
	assert.Exactly(t, expected, list)
}

func createCircularLinkedList(values ...int) ll.CircularLinkedList {
	list := ll.CircularLinkedList{Length: len(values)}
	if len(values) == 0 {
		return list
	}
	list.Tail = &ll.SNode{Value: values[0]}
	list.Tail.Next = list.Tail
	curr := list.Tail
	for _, value := range values[1:] {
		curr.Next = &ll.SNode{Value: value}
		curr = curr.Next
	}
	return list
}
