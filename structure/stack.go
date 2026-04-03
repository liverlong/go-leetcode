package structure

import "errors"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) GetSize() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	idx := s.GetSize() - 1
	val := s.elements[idx]

	// 切除最后一个元素
	s.elements = s.elements[:idx]
	return val, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.elements[s.GetSize()-1], nil
}
