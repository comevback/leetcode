# 1469. 寻找所有的独生节点
二叉树中，如果一个节点是其父节点的唯一子节点，则称这样的节点为 “独生节点” 。二叉树的根节点不会是独生节点，因为它没有父节点。

给定一棵二叉树的根节点 root ，返回树中 所有的独生节点的值所构成的数组 。数组的顺序 不限 。

示例 1：

输入：root = [1,2,3,null,4]
输出：[4]
解释：浅蓝色的节点是唯一的独生节点。
节点 1 是根节点，不是独生的。
节点 2 和 3 有共同的父节点，所以它们都不是独生的。
示例 2：

输入：root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
输出：[6,2]
输出：浅蓝色的节点是独生节点。
请谨记，顺序是不限的。 [2,6] 也是一种可接受的答案。
示例 3：

输入：root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
输出：[77,55,33,66,44,22]
解释：节点 99 和 88 有共同的父节点，节点 11 是根节点。
其他所有节点都是独生节点。
示例 4：

输入：root = [197]
输出：[]
示例 5：

输入：root = [31,null,78,null,28]
输出：[78,28]

提示：
tree 中节点个数的取值范围是 [1, 1000]。
每个节点的值的取值范围是 [1, 10^6]。

---

# Code
```go
package main

func main() {

}

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 全局变量用于存储结果
var res []int

// getLonelyNodes 返回二叉树中所有孤独节点的值
// 孤独节点定义为其父节点只有一个子节点的节点
func getLonelyNodes(root *TreeNode) []int {
	res = []int{}
	travel(root)
	return res
}

// travel 递归遍历二叉树，并记录所有孤独节点的值
func travel(root *TreeNode) {
	if root == nil {
		return // 如果当前节点为空，直接返回
	}

	// 如果当前节点只有一个右子节点
	if root.Left == nil && root.Right != nil {
		// 将右子节点的值添加到结果中
		res = append(res, root.Right.Val)
		// 递归遍历右子树
		travel(root.Right)
		// 如果当前节点只有一个左子节点
	} else if root.Left != nil && root.Right == nil {
		// 将左子节点的值添加到结果中
		res = append(res, root.Left.Val)
		// 递归遍历左子树
		travel(root.Left)
		// 如果当前节点有两个子节点
	} else {
		// 递归遍历左子树
		travel(root.Left)
		// 递归遍历右子树
		travel(root.Right)
	}
}
```