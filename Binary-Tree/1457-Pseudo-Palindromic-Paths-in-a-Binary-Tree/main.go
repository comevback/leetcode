package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	// 调用函数计算伪回文路径的数量
	result := pseudoPalindromicPaths(root)
	fmt.Println(result)
}

// TreeNode 是二叉树节点的结构定义
type TreeNode struct {
	Val   int       // 节点的整数值
	Left  *TreeNode // 左子节点指针
	Right *TreeNode // 右子节点指针
}

var res int // 全局变量，用于存储伪回文路径的数量

// pseudoPalindromicPaths 计算以 root 为根节点的伪回文路径的数量
func pseudoPalindromicPaths(root *TreeNode) int {
	res = 0
	travel(root, make(map[int]int)) // 调用 travel 函数进行遍历
	return res
}

// travel 函数递归遍历二叉树，并统计每条路径中节点值的出现次数
func travel(root *TreeNode, hsMap map[int]int) {
	if root == nil {
		return
	}

	hsMap[root.Val] += 1 // 将当前节点的值加入到 hsMap 中

	if root.Left == nil && root.Right == nil { // 如果是叶子节点
		if check(hsMap) { // 检查当前路径是否构成伪回文路径
			res += 1 // 如果是，则结果加一
		}
	}

	if root.Left != nil {
		travel(root.Left, hsMap) // 递归遍历左子树
	}

	if root.Right != nil {
		travel(root.Right, hsMap) // 递归遍历右子树
	}

	hsMap[root.Val] -= 1 // 回溯：将当前节点值的计数减一

	// 在 Go 语言中，当你从 map 中删除一个键值对时，如果删除操作使得该键对应的值变为 map 的零值（例如 0、""、nil 等），
	// 则该键值对会被自动从 map 中删除，不再占用额外的内存空间。
	// if hsMap[root.Val] == 0 {
	// 	delete(hsMap, root.Val)
	// }
}

// check 函数检查当前路径节点值的出现次数，判断是否构成伪回文路径
func check(hsMap map[int]int) bool {
	odd := 0
	for _, value := range hsMap {
		if value%2 == 1 {
			odd += 1
		}
	}

	return odd <= 1 // 如果奇数次出现的节点值不超过一个，则返回 true，表示构成伪回文路径
}
