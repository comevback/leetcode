# 437. Path Sum III
Solved
Medium
Topics
Companies
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

Example 1:

Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.
Example 2:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3

Constraints:

The number of nodes in the tree is in the range [0, 1000].
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000

---

# Code
```go
package main

import "fmt"

func main() {
	// 构建测试用的二叉树
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 1}

	// 计算路径和为 8 的路径数量
	res := pathSum(root, 8)
	fmt.Println(res) // 输出结果
}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果
var res int

// pathSum 返回二叉树中路径和为 targetSum 的路径数量
func pathSum(root *TreeNode, targetSum int) int {
	res = 0
	hsMap := map[int]int{0: 1} // 前缀和映射初始化为 0:1
	travel(root, []int{0}, hsMap, 0, targetSum)
	return res
}

// travel 递归遍历二叉树，并记录路径和为 targetSum 的路径数量
func travel(root *TreeNode, prefixSum []int, hsMap map[int]int, sum int, targetSum int) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	sum += root.Val                    // 更新当前路径和
	prefixSum = append(prefixSum, sum) // 更新前缀和数组

	need := sum - targetSum // 计算需要的前缀和
	if hsMap[need] > 0 {
		res += hsMap[need] // 如果前缀和映射中存在需要的前缀和，增加结果计数器
	}

	hsMap[sum] += 1                                      // 更新前缀和映射
	travel(root.Left, prefixSum, hsMap, sum, targetSum)  // 递归遍历左子树
	travel(root.Right, prefixSum, hsMap, sum, targetSum) // 递归遍历右子树
	hsMap[sum] -= 1                                      // 回溯时更新前缀和映射
}
```
