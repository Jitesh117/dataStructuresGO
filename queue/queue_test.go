package queue

import (
	"testing"
)

func TestNewQueue_IsEmpty(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
}

func TestQueue_PushAndSize(t *testing.T) {
	q := NewQueue()

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", q.Size())
	}
}

func TestQueue_IsEmptyAfterPush(t *testing.T) {
	q := NewQueue()

	q.Push(10)

	if q.IsEmpty() {
		t.Errorf("Expected queue to be non-empty")
	}
}

func TestQueue_Front(t *testing.T) {
	q := NewQueue()

	q.Push(10)
	q.Push(20)

	front, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error from Front: %v", err)
	}
	if front != 10 {
		t.Errorf("Expected Front to return 10, got %d", front)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()

	q.Push(10)
	q.Push(20)
	q.Push(30)

	popped, err := q.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop: %v", err)
	}
	if popped != 10 {
		t.Errorf("Expected Pop to return 10, got %d", popped)
	}

	if q.Size() != 2 {
		t.Errorf("Expected queue size to be 2, got %d", q.Size())
	}
}

func TestQueue_Reset(t *testing.T) {
	q := NewQueue()

	q.Push(10)
	q.Push(20)

	q.Reset()

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after Reset")
	}
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0 after Reset, got %d", q.Size())
	}
}

func TestQueue_PopOnEmptyQueue(t *testing.T) {
	q := NewQueue()

	_, err := q.Pop()
	if err == nil {
		t.Errorf("Expected error from Pop on empty queue")
	}
}

func TestQueue_FrontOnEmptyQueue(t *testing.T) {
	q := NewQueue()

	_, err := q.Front()
	if err == nil {
		t.Errorf("Expected error from Front on empty queue")
	}
}
