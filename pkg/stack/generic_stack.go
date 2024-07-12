package stack

type Stack[T any] struct {
	data []T
}

// New creates a new generic stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// NewWithCapacity creates a new generic stack with the specified capacity.
func NewWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, capacity),
	}
}

// Len returns the number of elements in the generic stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Empty returns true if the generic stack is empty.
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Top returns the top element of the generic stack.
// If the generic stack is empty, it returns zero value.
func (s *Stack[T]) Top() T {
	if s.Len() == 0 {
		var zero T
		return zero
	}

	return s.data[s.Len()-1]
}

// Push pushes a value onto the generic stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop pops a value from the generic stack.
// If the generic stack is empty, it returns zero value.
func (s *Stack[T]) Pop() T {
	if s.Len() == 0 {
		var zero T
		return zero
	}

	v := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]

	return v
}

// Clear removes all elements from the generic stack.
func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}
