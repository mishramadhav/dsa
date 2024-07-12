package queue

import "errors"

var ErrEmptyDequeue = errors.New("dequeue is empty")

type Dequeue[T any] struct {
	// slice to store elements
	Element []T
}

// NewDequeue returns a new Dequeue.
func NewDequeue[T any]() *Dequeue[T] {
	return &Dequeue[T]{}
}

// Empty returns true if the Dequeue is empty.
func (d *Dequeue[T]) Empty() bool {
	return d.Len() == 0
}

// Len returns the number of elements in the Dequeue.
func (d *Dequeue[T]) Len() int {
	return len(d.Element)
}

// PushFront adds an element to the front of the Dequeue.
func (d *Dequeue[T]) PushFront(e T) {
	d.Element = append([]T{e}, d.Element...)
}

// PushBack adds an element to the back of the Dequeue.
func (d *Dequeue[T]) PushBack(e T) {
	d.Element = append(d.Element, e)
}

// PopFront removes and returns the element from the front of the Dequeue.
func (d *Dequeue[T]) PopFront() (T, error) {
	if d.Empty() {
		var zeroValue T
		return zeroValue, ErrEmptyDequeue
	}
	e := d.Element[0]
	d.Element = d.Element[1:]
	return e, nil
}

// PopBack removes and returns the element from the back of the Dequeue.
func (d *Dequeue[T]) PopBack() (T, error) {
	if d.Empty() {
		var zeroValue T
		return zeroValue, ErrEmptyDequeue
	}
	e := d.Element[d.Len()-1]
	d.Element = d.Element[:d.Len()-1]
	return e, nil
}

// Front returns the element at the front of the Dequeue without removing it.
func (d *Dequeue[T]) Front() (T, error) {
	if d.Empty() {
		var zeroValue T
		return zeroValue, ErrEmptyDequeue
	}
	return d.Element[0], nil
}

// Back returns the element at the back of the Dequeue without removing it.
func (d *Dequeue[T]) Back() (T, error) {
	if d.Empty() {
		var zeroValue T
		return zeroValue, ErrEmptyDequeue
	}
	return d.Element[d.Len()-1], nil
}
