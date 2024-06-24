# 701. Insert into a Binary Search Tree
Solved
Medium
Topics
Companies
You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

Example 1:

Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:

Example 2:

Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]
Example 3:

Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
 

Constraints:

The number of nodes in the tree will be in the range [0, 104].
-108 <= Node.val <= 108
All the values Node.val are unique.
-108 <= val <= 108
It's guaranteed that val does not exist in the original BST.

---

# Code
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	current := root
	newNode := &TreeNode{
		Val: val,
	}

	// 如果树为空，直接返回新节点作为根节点
	if root == nil {
		root = newNode
		return root
	}

	// 遍历树，找到插入新节点的位置
	for {
		// 如果当前节点的值小于插入值，移动到右子树
		if current.Val < val {
			if current.Right != nil {
				current = current.Right
			} else {
				// 如果右子树为空，在此处插入新节点
				current.Right = newNode
				break
			}
		} else if current.Val > val {
			// 如果当前节点的值大于插入值，移动到左子树
			if current.Left != nil {
				current = current.Left
			} else {
				// 如果左子树为空，在此处插入新节点
				current.Left = newNode
				break
			}
		}
	}

	return root
}
```