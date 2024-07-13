# 863. All Nodes Distance K in Binary Tree

Solved
Medium
Topics
Companies
Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.

You can return the answer in any order.

Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
Output: [7,4,1]
Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.
Example 2:

Input: root = [1], target = 1, k = 3
Output: []

Constraints:

The number of nodes in the tree is in the range [1, 500].
0 <= Node.val <= 500
All the values Node.val are unique.
target is the value of one of the nodes in the tree.
0 <= k <= 1000

---

# Code

```go
package main

import "fmt"

func main() {
	// 构建树结构
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}

	// 调用 distanceK 函数，计算距离目标节点距离为 k 的所有节点
	res := distanceK1(root, root.Left, 2)
	fmt.Println(res)
}

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int // 存储结果节点值

// distanceK 函数找到距离目标节点 target 距离为 k 的所有节点
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	res = []int{}
	// 获取目标节点的深度和从根到目标节点的路径
	targetDepth, list := getDis(root, target)
	// 遍历树，找到距离为 k 的节点
	travel(root, targetDepth, k, list)
	return res
}

// getDis 函数获取目标节点的深度和从根到目标节点的路径
func getDis(root *TreeNode, target *TreeNode) (int, []int) {
	if root == nil {
		return -1, []int{}
	}

	dis := -1
	leftDis, leftList := getDis(root.Left, target)
	rightDis, rightList := getDis(root.Right, target)

	// 如果目标节点在左子树中，返回左子树的距离加1，并将当前节点加入路径
	if leftDis != -1 {
		dis = leftDis + 1
		leftList = append(leftList, root.Val)
		return dis, leftList
	}

	// 如果目标节点在右子树中，返回右子树的距离加1，并将当前节点加入路径
	if rightDis != -1 {
		dis = rightDis + 1
		rightList = append(rightList, root.Val)
		return dis, rightList
	}

	// 如果当前节点是目标节点，返回距离0，并将当前节点加入路径
	if root == target {
		dis = 0
		return dis, []int{root.Val}
	}

	return -1, []int{}
}

// travel 函数遍历树，找到距离为 targetDis 的节点
func travel(root *TreeNode, dis int, targetDis int, list []int) {
	if root == nil {
		return
	}

	// 如果当前节点在路径中，将其从路径中移除
	if len(list) > 0 && root.Val == list[len(list)-1] {
		list = list[:len(list)-1]
	}

	// 如果当前节点距离等于目标距离，将其值加入结果
	if dis == targetDis {
		res = append(res, root.Val)
		// 如果当前节点是目标节点，没必要继续向下遍历
		if len(list) == 0 {
			return
		} else {
			// 如果当前节点在路径中，继续向目标节点的方向遍历
			if root.Left != nil && root.Left.Val == list[len(list)-1] {
				travel(root.Left, dis-1, targetDis, list)
			}
			if root.Right != nil && root.Right.Val == list[len(list)-1] {
				travel(root.Right, dis-1, targetDis, list)
			}
		}
	} else if dis > targetDis { // 如果当前节点距离大于目标距离，则继续向下遍历子树
		// 如果当前节点是目标节点或者目标节点的子节点，没必要继续向下遍历
		if len(list) == 0 {
			return
		} else {
			// 如果当前节点在路径中，继续向目标节点的方向遍历
			if root.Left != nil && root.Left.Val == list[len(list)-1] {
				travel(root.Left, dis-1, targetDis, list)
			}
			if root.Right != nil && root.Right.Val == list[len(list)-1] {
				travel(root.Right, dis-1, targetDis, list)
			}
		}
	} else if dis < targetDis { // 如果当前节点距离小于目标距离，则继续向下遍历子树
		// 如果当前节点是目标节点或者目标节点的子节点，继续向下遍历
		if len(list) == 0 {
			travel(root.Left, dis+1, targetDis, list)
			travel(root.Right, dis+1, targetDis, list)
		} else {
			// 如果当前节点在路径中，继续向目标节点的方向遍历
			if root.Left != nil && root.Left.Val == list[len(list)-1] {
				travel(root.Left, dis-1, targetDis, list)
			}
			if root.Right != nil && root.Right.Val == list[len(list)-1] {
				travel(root.Right, dis-1, targetDis, list)
			}
			// 如果当前节点的左右子节点不在路径中，则向下遍历子树
			if root.Left != nil && root.Left.Val != list[len(list)-1] {
				travel(root.Left, dis+1, targetDis, list)
			}
			if root.Right != nil && root.Right.Val != list[len(list)-1] {
				travel(root.Right, dis+1, targetDis, list)
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------
// 生成图，BFS
// 声明全局变量用于存储图结构和访问记录
var graph map[int][]int
var memo map[int]bool

// distanceK1 函数用于查找距离 target 节点 k 距离的所有节点值
func distanceK1(root *TreeNode, target *TreeNode, k int) []int {
	graph = make(map[int][]int) // 初始化图
	memo = make(map[int]bool)   // 初始化访问记录
	dis := 0                    // 初始化距离为 0
	geneGraph(root)             // 生成图
	queue := []int{}            // 初始化队列
	queue = append(queue, target.Val)
	memo[target.Val] = true // 标记目标节点已访问
	for len(queue) != 0 {
		if dis == k {
			return queue // 当距离等于 k 时，返回当前队列中的节点
		}
		length := len(queue)
		for i := 0; i < length; i++ {
			for _, v := range graph[queue[i]] {
				if !memo[v] {
					queue = append(queue, v)
					memo[v] = true // 标记节点已访问
				}
			}
		}
		queue = queue[length:] // 更新队列，只保留未处理的节点
		dis += 1               // 增加距离
	}
	return []int{} // 如果未找到，返回空切片
}

// geneGraph 函数用于生成图结构
func geneGraph(root *TreeNode) {
	if root == nil {
		return
	}

	// 如果当前节点的左右子节点不为空，将其加入图中，同时将当前节点加入子节点的图中
	if root.Left != nil {
		graph[root.Val] = append(graph[root.Val], root.Left.Val)
		graph[root.Left.Val] = append(graph[root.Left.Val], root.Val)
	}
	if root.Right != nil {
		graph[root.Val] = append(graph[root.Val], root.Right.Val)
		graph[root.Right.Val] = append(graph[root.Right.Val], root.Val)
	}

	geneGraph(root.Left)
	geneGraph(root.Right)
}
```
