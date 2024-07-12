package stack_test

import (
	"testing"

	"github.com/mishramadhav/dsa/pkg/stack"
)

func TestStack(t *testing.T) {
	s := stack.New[int]()

	if s.Len() != 0 {
		t.Errorf("stack length - expected: %d, got: %d", 0, s.Len())
	}
	if !s.Empty() {
		t.Errorf("expected stack to be empty")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("stack length - expected: %d, got: %d", 1, s.Len())
	}
	if s.Empty() {
		t.Errorf("expected stack to not be empty")
	}
	if s.Top() != 1 {
		t.Errorf("top element - expected: %d, got: %d", 1, s.Top())
	}

	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("stack length - expected: %d, got: %d", 2, s.Len())
	}
	if s.Empty() {
		t.Errorf("expected stack to not be empty")
	}
	if s.Top() != 2 {
		t.Errorf("top element - expected: %d, got: %d", 2, s.Top())
	}

	if v := s.Pop(); v != 2 {
		t.Errorf("popped element - expected: %d, got: %d", 2, v)
	}
	if s.Len() != 1 {
		t.Errorf("stack length - expected: %d, got: %d", 1, s.Len())
	}

	if v := s.Pop(); v != 1 {
		t.Errorf("popped element - expected: %d, got: %d", 1, v)
	}
	if s.Len() != 0 {
		t.Errorf("stack length - expected: %d, got: %d", 0, s.Len())
	}

	s.Push(3)
	s.Push(4)
	if s.Len() != 2 {
		t.Errorf("stack length - expected: %d, got: %d", 2, s.Len())
	}
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("stack length - expected: %d, got: %d", 0, s.Len())
	}

	s = stack.NewWithCapacity[int](2)
	if s.Len() != 0 {
		t.Errorf("stack length - expected: %d, got: %d", 0, s.Len())
	}
	if !s.Empty() {
		t.Errorf("expected stack to be empty")
	}
	if s.Top() != 0 {
		t.Errorf("top element - expected: %d, got: %d", 0, s.Top())
	}
	if v := s.Pop(); v != 0 {
		t.Errorf("popped element - expected: %d, got: %d", 0, v)
	}
	s.Push(5)
	s.Push(6)
	if s.Len() != 2 {
		t.Errorf("stack length - expected: %d, got: %d", 2, s.Len())
	}
	if s.Top() != 6 {
		t.Errorf("top element - expected: %d, got: %d", 6, s.Top())
	}
	if v := s.Pop(); v != 6 {
		t.Errorf("popped element - expected: %d, got: %d", 6, v)
	}
}

func BenchmarkStackPush(b *testing.B) {
	s := stack.New[int]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStackPushWithCapacity(b *testing.B) {
	s := stack.NewWithCapacity[int](b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
