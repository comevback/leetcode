package main

func main() {

}

// TreeNode 表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindElements 用于处理被污染的二叉树，并支持查找元素
type FindElements struct {
	tree  *TreeNode
	hsMap map[int]bool
}

// Constructor 初始化 FindElements 结构体，修复二叉树并建立值的哈希集合
func Constructor(root *TreeNode) FindElements {
	hsMap := make(map[int]bool)
	travel(root, 0, hsMap)
	return FindElements{
		tree:  root,
		hsMap: hsMap,
	}
}

// travel 递归地遍历二叉树，修复节点的值，并将值存入哈希集合
func travel(root *TreeNode, value int, hsMap map[int]bool) {
	if root == nil {
		return
	}

	// 修复当前节点的值
	root.Val = value
	// 将修复后的值存入哈希集合
	hsMap[root.Val] = true

	// 递归修复左子节点，值为当前值 * 2 + 1
	travel(root.Left, value*2+1, hsMap)
	// 递归修复右子节点，值为当前值 * 2 + 2
	travel(root.Right, value*2+2, hsMap)
}

// Find 在修复后的二叉树中查找目标值
func (this *FindElements) Find(target int) bool {
	// 在哈希集合中查找目标值
	return this.hsMap[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
