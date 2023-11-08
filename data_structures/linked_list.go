package data_structures

import "errors"

type LinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
	len  uint
}

type node[T comparable] struct {
	value T
	next  *node[T]
}

var ErrIndexOutOfRange = errors.New("index out of range")

func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{}
}

// Length returns the current size of the list
func (l LinkedList[T]) Length() uint {
	return l.len
}

// Append inserts item at the end of the list
func (l *LinkedList[T]) Append(item T) {
	n := &node[T]{value: item}
	l.len++

	if l.tail == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
}

// Preappend inserts item in the begining of the list
func (l *LinkedList[T]) Preappend(item T) {
	n := &node[T]{value: item}
	l.len++

	if l.head == nil {
		l.head = n
		return
	}

	n.next = l.head
	l.head = n
}

// Get gets the value at index
func (l *LinkedList[T]) Get(index uint) (T, error) {
	if index > l.len {
		return *new(T), ErrIndexOutOfRange
	}

	return l.get(index).value, nil
}

// IndexOf finds the index of value, returns -1 if not found
func (l LinkedList[T]) IndexOf(value T) int {
	idx := 0
	n := l.walk(l.head, value, &idx)
	if n != nil {
		return idx
	}

	return -1
}

// Set sets the new value at index
// returns ErrIndexOutOfRange in case index does not exist in the list
func (l *LinkedList[T]) Set(index uint, value T) error {
	if index > l.len {
		return ErrIndexOutOfRange
	}

	l.get(index).value = value

	return nil
}

func (l LinkedList[T]) walk(n *node[T], value T, index *int) *node[T] {
	if n == nil {
		return nil
	}

	if n.value == value {
		return n
	}

	if index != nil {
		*index++
	}

	return l.walk(n.next, value, index)
}

func (l LinkedList[T]) get(index uint) *node[T] {
	if index == 0 {
		return l.head
	}

	if index == l.len {
		return l.tail
	}

	curr := l.head
	for i := uint(0); i < index; i++ {
		curr = curr.next
	}

	return curr
}
