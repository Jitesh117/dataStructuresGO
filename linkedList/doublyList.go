package linkedlist

type DoubleNode struct {
	Val  int
	Prev *DoubleNode
	Next *DoubleNode
}

type DoublyLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) AddToHead(val int) {
	newNode := &DoubleNode{Val: val}
	if dll.Head == nil { // If the list is empty
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) AddToTail(val int) {
	newNode := &DoubleNode{Val: val}
	if dll.Tail == nil { // If the list is empty
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) RemoveFromHead() (int, bool) {
	if dll.Head == nil { // If the list is empty
		return 0, false
	}
	val := dll.Head.Val
	if dll.Head == dll.Tail { // If there's only one element
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	}
	dll.size--
	return val, true
}

func (dll *DoublyLinkedList) RemoveFromTail() (int, bool) {
	if dll.Tail == nil { // If the list is empty
		return 0, false
	}
	val := dll.Tail.Val
	if dll.Head == dll.Tail { // If there's only one element
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	}
	dll.size--
	return val, true
}

func (dll *DoublyLinkedList) Find(val int) *DoubleNode {
	current := dll.Head
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}

func (dll *DoublyLinkedList) Delete(val int) bool {
	current := dll.Head
	for current != nil {
		if current.Val == val {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else { // Removing head
				dll.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else { // Removing tail
				dll.Tail = current.Prev
			}
			dll.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (dll *DoublyLinkedList) GetSize() int {
	return dll.size
}

func (dll *DoublyLinkedList) Traverse() []int {
	values := []int{}
	current := dll.Head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}
