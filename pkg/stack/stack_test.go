package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mishramadhav/dsa/pkg/stack"
)

func TestStack(t *testing.T) {
	s := stack.New()

	assert := assert.New(t)

	assert.Equal(0, s.Len(), "stack length should be 0")
	assert.True(s.Empty(), "stack should be empty")

	s.Push(1)

	assert.Equal(1, s.Len(), "stack length should be 1")
	assert.False(s.Empty(), "stack should not be empty")
	assert.Equal(1, s.Top(), "top element should be 1")

	s.Push(2)

	assert.Equal(2, s.Len(), "stack length should be 2")
	assert.False(s.Empty(), "stack should not be empty")
	assert.Equal(2, s.Top(), "top element should be 2")
	assert.Equal(2, s.Pop(), "popped element should be 2")
	assert.Equal(1, s.Len(), "stack length should be 1")
	assert.False(s.Empty(), "stack should not be empty")
	assert.Equal(1, s.Top(), "top element should be 1")
	assert.Equal(1, s.Pop(), "popped element should be 1")
	assert.Equal(0, s.Len(), "stack length should be 0")
	assert.True(s.Empty(), "stack should be empty")
	assert.Nil(s.Top(), "top element should be nil")
	assert.Nil(s.Pop(), "popped element should be nil")

	stackWithCapacity := stack.NewWithCapacity(2)

	assert.Equal(0, stackWithCapacity.Len(), "stack length should be 0")
	assert.True(stackWithCapacity.Empty(), "stack should be empty")

	stackWithCapacity.Push(1)

	assert.Equal(1, stackWithCapacity.Len(), "stack length should be 1")
	assert.False(stackWithCapacity.Empty(), "stack should not be empty")
	assert.Equal(1, stackWithCapacity.Top(), "top element should be 1")

	stackWithCapacity.Push(2)

	assert.Equal(2, stackWithCapacity.Len(), "stack length should be 2")
	assert.False(stackWithCapacity.Empty(), "stack should not be empty")
	assert.Equal(2, stackWithCapacity.Top(), "top element should be 2")
	assert.Equal(2, stackWithCapacity.Pop(), "popped element should be 2")
	assert.Equal(1, stackWithCapacity.Len(), "stack length should be 1")
	assert.False(stackWithCapacity.Empty(), "stack should not be empty")
	assert.Equal(1, stackWithCapacity.Top(), "top element should be 1")
	assert.Equal(1, stackWithCapacity.Pop(), "popped element should be 1")
	assert.Equal(0, stackWithCapacity.Len(), "stack length should be 0")
	assert.True(stackWithCapacity.Empty(), "stack should be empty")
	assert.Nil(stackWithCapacity.Top(), "top element should be nil")
	assert.Nil(stackWithCapacity.Pop(), "popped element should be nil")
}

func BenchmarkStackPush(b *testing.B) {
	s := stack.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStackPushWithCapacity(b *testing.B) {
	s := stack.NewWithCapacity(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
