# 108. Convert Sorted Array to Binary Search Tree

Solved
Easy
Topics
Companies
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced
binary search tree.

Example 1:

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:

Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.

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

func sortedArrayToBST(nums []int) *TreeNode {
	// 如果数组为空，返回 nil
	if len(nums) == 0 {
		return nil
	}

	// 定义数组的左右边界
	left := 0
	right := len(nums) - 1
	// 计算中间索引
	mid := left + (right-left)/2

	// 创建一个新节点，值为数组中间索引处的值
	newNode := &TreeNode{Val: nums[mid]}
	// 递归构建左子树，传入中间索引左边的子数组
	newNode.Left = sortedArrayToBST(nums[:mid])
	// 递归构建右子树，传入中间索引右边的子数组
	newNode.Right = sortedArrayToBST(nums[mid+1:])

	// 返回新构建的树节点
	return newNode
}
```
