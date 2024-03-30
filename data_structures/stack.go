package data_structures

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		NewLinkedList[T](),
	}
}

type Stack[T any] struct {
	data DoubleLinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.data.Append(value)
}

func (s *Stack[T]) Pop() T {
	res := s.data.tail.value

	s.data.Remove(s.data.Length())

	return res
}
