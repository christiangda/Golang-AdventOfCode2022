package stack

import (
	"slices"
	"sync"
)

type Stack[T any] struct {
	mu  sync.RWMutex
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
	s.mu.Lock()
	defer s.mu.Unlock()

	s.e = append(s.e, e)
	s.top++
}

func (s *Stack[T]) Pop() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.IsEmpty() {
		var v T
		return v
	}
	v := s.e[(s.top - 1)]
	s.e = s.e[:(s.top - 1)]
	s.top--

	return v
}

func (s *Stack[T]) Reverse() {
	s.mu.Lock()
	defer s.mu.Unlock()
	slices.Reverse(s.e)
}
