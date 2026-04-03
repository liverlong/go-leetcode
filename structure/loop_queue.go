package structure

import "errors"

type LoopQueue[T any] struct {
	elements []T
	front    int
	tail     int
	size     int
}

func NewLoopQueue[T any](capacity int) *LoopQueue[T] {
	return &LoopQueue[T]{
		elements: make([]T, 0, capacity+1),
		front:    0,
		tail:     0,
		size:     0,
	}
}

func (q *LoopQueue[T]) IsEmpty() bool {
	return q.front == q.tail
}

func (q *LoopQueue[T]) GetCapacity() int {
	return len(q.elements) - 1
}

func (q *LoopQueue[T]) Enqueue(val T) {
	if (q.tail+1)%q.GetCapacity() == q.front {
		q.resize(q.GetCapacity() * 2)
	}
	q.elements[q.tail] = val
	q.tail = (q.tail + 1) % q.GetCapacity()
	q.size++
}

func (q *LoopQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}

	res := q.elements[q.front]
	var zero T
	q.elements[q.front] = zero
	q.front = (q.front + 1) % len(q.elements)
	q.size--

	if q.size == q.GetCapacity()/4 && q.GetCapacity()/2 != 0 {
		q.resize(q.GetCapacity() / 2)
	}

	return res, nil
}

func (q *LoopQueue[T]) GetFront() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}

	return q.elements[q.front], nil
}

func (q *LoopQueue[T]) resize(newCapacity int) {
	newData := make([]T, 0, newCapacity+1)
	for i := 0; i < q.size; i++ {
		newData[i] = q.elements[(i+q.front)%q.GetCapacity()]
	}
	q.elements = newData
	q.front = 0
	q.tail = q.size
}
