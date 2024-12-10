package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	// Test IsEmpty on a new stack
	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty initially")
	}

	// Test Push
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.IsEmpty() {
		t.Errorf("expected stack to be non-empty after pushes")
	}

	// Test Size
	if stack.Size != 3 {
		t.Errorf("expected size 3, got %d", stack.Size)
	}

	// Test Peek
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if top != 30 {
		t.Errorf("expected top element 30, got %d", top)
	}

	// Test Pop
	popped, err := stack.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if popped != 30 {
		t.Errorf("expected popped element 30, got %d", popped)
	}

	// Test Size after pop
	if stack.Size != 2 {
		t.Errorf("expected size 2 after pop, got %d", stack.Size)
	}

	// Test Peek after pop
	top, err = stack.Peek()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if top != 20 {
		t.Errorf("expected top element 20 after pop, got %d", top)
	}

	// Test Pop until empty
	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty after popping all elements")
	}

	// Test Pop on empty stack
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("expected error when popping from an empty stack")
	}

	// Test Peek on empty stack
	_, err = stack.Peek()
	if err == nil {
		t.Errorf("expected error when peeking on an empty stack")
	}
}
