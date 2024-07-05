# 889. Construct Binary Tree from Preorder and Postorder Traversal

<https://labuladong.online/algo/data-structure/binary-tree-part2/#%E9%80%9A%E8%BF%87%E5%90%8E%E5%BA%8F%E5%92%8C%E5%89%8D%E5%BA%8F%E9%81%8D%E5%8E%86%E7%BB%93%E6%9E%9C%E6%9E%84%E9%80%A0%E4%BA%8C%E5%8F%89%E6%A0%91>

Medium

Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary tree of distinct values and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.

If there exist multiple answers, you can return any of them.

 

Example 1:
![lc-prepost](https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg)
> Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]

Example 2:
> Input: preorder = [1], postorder = [1]
Output: [1]
 

Constraints:
> 1 <= preorder.length <= 30
1 <= preorder[i] <= preorder.length
All the values of preorder are unique.
postorder.length == preorder.length
1 <= postorder[i] <= postorder.length
All the values of postorder are unique.
It is guaranteed that preorder and postorder are the preorder traversal and postorder traversal of the same binary tree.


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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(preorder) == 0 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	lefthead := preorder[1]

	var divide int
	for i := 1; i < len(postorder); i++ {
		if postorder[i] == lefthead {
			divide = i
			break
		}
	}

	leftPostOrder := postorder[:divide+1]
	rightPostOrder := postorder[divide+1 : len(postorder)-1]

	leftPreOrder := preorder[1 : 1+len(leftPostOrder)]
	rightPreOrder := preorder[1+len(leftPostOrder):]

	root.Left = constructFromPrePost(leftPreOrder, leftPostOrder)
	root.Right = constructFromPrePost(rightPreOrder, rightPostOrder)

	return root
}
// **************************************************************************************************************
// review 7.5
func constructFromPrePost_75(preorder []int, postorder []int) *TreeNode {
	return build_75(preorder, postorder)
}

// build_75 递归地构建二叉树
func build_75(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 创建新的树节点
	newNode := &TreeNode{
		Val: preorder[0],
	}

	// 如果只有一个节点，直接返回
	if len(postorder) == 1 {
		return newNode
	}

	leftPreOrder, leftPostOrder, rightPreOrder, rightPostOrder := []int{}, []int{}, []int{}, []int{}

	// 查找左子树的前序和后序遍历数组
	for i, v := range preorder {
		if v == postorder[len(postorder)-2] {
			leftPreOrder = preorder[1:i]
			rightPreOrder = preorder[i:]
			break
		}
	}

	// 分割右子树的后序遍历数组
	rightPostOrder = postorder[len(postorder)-1-len(rightPreOrder) : len(postorder)-1]
	// 分割左子树的后序遍历数组
	leftPostOrder = postorder[:len(postorder)-1-len(rightPreOrder)]

	// 递归构建左子树
	newNode.Left = build_75(leftPreOrder, leftPostOrder)
	// 递归构建右子树
	newNode.Right = build_75(rightPreOrder, rightPostOrder)

	return newNode
}
```