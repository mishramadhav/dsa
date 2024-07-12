package queue

import "fmt"

var ErrEmptyQueue = fmt.Errorf("queue is empty")

type Queue[T any] struct {
	// slice to store the elements
	elements []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element to the back of the queue.
func (q *Queue[T]) Push(element T) {
	q.elements = append(q.elements, element)
}

// Pop removes and returns the element at the front of the queue.
func (q *Queue[T]) Pop() (T, error) {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue, ErrEmptyQueue
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool {
	return len(q.elements) == 0
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(q.elements)
}
