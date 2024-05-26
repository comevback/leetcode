# 106. Construct Binary Tree from Inorder and Postorder Traversal

the same as:
<https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/>

Medium

Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.


Example 1
![tree](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)
> Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]

Example 2:
> Input: inorder = [-1], postorder = [-1]
Output: [-1]
 

Constraints:
> 1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder and postorder consist of unique values.
Each value of postorder also appears in inorder.
inorder is guaranteed to be the inorder traversal of the tree.
postorder is guaranteed to be the postorder traversal of the tree.

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	var devide int
	for i := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			devide = i
			break
		}
	}

	rightInOrder := inorder[devide+1:]
	leftInOrder := inorder[:devide]

	rightPostOrder := postorder[len(leftInOrder) : len(postorder)-1]
	leftPostOrder := postorder[:len(leftInOrder)]

	root.Right = buildTree(rightInOrder, rightPostOrder)
	root.Left = buildTree(leftInOrder, leftPostOrder)

	return root
}
```