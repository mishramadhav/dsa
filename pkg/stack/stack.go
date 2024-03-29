package stack

type stack struct {
	data []any
}

// New creates a new stack.
func New() *stack {
	return &stack{
		data: make([]any, 0),
	}
}

// NewWithCapacity creates a new stack with the specified capacity.
func NewWithCapacity(capacity int) *stack {
	return &stack{
		data: make([]any, 0, capacity),
	}
}

// Len returns the number of elements in the stack.
func (s *stack) Len() int {
	return len(s.data)
}

// Empty returns true if the stack is empty.
func (s *stack) Empty() bool {
	return s.Len() == 0
}

// Top returns the top element of the stack.
// If the stack is empty, it returns nil.
func (s *stack) Top() any {
	if s.Len() == 0 {
		return nil
	}

	return s.data[s.Len()-1]
}

// Push pushes a value onto the stack.
func (s *stack) Push(v any) {
	s.data = append(s.data, v)
}

// Pop pops a value from the stack.
// If the stack is empty, it returns nil.
func (s *stack) Pop() any {
	if s.Len() == 0 {
		return nil
	}

	v := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]

	return v
}
