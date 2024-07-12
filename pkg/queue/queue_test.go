package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	if !q.Empty() {
		t.Errorf("expected queue to be empty")
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Empty() {
		t.Errorf("expected queue to not be empty")
	}
	if q.Len() != 3 {
		t.Errorf("expected queue length to be 3")
	}
	for i := 0; i < 3; i++ {
		if element, err := q.Pop(); err != nil || element != i+1 {
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			t.Errorf("expected %d, got %v", i+1, element)
		}
	}
	if !q.Empty() {
		t.Errorf("expected queue to be empty")
	}

	_, err := q.Pop()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err != ErrEmptyQueue {
		t.Errorf("expected %v, got %v", ErrEmptyQueue, err)
	}
}
