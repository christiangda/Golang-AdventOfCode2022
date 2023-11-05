package stack

type Stack[T any] struct {
	e   []T
	top int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		e:   make([]T, 0),
		top: 0,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack[T]) Push(e T) {
	s.e = append(s.e, e)
	s.top++
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var v T
		return v
	}
	v := s.e[(s.top - 1)]

	return v
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var v T
		return v
	}
	v := s.e[(s.top - 1)]
	s.e = s.e[:(s.top - 1)]
	s.top--

	return v
}

func (s *Stack[T]) Size() int {
	return len(s.e)
}
