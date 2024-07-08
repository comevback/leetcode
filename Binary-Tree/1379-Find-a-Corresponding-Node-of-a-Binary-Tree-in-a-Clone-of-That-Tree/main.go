package main

func main() {

}

// 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res *TreeNode // 全局变量，用于存储找到的目标节点的克隆节点

// getTargetCopy 函数在克隆树中找到目标节点对应的节点
func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	res = nil                        // 重置 res 为 nil
	travel(original, cloned, target) // 调用 travel 函数进行遍历
	return res                       // 返回找到的节点
}

// travel 函数递归遍历原始树和克隆树，查找目标节点
func travel(original *TreeNode, cloned *TreeNode, target *TreeNode) {
	if original == nil {
		return // 如果当前节点为空，直接返回
	}

	if original == target {
		res = cloned // 如果找到了目标节点，将克隆节点赋值给 res
		return       // 找到目标节点后可以直接返回，避免不必要的递归
	}

	// 递归遍历左子树
	travel(original.Left, cloned.Left, target)
	// 递归遍历右子树
	travel(original.Right, cloned.Right, target)
}
