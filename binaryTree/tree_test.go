package binarytree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	// Check if the root value is correct
	if root.Val != 10 {
		t.Errorf("Expected root value 10, but got %d", root.Val)
	}

	// Check left and right children
	if root.Left.Val != 5 {
		t.Errorf("Expected left child of root to be 5, but got %d", root.Left.Val)
	}
	if root.Right.Val != 15 {
		t.Errorf("Expected right child of root to be 15, but got %d", root.Right.Val)
	}
}

func TestSearch(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)

	// Test searching for an existing value
	found := root.Search(7)
	if found == nil || found.Val != 7 {
		t.Errorf("Expected to find 7, but found %v", found)
	}

	// Test searching for a non-existing value
	notFound := root.Search(20)
	if notFound != nil {
		t.Errorf("Expected to not find 20, but found %v", notFound)
	}
}

func TestMin(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)

	// Test finding the minimum value
	minNode := root.Min()
	if minNode.Val != 3 {
		t.Errorf("Expected minimum value to be 3, but got %d", minNode.Val)
	}
}

func TestMax(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)

	// Test finding the maximum value
	maxNode := root.Max()
	if maxNode.Val != 15 {
		t.Errorf("Expected maximum value to be 15, but got %d", maxNode.Val)
	}
}

func TestDelete(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	// Delete a node with no children
	root.Delete(3)
	if root.Left.Left != nil {
		t.Errorf("Expected left child of 5 to be nil, but got %v", root.Left.Left)
	}

	// Delete a node with one child
	root.Delete(7)
	if root.Left.Right != nil {
		t.Errorf("Expected right child of 5 to be nil, but got %v", root.Left.Right)
	}

	// Delete a node with two children
	root.Delete(15)
	if root.Right.Val != 18 {
		t.Errorf("Expected right child of root to be 18, but got %d", root.Right.Val)
	}

	// Delete the root node
	root.Delete(10)
	if root.Val != 12 {
		t.Errorf("Expected new root value to be 12 after deleting 10, but got %d", root.Val)
	}
}
