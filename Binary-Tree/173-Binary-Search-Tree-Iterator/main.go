package main

import "fmt"

func main() {
	// 示例用法
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}

	// 使用递归中序遍历的迭代器
	iter1 := Constructor_1(root)
	for iter1.HasNext() {
		fmt.Println(iter1.Next()) // 输出 3, 7, 9, 15, 20
	}

	// 使用迭代中序遍历的迭代器
	iter2 := Constructor(root)
	for iter2.HasNext() {
		fmt.Println(iter2.Next()) // 输出 3, 7, 9, 15, 20
	}
}

// TreeNode 定义了二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator_1 使用递归中序遍历构建一个数组来实现迭代器
type BSTIterator_1 struct {
	arr   []int // 存储中序遍历的结果
	index int   // 当前遍历到的数组索引
}

// Constructor_1 初始化 BSTIterator_1
func Constructor_1(root *TreeNode) BSTIterator_1 {
	arr := []int{}
	travel(root, &arr)
	return BSTIterator_1{
		arr:   arr,
		index: 0,
	}
}

// Next 返回下一个最小元素
func (this *BSTIterator_1) Next() int {
	res := this.arr[this.index]
	this.index += 1
	return res
}

// HasNext 判断是否还有下一个元素
func (this *BSTIterator_1) HasNext() bool {
	return this.index < len(this.arr)
}

// travel 递归中序遍历二叉树并将结果存入数组
func travel(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	travel(root.Left, arr)
	*arr = append(*arr, root.Val)
	travel(root.Right, arr)
}

// ****************************************************  迭代中序遍历  *********************************************************
// BSTIterator 使用迭代的方式进行中序遍历
type BSTIterator struct {
	stack []*TreeNode // 使用栈来实现中序遍历
	root  *TreeNode   // 当前遍历的节点
}

// Constructor 初始化 BSTIterator
func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: []*TreeNode{},
		root:  root,
	}
}

// Next 返回下一个最小元素
func (this *BSTIterator) Next() int {
	for this.root != nil {
		this.stack = append(this.stack, this.root)
		this.root = this.root.Left
	}
	popped := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.root = popped.Right
	return popped.Val
}

// HasNext 判断是否还有下一个元素
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.root != nil
}
