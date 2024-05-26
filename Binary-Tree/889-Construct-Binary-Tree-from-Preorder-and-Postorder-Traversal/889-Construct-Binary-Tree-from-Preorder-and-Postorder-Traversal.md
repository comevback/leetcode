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
```