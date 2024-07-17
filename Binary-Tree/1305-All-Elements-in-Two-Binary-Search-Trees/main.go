package main

import (
	"errors"
	"fmt"
)

func main() {
	// 示例用法，构建两个BST
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 0}
	root2.Right = &TreeNode{Val: 3}

	// 合并两个BST中的所有元素并返回排序后的结果
	res := getAllElements(root, root2)
	fmt.Println(res) // 输出 [0, 1, 1, 2, 3, 4]
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int       // 节点值
	Left  *TreeNode // 左子节点
	Right *TreeNode // 右子节点
}

// getAllElements 合并两个BST中的所有元素并返回排序后的结果
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	Bst1 := NewBST(root1) // 初始化第一个BST的迭代器
	Bst2 := NewBST(root2) // 初始化第二个BST的迭代器

	for {
		check1, err1 := Bst1.checkNext() // 查看第一个BST的下一个元素
		check2, err2 := Bst2.checkNext() // 查看第二个BST的下一个元素

		if err1 != nil && err2 != nil {
			break // 如果两个BST都没有下一个元素，跳出循环
		} else if err1 != nil {
			poped2, _ := Bst2.next() // 如果第一个BST没有下一个元素，取第二个BST的下一个元素
			res = append(res, poped2)
		} else if err2 != nil {
			poped1, _ := Bst1.next() // 如果第二个BST没有下一个元素，取第一个BST的下一个元素
			res = append(res, poped1)
		} else {
			if check1 <= check2 {
				poped1, _ := Bst1.next() // 如果第一个BST的下一个元素小于或等于第二个BST的下一个元素，取第一个BST的下一个元素
				res = append(res, poped1)
			} else {
				poped2, _ := Bst2.next() // 否则，取第二个BST的下一个元素
				res = append(res, poped2)
			}
		}
	}

	return res
}

// inOrder 返回二叉树的中序遍历结果
func inOrder(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	current := root
	for len(stack) != 0 || current != nil {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, popped.Val)
		current = popped.Right
	}
	return res
}

// BST 结构定义，用于迭代二叉搜索树
type BST struct {
	stack []*TreeNode // 栈用于迭代中序遍历
	root  *TreeNode   // 当前遍历的节点
}

// NewBST 初始化一个新的BST迭代器
func NewBST(root *TreeNode) BST {
	return BST{
		stack: []*TreeNode{},
		root:  root,
	}
}

// next 返回BST的下一个最小元素
func (bst *BST) next() (int, error) {
	if len(bst.stack) == 0 && bst.root == nil {
		return 0, errors.New("empty") // 如果没有下一个元素，返回错误
	}
	for bst.root != nil {
		bst.stack = append(bst.stack, bst.root)
		bst.root = bst.root.Left
	}
	popped := bst.stack[len(bst.stack)-1]
	bst.stack = bst.stack[:len(bst.stack)-1]
	bst.root = popped.Right
	return popped.Val, nil
}

// checkNext 查看BST的下一个最小元素，但不移除
func (bst *BST) checkNext() (int, error) {
	if len(bst.stack) == 0 && bst.root == nil {
		return 0, errors.New("empty") // 如果没有下一个元素，返回错误
	}
	for bst.root != nil {
		bst.stack = append(bst.stack, bst.root)
		bst.root = bst.root.Left
	}
	popped := bst.stack[len(bst.stack)-1]
	return popped.Val, nil
}
