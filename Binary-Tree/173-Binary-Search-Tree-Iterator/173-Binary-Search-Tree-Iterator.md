# 173. Binary Search Tree Iterator

Solved
Medium
Topics
Companies
Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):

BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
int next() Moves the pointer to the right, then returns the number at the pointer.
Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.

Example 1:

Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next(); // return 3
bSTIterator.next(); // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next(); // return 20
bSTIterator.hasNext(); // return False

Constraints:

The number of nodes in the tree is in the range [1, 105].
0 <= Node.val <= 106
At most 105 calls will be made to hasNext, and next.

Follow up:

Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?

---

# Code

```go
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
```
