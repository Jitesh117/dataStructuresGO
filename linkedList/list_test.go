package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(30)
	if got, want := ll.Display(), []int{10, 20, 30}; !equal(got, want) {
		t.Errorf("Display() = %v; want %v", got, want)
	}

	ll.InsertAtBeginning(5)
	if got, want := ll.Display(), []int{5, 10, 20, 30}; !equal(got, want) {
		t.Errorf("Display() = %v; want %v", got, want)
	}

	if err := ll.Delete(20); err != nil {
		t.Errorf("Delete(20) failed: %v", err)
	}
	if got, want := ll.Display(), []int{5, 10, 30}; !equal(got, want) {
		t.Errorf("Display() = %v; want %v", got, want)
	}

	if err := ll.Delete(100); err == nil {
		t.Errorf("Delete(100) succeeded; want error")
	}

	if err := ll.Delete(5); err != nil {
		t.Errorf("Delete(5) failed: %v", err)
	}
	if got, want := ll.Display(), []int{10, 30}; !equal(got, want) {
		t.Errorf("Display() = %v; want %v", got, want)
	}

	if err := ll.Delete(10); err != nil {
		t.Errorf("Delete(10) failed: %v", err)
	}
	if err := ll.Delete(30); err != nil {
		t.Errorf("Delete(30) failed: %v", err)
	}
	if got, want := ll.Display(), []int{}; !equal(got, want) {
		t.Errorf("Display() = %v; want %v", got, want)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
