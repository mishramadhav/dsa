package queue

import (
	"errors"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	// slice to store the Elements
	Elements []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element to the back of the queue.
func (q *Queue[T]) Push(element T) {
	q.Elements = append(q.Elements, element)
}

// Pop removes and returns the element at the front of the queue.
func (q *Queue[T]) Pop() (T, error) {
	if len(q.Elements) == 0 {
		var zeroValue T
		return zeroValue, ErrEmptyQueue
	}

	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element, nil
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return len(q.Elements) == 0
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.Elements)
}
