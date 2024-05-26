# 652. Find Duplicate Subtrees

Medium

Given the root of a binary tree, return all duplicate subtrees.

For each kind of duplicate subtrees, you only need to return the root node of any one of them.

Two trees are duplicate if they have the same structure with the same node values.


Example 1:
![e1](https://assets.leetcode.com/uploads/2020/08/16/e1.jpg)
> Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]

Example 2:
![e2](https://assets.leetcode.com/uploads/2020/08/16/e2.jpg)
> Input: root = [2,1,1]
Output: [[1]]

Example 3:
![e3](https://assets.leetcode.com/uploads/2020/08/16/e33.jpg)
> Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]
 

Constraints:

The number of the nodes in the tree will be in the range [1, 5000]
-200 <= Node.val <= 200

---

## 解题思路

需要知道以下两点：

1、以我为根的这棵二叉树（子树）长啥样？

2、以其他节点为根的子树都长啥样？

对于第一个问题，只要知道左右子树长啥样，再加上自己长啥样，就能知道以我为根的这棵二叉树长啥样。所以我们可以通过递归的方式，DFS 后序遍历这棵二叉树，求出每棵子树的序列化字符串。

对于第二个问题，我们可以使用哈希表辅助，对于每棵子树，我们将其序列化，然后存入哈希表中，如果这棵子树在哈希表中已经出现过了，说明以这棵子树为根的子树是重复的。

# Code

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//root = [0,0,0,0,null,null,0,null,null,0,0]
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  0,
			Left: nil,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	res := findDuplicateSubtrees(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var existed map[string]int
var duplicate []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	existed = make(map[string]int)
	duplicate = []*TreeNode{}

	if root == nil {
		return nil
	}

	dfs(root)
	return duplicate

}

// 1. 我的解法，检查两个子节点
func dfs(root *TreeNode) string {
	if root == nil {
		return ""
	}

	leftStr := dfs(root.Left)
	rightStr := dfs(root.Right)

	// =========================left============================

	if existed[leftStr] == 1 {
		duplicate = append(duplicate, root.Left)
		existed[leftStr] += 1
	}

	if existed[leftStr] == 0 && root.Left != nil {
		existed[leftStr] = 1
	}

	// =========================right============================

	if existed[rightStr] == 1 {
		duplicate = append(duplicate, root.Right)
		existed[rightStr] += 1
	}

	if existed[rightStr] == 0 && root.Right != nil {
		existed[rightStr] = 1
	}

	res := ""
	res += strconv.Itoa(root.Val)
	res = res + "-left->" + leftStr
	res = res + "-right->" + rightStr

	return res
}

// 2. 简化版，只检查本节点，而不是两个子节点
func dfs1(root *TreeNode) string {
	if root == nil {
		return ""
	}

	leftStr := dfs(root.Left)
	rightStr := dfs(root.Right)

	res := ""
	res += strconv.Itoa(root.Val)
	res = res + "-left->" + leftStr
	res = res + "-right->" + rightStr

	if existed[res] == 1 {
		duplicate = append(duplicate, root)
		existed[res] += 1
	}

	if existed[res] == 0 {
		existed[res] = 1
	}

	return res
}
```