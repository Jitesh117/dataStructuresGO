package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Insert(value int) {
	if value < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

func (node *TreeNode) Search(value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value == node.Val {
		return node
	} else if value < node.Val {
		return node.Left.Search(value)
	} else {
		return node.Right.Search(value)
	}
}

func (node *TreeNode) Min() *TreeNode {
	if node.Left == nil {
		return node
	}
	return node.Left.Min()
}

func (node *TreeNode) Max() *TreeNode {
	if node.Right == nil {
		return node
	}
	return node.Right.Max()
}

func (node *TreeNode) Delete(value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Val {
		node.Left = node.Left.Delete(value)
	} else if value > node.Val {
		node.Right = node.Right.Delete(value)
	} else {
		// Case 1: No child
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// Case 2: One child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Case 3: Two children
		minRight := node.Right.Min()
		node.Val = minRight.Val
		node.Right = node.Right.Delete(minRight.Val)
	}
	return node
}
