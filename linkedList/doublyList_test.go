package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("Test AddToHead", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToHead(10)
		dll.AddToHead(20)
		dll.AddToHead(30)

		got := dll.Traverse()
		want := []int{30, 20, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddToHead failed. Got %v, want %v", got, want)
		}
	})

	t.Run("Test AddToTail", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToTail(10)
		dll.AddToTail(20)
		dll.AddToTail(30)

		got := dll.Traverse()
		want := []int{10, 20, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddToTail failed. Got %v, want %v", got, want)
		}
	})

	t.Run("Test RemoveFromHead", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToHead(10)
		dll.AddToHead(20)

		val, ok := dll.RemoveFromHead()
		if !ok || val != 20 {
			t.Errorf("RemoveFromHead failed. Got %d, want 20", val)
		}

		val, ok = dll.RemoveFromHead()
		if !ok || val != 10 {
			t.Errorf("RemoveFromHead failed. Got %d, want 10", val)
		}

		if _, ok := dll.RemoveFromHead(); ok {
			t.Error("RemoveFromHead should fail on empty list")
		}
	})

	t.Run("Test RemoveFromTail", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToTail(10)
		dll.AddToTail(20)

		val, ok := dll.RemoveFromTail()
		if !ok || val != 20 {
			t.Errorf("RemoveFromTail failed. Got %d, want 20", val)
		}

		val, ok = dll.RemoveFromTail()
		if !ok || val != 10 {
			t.Errorf("RemoveFromTail failed. Got %d, want 10", val)
		}

		if _, ok := dll.RemoveFromTail(); ok {
			t.Error("RemoveFromTail should fail on empty list")
		}
	})

	t.Run("Test Find", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToTail(10)
		dll.AddToTail(20)
		dll.AddToTail(30)

		node := dll.Find(20)
		if node == nil || node.Val != 20 {
			t.Errorf("Find failed. Got %v, want 20", node)
		}

		if dll.Find(40) != nil {
			t.Error("Find should return nil for non-existent value")
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToTail(10)
		dll.AddToTail(20)
		dll.AddToTail(30)

		if !dll.Delete(20) {
			t.Error("Delete failed for existing value 20")
		}

		got := dll.Traverse()
		want := []int{10, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Delete failed. Got %v, want %v", got, want)
		}

		if dll.Delete(40) {
			t.Error("Delete should return false for non-existent value")
		}
	})

	t.Run("Test Traverse", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.AddToHead(10)
		dll.AddToTail(20)
		dll.AddToTail(30)

		got := dll.Traverse()
		want := []int{10, 20, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Traverse failed. Got %v, want %v", got, want)
		}
	})

	t.Run("Test GetSize", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		if dll.GetSize() != 0 {
			t.Errorf("GetSize failed on empty list. Got %d, want 0", dll.GetSize())
		}

		dll.AddToHead(10)
		dll.AddToTail(20)

		if dll.GetSize() != 2 {
			t.Errorf("GetSize failed. Got %d, want 2", dll.GetSize())
		}
	})
}
