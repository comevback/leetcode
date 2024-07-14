# 99. Recover Binary Search Tree

Solved
Medium
Topics
Companies
You are given the root of a binary search tree (BST), where the values of exactly two nodes of the tree were swapped by mistake. Recover the tree without changing its structure.

Example 1:

Input: root = [1,3,null,null,2]
Output: [3,1,null,null,2]
Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.
Example 2:

Input: root = [3,1,4,null,null,2]
Output: [2,1,4,null,null,3]
Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.

Constraints:

The number of nodes in the tree is in the range [2, 1000].
-231 <= Node.val <= 231 - 1

Follow up: A solution using O(n) space is pretty straight-forward. Could you devise a constant O(1) space solution?

---

# Code

```go
package main

import "math"

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}

	recoverTree1(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 快速排序方法
func recoverTree(root *TreeNode) {
	// 将树的节点值存储到数组中
	arr := []int{}
	travel(root, &arr)
	// 对数组进行排序
	insertSort(arr)
	// 将排序后的值重新赋值给树节点
	index := 0
	setVal(root, arr, &index)
}

// 中序遍历将树的节点值存储到数组中
func travel(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	travel(root.Left, arr)
	(*arr) = append((*arr), root.Val)
	travel(root.Right, arr)
}

// 将排序后的数组值重新赋值给树节点
func setVal(root *TreeNode, arr []int, index *int) {
	if root == nil {
		return
	}
	setVal(root.Left, arr, index)
	root.Val = arr[*index]
	*index += 1
	setVal(root.Right, arr, index)
}

// 插入排序算法
func insertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// --------------------------------------------------------------------------------------------------------------
// 空间复杂度为 O(1) 的方法
var currentLeft *TreeNode
var currentRight *TreeNode
var exchangeLeft *TreeNode
var exchangeRight *TreeNode

func recoverTree1(root *TreeNode) {
	if root == nil {
		return
	}
	// 初始化变量
	currentLeft = &TreeNode{Val: math.MinInt}
	currentRight = &TreeNode{Val: math.MaxInt}
	// 遍历树找到需要交换的节点
	travelCompare(root)
	// 设置正确的节点值
	setValto(root)
	// 交换节点值
	exchangeLeft.Val, exchangeRight.Val = exchangeRight.Val, exchangeLeft.Val
}

// 遍历树找到需要交换的节点（从左到右）
func travelCompare(root *TreeNode) {
	if root == nil {
		return
	}

	travelCompare(root.Left)
	if currentLeft.Val <= root.Val {
		currentLeft = root
	} else {
		exchangeLeft = currentLeft
		return
	}
	travelCompare(root.Right)
}

// 遍历树找到需要交换的节点（从右到左）
func setValto(root *TreeNode) {
	if root == nil {
		return
	}

	setValto(root.Right)
	if currentRight.Val >= root.Val {
		currentRight = root
	} else {
		exchangeRight = currentRight
		return
	}
	setValto(root.Left)
}

// -------------------------------------------------------------------------------------------------------------
// 改进后的空间复杂度为 O(1) 的方法
func recoverTree2(root *TreeNode) {
	var first, second *TreeNode
	prev := &TreeNode{Val: math.MinInt32} // 判断中序遍历的有序性
	// 中序遍历找到需要交换的节点
	inorderTraverse(root, &first, &second, &prev)
	// 交换节点值
	first.Val, second.Val = second.Val, first.Val
}

// 中序遍历函数
func inorderTraverse(root *TreeNode, first, second, prev **TreeNode) {
	if root == nil {
		return
	}
	inorderTraverse(root.Left, first, second, prev)
	// 中序遍历代码位置，找出那两个节点
	if root.Val < (*prev).Val {
		// root 不符合有序性
		if *first == nil {
			// 第一个错位节点是 prev
			*first = *prev
		}
		// 第二个错位节点是 root
		*second = root
	}
	*prev = root
	inorderTraverse(root.Right, first, second, prev)
}
```
