package main

func main() {

}

// TreeNode 定义了二叉树节点的结构
type TreeNode struct {
	Val   int       // 节点存储的整数值
	Left  *TreeNode // 指向左子节点的指针
	Right *TreeNode // 指向右子节点的指针
}

// res 全局变量，用于记录最长的路径长度
var res int

// amountOfTime 计算从给定起始节点到二叉树中任意节点的最长路径长度
func amountOfTime(root *TreeNode, start int) int {
	res = 0                      // 初始化全局变量 res
	HsMap := make(map[int][]int) // 创建一个哈希表来存储每个节点的邻接节点
	travel(root, HsMap)          // 遍历树，填充哈希表

	getLength(HsMap, start, -1, 0) // 从起始节点开始，使用 DFS 寻找最长路径
	return res                     // 返回最长路径的长度
}

// travel 递归遍历二叉树，并将每个节点的邻接节点信息存储在哈希表中
func travel(root *TreeNode, hsMap map[int][]int) {
	if root == nil {
		return // 如果节点为空，则返回
	}
	if root.Left != nil {
		// 建立当前节点与左子节点的双向连接
		hsMap[root.Val] = append(hsMap[root.Val], root.Left.Val)
		hsMap[root.Left.Val] = append(hsMap[root.Left.Val], root.Val)
	}

	if root.Right != nil {
		// 建立当前节点与右子节点的双向连接
		hsMap[root.Val] = append(hsMap[root.Val], root.Right.Val)
		hsMap[root.Right.Val] = append(hsMap[root.Right.Val], root.Val)
	}

	// 递归遍历左右子树
	travel(root.Left, hsMap)
	travel(root.Right, hsMap)
}

// getLength 使用 DFS 方法计算从起始节点开始的最长路径
func getLength(hsMap map[int][]int, current int, last int, length int) {
	if length > res {
		res = length // 更新最长路径长度
	}
	for _, v := range hsMap[current] {
		if v != last { // 避免回访父节点
			getLength(hsMap, v, current, length+1) // 递归地访问邻接节点，增加路径长度
		}
	}
}
