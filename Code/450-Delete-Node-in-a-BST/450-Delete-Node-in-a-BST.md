# 450. Delete Node in a BST

Medium

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
 

Example 1:
![del_node_1](https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg)

> Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:
![del_node_supp](https://assets.leetcode.com/uploads/2020/09/04/del_node_supp.jpg)

> Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.

Example 3:
> Input: root = [], key = 0
Output: []
 

Constraints:
> The number of nodes in the tree is in the range [0, 104].
-105 <= Node.val <= 105
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105
 

> Follow up: Could you solve it with time complexity O(height of tree)?


---

# 思路
- 建立四个变量，分别是：current（跟踪当前节点的指针）、parentNode（跟踪当前节点的父节点的指针）、replacer（要替换当前节点的节点）、parentReplacer（要替换当前节点的节点的父节点）
- 分三种情况：要删除的节点没有子节点、要删除的节点只有一个子节点、要删除的节点有两个子节点

# 代码
```go
package main

import "fmt"

func main() {
	// root = [3,2,5,null,null,4,10,null,null,8,15,7]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	root = deleteNode(root, 5)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 删除节点 见 DSA/Binary-Tree/main.go
// 建立四个变量，分别是：current（跟踪当前节点的指针）、parentNode（跟踪当前节点的父节点的指针）、replacer（要替换当前节点的节点）、parentReplacer（要替换当前节点的节点的父节点）
// 分三种情况：要删除的节点没有子节点、要删除的节点只有一个子节点、要删除的节点有两个子节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	current := root
	var parentNode *TreeNode
	var replacer, parentReplacer *TreeNode

	if root == nil {
		return nil
	}

	// 找到要删除的节点
	for current != nil {
		if current.Val > key {
			parentNode = current
			current = current.Left
		} else if current.Val < key {
			parentNode = current
			current = current.Right
		} else {
			break
		}
	}

	// 没有找到要删除的节点，直接返回原树
	if current == nil {
		return root
	}

	// 找到了要删除的节点，分三种情况：没有子节点、只有一个子节点、有两个子节点
	// 没有子节点
	if current.Left == nil && current.Right == nil {
		if parentNode == nil {
			root = nil
		} else {
			if parentNode.Val > current.Val {
				parentNode.Left = nil
			} else {
				parentNode.Right = nil
			}
		}
		return root
	} else if current.Right == nil {
		// 只有一个左节点
		if parentNode == nil {
			root = current.Left
		} else {
			if parentNode.Val > current.Val {
				parentNode.Left = current.Left
			} else {
				parentNode.Right = current.Left
			}
		}
		return root
	} else {
		// 有两个子节点
		replacer = current.Right
		for replacer.Left != nil {
			parentReplacer = replacer
			replacer = replacer.Left
		}
	}

	// 如果replacer不是current的右子节点，即replacer不是current的直接右子节点，需要重新链接replacer的父节点和replacer的右子节点
	if parentReplacer != nil {
		parentReplacer.Left = replacer.Right
		replacer.Right = current.Right
	}
	// 无论replacer是否是current的直接右子节点，都需要重新链接replacer的左子节点
	replacer.Left = current.Left

	// 重新链接parentNode到replacer，如果parentNode为空，即current是根节点，把replacer作为新的根节点
	if parentNode == nil {
		root = replacer
	} else {
		// 如果parentNode不为空，重新链接parentNode到replacer
		if parentNode.Val > current.Val {
			parentNode.Left = replacer
		} else {
			parentNode.Right = replacer
		}
	}

	// 返回新的根节点
	return root
}
```