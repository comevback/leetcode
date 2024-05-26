# 105. Construct Binary Tree from Preorder and Inorder Traversal

<https://labuladong.online/algo/data-structure/binary-tree-part2/#%E9%80%9A%E8%BF%87%E5%89%8D%E5%BA%8F%E5%92%8C%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86%E7%BB%93%E6%9E%9C%E6%9E%84%E9%80%A0%E4%BA%8C%E5%8F%89%E6%A0%91>

Medium

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

 

Example 1:

> Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
> Input: preorder = [-1], inorder = [-1]
Output: [-1]
 

Constraints:
> 1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果给的遍历数组长度为0，说明是空，返回nil
	if len(preorder) == 0 {
		return nil
	}

	// 这个子树的根节点一定是前序遍历数组的第一个，把这个根节点定义出来
	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}

	// devide定义为在中序遍历数组中分割左右子树的index，也就是前序遍历数组中头元素在中序遍历数组中的index
	var devide int

	// 找到这个divide
	for i, v := range inorder {
		if v == preorder[0] {
			devide = i
			break
		}
	}

	// 找到左右子树的中序遍历数组
	var leftInOrder, rightInOrder []int

	leftInOrder = inorder[:devide]
	if devide < len(inorder)-1 {
		rightInOrder = inorder[devide+1:]
	}

	// 找到左右子树的前序遍历数组
	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	rightPreOrder := preorder[1+len(leftInOrder):]

	// 把左右子树通过递归得到，附到这个节点上
	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)

	// 返回这个节点
	return root
}
```