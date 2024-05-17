# 1325. Delete Leaves With a Given Value

Medium

Given a binary tree root and an integer target, delete all the leaf nodes with value target.

Note that once you delete a leaf node with value target, if its parent node becomes a leaf node and has the value target, it should also be deleted (you need to continue doing that until you cannot).


Example 1:
![1](https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png)
> Input: root = [1,2,3,2,null,2,4], target = 2
Output: [1,null,3,null,4]
Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left). 
After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).

Example 2:
![2](https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png)
> Input: root = [1,3,3,3,2], target = 3
Output: [1,3,null,null,2]

Example 3:
![3](https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png)
> Input: root = [1,2,null,2,null,2], target = 2
Output: [1]
Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
 
Constraints:
> The number of nodes in the tree is in the range [1, 3000].
1 <= Node.val, target <= 1000

---

# Code
```go
package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	printTree(tree)
	removeLeafNodes(tree, 2)
	printTree(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// =====================================================================================================================
// 1. 双重指针
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	dfs(&root, target) // 传入指针的地址，这样可以修改指针的值
	return root
}

// dfs函数，接收一个指向指针的指针，以及目标值
func dfs(root **TreeNode, target int) {
	if (*root).Left != nil {
		dfs(&(*root).Left, target)
	}

	if (*root).Right != nil {
		dfs(&(*root).Right, target)
	}
	// 如果当前节点是叶子节点，且值等于目标值，将当前节点置为nil
	if (*root).Val == target && (*root).Left == nil && (*root).Right == nil {
		*root = nil
	}
}

// =====================================================================================================================
// 2. 直接赋值，from leetcode
func removeLeafNodes1(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}

	// 因为removeLeafNodes函数返回的是 *TreeNode，所以可以在此直接把指针赋值为nil
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Val == target && root.Left == nil && root.Right == nil {
		root = nil
		return root
	}
	return root
}

// 打印树方法
func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
```