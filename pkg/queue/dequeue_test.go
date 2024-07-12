package queue

import "testing"

func TestDequeue(t *testing.T) {
	dq := NewDequeue[int]()
	if !dq.Empty() {
		t.Errorf("expected empty dequeue")
	}

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushBack(3)
	dq.PushBack(4)

	if dq.Len() != 4 {
		t.Errorf("expected dequeue length 4, got %d", dq.Len())
	}
	v, err := dq.Front()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 2 {
		t.Errorf("expected front element 2, got %d", v)
	}

	v, err = dq.Back()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 4 {
		t.Errorf("expected back element 4, got %d", v)
	}

	v, err = dq.PopFront()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 2 {
		t.Errorf("expected popped element 2, got %d", v)
	}

	v, err = dq.PopBack()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 4 {
		t.Errorf("expected popped element 4, got %d", v)
	}

	if dq.Len() != 2 {
		t.Errorf("expected dequeue length 2, got %d", dq.Len())
	}

	v, err = dq.PopFront()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 1 {
		t.Errorf("expected popped element 1, got %d", v)
	}

	v, err = dq.PopBack()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v != 3 {
		t.Errorf("expected popped element 3, got %d", v)
	}

	if !dq.Empty() {
		t.Errorf("expected empty dequeue")
	}

	_, err = dq.PopFront()
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyDequeue {
		t.Errorf("expected error ErrEmptyDequeue, got %v", err)
	}

	_, err = dq.PopBack()
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyDequeue {
		t.Errorf("expected error ErrEmptyDequeue, got %v", err)
	}

	_, err = dq.Front()
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyDequeue {
		t.Errorf("expected error ErrEmptyDequeue, got %v", err)
	}

	_, err = dq.Back()
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyDequeue {
		t.Errorf("expected error ErrEmptyDequeue, got %v", err)
	}
}
