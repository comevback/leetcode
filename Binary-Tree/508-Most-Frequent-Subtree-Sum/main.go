package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var hsMap map[int]int // 全局变量，用于存储每个子树和出现的频率

// findFrequentTreeSum 函数返回二叉树中出现频率最高的子树和
func findFrequentTreeSum(root *TreeNode) []int {
	hsMap = make(map[int]int) // 初始化哈希表
	travel(root)              // 遍历二叉树，计算每个子树的和并记录频率

	res := []int{} // 存储结果
	max := 0       // 存储最高频率

	// 遍历哈希表，找出出现频率最高的子树和
	for key, value := range hsMap {
		if value > max {
			res = []int{} // 发现新的最高频率时，清空结果数组
			res = append(res, key)
			max = value
		} else if value == max {
			res = append(res, key) // 频率相同时，添加到结果数组
		}
	}

	return res // 返回结果
}

// travel 函数递归遍历二叉树，计算每个子树的和并记录频率
func travel(root *TreeNode) int {
	if root == nil {
		return 0 // 如果当前节点为空，返回 0
	}

	// 计算当前子树的和
	sum := root.Val
	sum += travel(root.Left) + travel(root.Right)

	// 将当前子树和加入哈希表并更新频率
	hsMap[sum] += 1

	return sum // 返回当前子树的和
}
