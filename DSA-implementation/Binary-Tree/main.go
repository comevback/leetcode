package main

import "errors"

func main() {

}

type ListNode struct {
	Val       int
	Frequency int
	Left      *ListNode
	Right     *ListNode
}

type BinaryTree struct {
	root *ListNode
}

func NewBinary() *BinaryTree {
	return &BinaryTree{}
}

func (binTree *BinaryTree) insert(value int) {
	newNode := &ListNode{
		Val: value,
	}

	if binTree.root == nil {
		binTree.root = newNode
		return
	}

	current := binTree.root
	for {
		if value < current.Val {
			if current.Left == nil {
				current.Left = newNode
				break
			}
			current = current.Left
		} else if value > current.Val {
			if current.Right == nil {
				current.Right = newNode
				break
			}
			current = current.Right
		} else {
			current.Frequency += 1
			break
		}
	}
}

func (binTree *BinaryTree) lookup(value int) (*ListNode, error) {
	current := binTree.root

	for current != nil {
		if value < current.Val {
			current = current.Left
		} else if value > current.Val {
			current = current.Right
		} else {
			return current, nil
		}
	}

	return nil, errors.New("didn't find")
}

// =============================================================================================================================

func (binTree *BinaryTree) remove(value int) {
	var parentNode *ListNode
	current := binTree.root
	if current == nil {
		return
	}

	for current != nil {
		if value < current.Val {
			parentNode = current
			current = current.Left
		} else if value > current.Val {
			parentNode = current
			current = current.Right
		} else {
			break
		}
	}

	if parentNode == nil {
		if current.Left == nil && current.Right == nil {
			binTree.root = nil
			return
		}

		if current.Right == nil && current.Left != nil {
			binTree.root = current.Left
			return
		}

		if current.Right != nil && current.Right.Left == nil {
			current.Right.Left = binTree.root.Left
			binTree.root = current.Right
			return
		}

	}

	if current.Left == nil && current.Right == nil {
		if parentNode.Val > current.Val {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
		current = nil
		return
	}

	if current.Right == nil && current.Left != nil {
		if parentNode.Val > current.Val {
			parentNode.Left = current.Left
		} else {
			parentNode.Right = current.Left
		}
		current = nil
		return
	}

	if current.Right != nil && current.Right.Left == nil {
		current.Left = current.Right.Left
		current = current.Right
		return
	}

	var replacer, replacerParent *ListNode

}
