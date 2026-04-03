package structure

import (
	"errors"
)

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

type LinkedList[T any] struct {
	dummyHead *listNode[T]
	tail      *listNode[T]
	size      int
	equals    func(a, b T) bool
}

func NewLinkedList[T any](equals func(a, b T) bool) *LinkedList[T] {
	var zero T
	return &LinkedList[T]{
		dummyHead: &listNode[T]{val: zero},
		size:      0,
		equals:    equals,
	}
}

func (l *LinkedList[T]) GetSize() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Contains linked list contains element
func (l *LinkedList[T]) Contains(t T) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if l.equals(cur.val, t) {
			return true
		}
		cur = cur.next
	}

	return false
}

func (l *LinkedList[T]) Enqueue(t T) {
	// empty linked list
	if l.tail == nil {
		l.tail = &listNode[T]{val: t}
		l.dummyHead.next = l.tail
	} else {
		l.tail.next = &listNode[T]{val: t}
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedList[T]) Dequeue() (T, error) {
	if l.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}

	retNode := l.dummyHead.next
	l.dummyHead.next = retNode.next
	retNode.next = nil

	// only one node
	if l.dummyHead.next == nil {
		l.tail = nil
	}

	l.size--
	return retNode.val, nil
}
