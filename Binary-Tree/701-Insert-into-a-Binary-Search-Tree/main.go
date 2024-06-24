package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
