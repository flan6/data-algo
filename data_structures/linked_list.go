package datastructures

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	len  uint
}

type node[T any] struct {
	value T
	next  *node[T]
}

func (l LinkedList[T]) Length() uint {
	return l.len
}

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

func (l *LinkedList[T]) Get(index uint) T {
	if index == 0 {
		return l.head.value
	}

	if index == l.len {
		return l.tail.value
	}

	curr := l.head
	for i := uint(0); i < index ; i++ {
		curr = curr.next
	}

	return curr.value
}
