# 998. Maximum Binary Tree II
Solved
Medium
Topics
Companies
A maximum tree is a tree where every node has a value greater than any other value in its subtree.

You are given the root of a maximum binary tree and an integer val.

Just as in the previous problem, the given tree was constructed from a list a (root = Construct(a)) recursively with the following Construct(a) routine:

If a is empty, return null.
Otherwise, let a[i] be the largest element of a. Create a root node with the value a[i].
The left child of root will be Construct([a[0], a[1], ..., a[i - 1]]).
The right child of root will be Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]]).
Return root.
Note that we were not given a directly, only a root node root = Construct(a).

Suppose b is a copy of a with the value val appended to it. It is guaranteed that b has unique values.

Return Construct(b).

Example 1:

Input: root = [4,1,3,null,null,2], val = 5
Output: [5,4,null,1,3,null,null,2]
Explanation: a = [1,4,2,3], b = [1,4,2,3,5]
Example 2:

Input: root = [5,2,4,null,1], val = 3
Output: [5,2,4,null,1,null,3]
Explanation: a = [2,1,5,4], b = [2,1,5,4,3]
Example 3:

Input: root = [5,2,3,null,1], val = 4
Output: [5,2,4,null,1,3]
Explanation: a = [2,1,5,3], b = [2,1,5,3,4]

Constraints:

The number of nodes in the tree is in the range [1, 100].
1 <= Node.val <= 100
All the values of the tree are unique.
1 <= val <= 100

---

# Code
```go
package main

import "fmt"

func main() {
	// 构建初始的最大二叉树
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}

	// 向最大二叉树中插入新值
	res := insertIntoMaxTree(root, 5)
	fmt.Println(res) // 输出插入新值后的树
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ***************************************************************************************************************
// insertIntoMaxTree1 在最大二叉树中插入新值的递归方法
func insertIntoMaxTree1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		// 如果树为空，则创建一个新节点
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		// 如果新值大于当前节点值，创建新节点，并将当前树作为左子树
		newNode := &TreeNode{
			Val:  val,
			Left: root,
		}
		return newNode
	} else {
		// 否则递归插入到右子树
		root.Right = insertIntoMaxTree1(root.Right, val)
		return root
	}
}

// ***************************************************************************************************************
// insertIntoMaxTree 通过转换数组和重建树的方式在最大二叉树中插入新值
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	res := TreeToArr(root)                     // 将树转换为数组
	res = append(res, val)                     // 将新值添加到数组
	resTree := constructMaximumBinaryTree(res) // 通过数组重建最大二叉树
	return resTree
}

// TreeToArr 将二叉树转换为数组
func TreeToArr(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	res = append(res, root.Val)

	res = append(TreeToArr(root.Left), res...)
	res = append(res, TreeToArr(root.Right)...)
	return res
}

// constructMaximumBinaryTree 通过数组构建最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var max, index int
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}

	node := &TreeNode{
		Val: max,
	}

	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])

	return node
}
```