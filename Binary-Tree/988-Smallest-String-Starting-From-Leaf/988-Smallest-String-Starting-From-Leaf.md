# 988. Smallest String Starting From Leaf
Solved
Medium
Topics
Companies
You are given the root of a binary tree where each node has a value in the range [0, 25] representing the letters 'a' to 'z'.

Return the lexicographically smallest string that starts at a leaf of this tree and ends at the root.

As a reminder, any shorter prefix of a string is lexicographically smaller.

For example, "ab" is lexicographically smaller than "aba".
A leaf of a node is a node that has no children.

Example 1:

Input: root = [0,1,2,3,4,3,4]
Output: "dba"
Example 2:

Input: root = [25,1,3,1,3,0,2]
Output: "adz"
Example 3:

Input: root = [2,2,1,null,1,0,null,0]
Output: "abc"

Constraints:

The number of nodes in the tree is in the range [1, 8500].
0 <= Node.val <= 25

---

# Code
```go
package main

import "fmt"

func main() {
	// 构造测试用例 [4,0,1,1]
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	// 使用 smallestFromLeaf1 函数获取从叶子到根的最小字符串
	res := smallestFromLeaf1(root)
	fmt.Println(res)
}

// TreeNode 定义二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// compare 函数比较两个整数数组，返回字典序较小的数组
func compare(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 && len(arr2) == 0 {
		return []int{}
	} else if len(arr1) == 0 {
		return arr2
	} else if len(arr2) == 0 {
		return arr1
	}

	min := 0
	if len(arr1) <= len(arr2) {
		min = len(arr1)
	} else {
		min = len(arr2)
	}

	// 比较两个数组的每个元素
	for i := 0; i < min; i++ {
		if arr1[i] < arr2[i] {
			return arr1
		} else if arr1[i] > arr2[i] {
			return arr2
		}
	}

	// 如果前面的元素都相等，返回较短的数组
	if len(arr1) <= len(arr2) {
		return arr1
	} else {
		return arr2
	}
}

// *************************************************  遍历  *******************************************************
// 定义一个全局变量 minStr 用于存储最小字符串表示
var minStr []int

// smallestFromLeaf 函数返回从叶子节点到根节点的最小字符串表示
func smallestFromLeaf(root *TreeNode) string {
	minStr = []int{}
	// 递归遍历树，收集所有路径
	travel(root, []int{})
	res := ""
	// 将最小字符串表示转换为字符
	for _, v := range minStr {
		res += string(v + 'a')
	}
	return res
}

// travel 函数递归遍历树，收集从叶子节点到根节点的路径
func travel(root *TreeNode, arr []int) {
	// 将当前节点的值添加到路径的开头
	arr = append([]int{root.Val}, arr...)

	// 如果当前节点是叶子节点，比较并更新最小字符串表示
	if root.Left == nil && root.Right == nil {
		if len(minStr) == 0 {
			minStr = arr
		} else {
			minStr = compare(minStr, arr)
		}
	}

	// 递归遍历左子树
	if root.Left != nil {
		travel(root.Left, arr)
	}

	// 递归遍历右子树
	if root.Right != nil {
		travel(root.Right, arr)
	}
}

// *************************************************  分解  *******************************************************
// smallestFromLeaf1 函数返回从叶子节点到根节点的最小字符串表示
func smallestFromLeaf1(root *TreeNode) string {
	// 收集所有路径
	arrs := gatherArrs(root)

	minStr := []int{}
	// 比较并找到最小的路径
	for _, v := range arrs {
		minStr = compare(minStr, v)
	}

	res := ""
	// 将最小路径转换为字符串
	for _, v := range minStr {
		res += string(v + 'a')
	}
	return res
}

// gatherArrs 函数收集树中所有从叶子节点到根节点的路径
func gatherArrs(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int

	// 如果当前节点是叶子节点，返回包含当前节点值的路径
	if root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}

	// 递归遍历左子树，并将当前节点值添加到每条路径中
	if root.Left != nil {
		left := gatherArrs(root.Left)
		for i := range left {
			left[i] = append(left[i], root.Val)
		}
		res = append(res, left...)
	}

	// 递归遍历右子树，并将当前节点值添加到每条路径中
	if root.Right != nil {
		right := gatherArrs(root.Right)
		for i := range right {
			right[i] = append(right[i], root.Val)
		}
		res = append(res, right...)
	}

	return res
}
```