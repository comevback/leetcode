# 98. Validate Binary Search Tree
Solved
Medium
Topics
Companies
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:

Input: root = [2,1,3]
Output: true
Example 2:

Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1

---

# Code
```go
// ****************************************************  遍历  ******************************************************
var measureVal int = math.MinInt // 初始化全局变量 measureVal 用于比较节点值
var res bool = true              // 初始化全局变量 res 用于存储验证结果

// isValidBST 函数用于检查二叉树是否为有效的二叉搜索树
func isValidBST(root *TreeNode) bool {
	measureVal = math.MinInt // 重置 measureVal 为最小整数值
	res = true               // 重置 res 为 true
	inOrderCheck(root)       // 调用中序遍历检查函数
	return res
}

// inOrderCheck 函数用于中序遍历二叉树并检查是否为有效的二叉搜索树
func inOrderCheck(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderCheck(root.Left)     // 递归遍历左子树
	if root.Val <= measureVal { // 如果当前节点值小于或等于前一个节点值，设置 res 为 false
		res = false
	} else {
		measureVal = root.Val // 更新 measureVal 为当前节点值
	}
	inOrderCheck(root.Right) // 递归遍历右子树
}

// ****************************************************  分治  ******************************************************
// isValidBST1 函数用于通过分治法检查二叉树是否为有效的二叉搜索树
func isValidBST1(root *TreeNode) bool {
	return check(root, math.MaxInt, math.MinInt) // 调用检查函数，初始最大值和最小值
}

// check 函数用于分治法检查二叉树是否为有效的二叉搜索树
func check(root *TreeNode, max int, min int) bool {
	if root.Left != nil && root.Right != nil { // 如果左右子节点都存在
		if root.Left.Val < root.Val && root.Right.Val > root.Val &&
			root.Left.Val < max && root.Right.Val < max &&
			root.Left.Val > min && root.Right.Val > min { // 检查当前节点值是否满足二叉搜索树条件
			return check(root.Left, root.Val, min) && check(root.Right, max, root.Val) // 递归检查左右子树
		} else {
			return false
		}
	} else if root.Left != nil { // 如果只有左子节点
		if root.Left.Val < root.Val &&
			root.Left.Val < max &&
			root.Left.Val > min { // 检查左子节点值是否满足二叉搜索树条件
			return check(root.Left, root.Val, min) // 递归检查左子树
		} else {
			return false
		}
	} else if root.Right != nil { // 如果只有右子节点
		if root.Right.Val > root.Val &&
			root.Right.Val < max &&
			root.Right.Val > min { // 检查右子节点值是否满足二叉搜索树条件
			return check(root.Right, max, root.Val) // 递归检查右子树
		} else {
			return false
		}
	} else { // 如果没有子节点，返回 true
		return true
	}
}
```