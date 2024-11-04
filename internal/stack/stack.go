package stack

import (
	"errors"
	"sync"
)

var ErrEmptyStack = errors.New("empty stack")

type Stack[T any] struct {
	mu      *sync.RWMutex
	storage []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		mu:      &sync.RWMutex{},
		storage: make([]T, 0),
	}
}

func (s *Stack[T]) Top() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.storage) == 0 {
		var empty T
		return empty, ErrEmptyStack
	}

	last := s.storage[len(s.storage)-1]
	return last, nil
}

func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.storage)
}

func (s *Stack[T]) Pop() (T, error) {
	last, err := s.Top()
	if err != nil {
		return last, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage = s.storage[:len(s.storage)-1]
	return last, nil
}

func (s *Stack[T]) Push(obj T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage = append(s.storage, obj)
}
