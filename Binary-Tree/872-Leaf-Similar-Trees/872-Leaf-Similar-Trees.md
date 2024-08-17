# 872. Leaf-Similar Trees

Solved
Easy
Topics
Companies
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Example 1:

Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true
Example 2:

Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false

Constraints:

The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].

---

# Code

```go
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leafSimilar 函数计算两棵树的叶值序列是否相同
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1, arr2 := []int{}, []int{}

	travel(root1, &arr1)
	travel(root2, &arr2)

	// 如果两个数组长度不同，则两棵树的叶值序列不同
	if len(arr1) != len(arr2) {
		return false
	}

	// 比较两个数组的元素是否相同，如果有不同的元素，则两棵树的叶值序列不同
	for i, v := range arr1 {
		if arr2[i] != v {
			return false
		}
	}

	// 如果两个数组的元素都相同，则两棵树的叶值序列相同
	return true
}

// 定义 travel 函数，用于遍历树的叶值序列，如果是叶子节点，则添加到数组中
func travel(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	travel(root.Left, arr)
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
	}
	travel(root.Right, arr)
}
```
