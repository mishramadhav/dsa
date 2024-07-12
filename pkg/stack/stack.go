package stack

import "fmt"

var ErrEmptyStack = fmt.Errorf("stack is empty")

type Stack[T any] struct {
	data []T
}

// New creates a new stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// NewWithCapacity creates a new stack with the specified capacity.
func NewWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, capacity),
	}
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Empty returns true if the stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Top returns the top element of the stack.
// If the stack is empty, it returns error.
func (s *Stack[T]) Top() (T, error) {
	if s.Len() == 0 {
		var zeroValue T
		return zeroValue, ErrEmptyStack
	}

	return s.data[s.Len()-1], nil
}

// Push pushes a value onto the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop pops a value from the stack.
// If the stack is empty, it returns error.
func (s *Stack[T]) Pop() (T, error) {
	if s.Len() == 0 {
		var zeroValue T
		return zeroValue, ErrEmptyStack
	}

	v := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]

	return v, nil
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}
