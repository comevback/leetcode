# 1379. Find a Corresponding Node of a Binary Tree in a Clone of That Tree

Solved
Easy
Topics
Companies
Given two binary trees original and cloned and given a reference to a node target in the original tree.

The cloned tree is a copy of the original tree.

Return a reference to the same node in the cloned tree.

Note that you are not allowed to change any of the two trees or the target node and the answer must be a reference to a node in the cloned tree.

Example 1:

Input: tree = [7,4,3,null,null,6,19], target = 3
Output: 3
Explanation: In all examples the original and cloned trees are shown. The target node is a green node from the original tree. The answer is the yellow node from the cloned tree.
Example 2:

Input: tree = [7], target = 7
Output: 7
Example 3:

Input: tree = [8,null,6,null,5,null,4,null,3,null,2,null,1], target = 4
Output: 4

Constraints:

The number of nodes in the tree is in the range [1, 104].
The values of the nodes of the tree are unique.
target node is a node from the original tree and is not null.

Follow up: Could you solve the problem if repeated values on the tree are allowed?

---

# Code

```go
package main

func main() {

}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res *TreeNode // 全局变量，用于存储找到的目标节点的克隆节点

// getTargetCopy 函数在克隆树中找到目标节点对应的节点
func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	res = nil                        // 重置 res 为 nil
	travel(original, cloned, target) // 调用 travel 函数进行遍历
	return res                       // 返回找到的节点
}

// travel 函数递归遍历原始树和克隆树，查找目标节点
func travel(original *TreeNode, cloned *TreeNode, target *TreeNode) {
	if original == nil {
		return // 如果当前节点为空，直接返回
	}

	if original == target {
		res = cloned // 如果找到了目标节点，将克隆节点赋值给 res
		return       // 找到目标节点后可以直接返回，避免不必要的递归
	}

	// 递归遍历左子树
	travel(original.Left, cloned.Left, target)
	// 递归遍历右子树
	travel(original.Right, cloned.Right, target)
}
```
