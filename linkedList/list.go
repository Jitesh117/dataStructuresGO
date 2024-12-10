package linkedlist

import "errors"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) Delete(data int) error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.Size--
		return nil
	}

	current := ll.Head
	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	if current.Next == nil {
		return errors.New("value not found in the list")
	}

	current.Next = current.Next.Next
	ll.Size--
	return nil
}

func (ll *LinkedList) Display() []int {
	var result []int
	current := ll.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}
