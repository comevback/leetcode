# 1339. Maximum Product of Splitted Binary Tree

Solved
Medium
Topics
Companies
Hint
Given the root of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized.

Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it modulo 109 + 7.

Note that you need to maximize the answer before taking the mod and not after taking it.

Example 1:

Input: root = [1,2,3,4,5,6]
Output: 110
Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11\*10)
Example 2:

Input: root = [1,null,2,3,4,null,null,5,6]
Output: 90
Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15\*6)

Constraints:

The number of nodes in the tree is in the range [2, 5 * 104].
1 <= Node.val <= 104

---

# Code

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 创建一个测试用例的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}

	// 计算并输出二叉树中分割后的最大乘积
	res := maxProduct(root)
	fmt.Println(res) // 应输出 110
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sums []int // 全局变量，用于存储每个子树的和

// maxProduct 函数计算二叉树中分割后的最大乘积
func maxProduct(root *TreeNode) int {
	sums = []int{}             // 初始化 sums 数组
	allSum := getSum(root)     // 计算整个树的和
	maxPro := 0                // 初始化最大乘积为 0
	num := math.Pow(10, 9) + 7 // 定义模数，用于计算最终结果

	// 遍历 sums 数组，计算每个子树分割后的乘积
	for i := len(sums) - 1; i >= 0; i-- {
		prod := sums[i] * (allSum - sums[i]) // 计算乘积
		if prod > maxPro {
			maxPro = prod // 更新最大乘积
		}
	}
	maxPro = maxPro % int(num) // 取模
	return maxPro              // 返回最终结果
}

// getSum 函数递归计算每个子树的和，并将其存储在 sums 数组中
func getSum(root *TreeNode) int {
	if root == nil {
		return 0 // 处理根节点为 nil 的情况
	}

	left := getSum(root.Left)      // 递归计算左子树的和
	right := getSum(root.Right)    // 递归计算右子树的和
	sum := left + right + root.Val // 计算当前子树的和
	sums = append(sums, sum)       // 将当前子树的和存储在 sums 数组中
	return sum                     // 返回当前子树的和
}
```
