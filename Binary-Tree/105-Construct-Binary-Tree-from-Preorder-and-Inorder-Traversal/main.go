package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果给的遍历数组长度为0，说明是空，返回nil
	if len(preorder) == 0 {
		return nil
	}

	// 这个子树的根节点一定是前序遍历数组的第一个，把这个根节点定义出来
	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}

	// devide定义为在中序遍历数组中分割左右子树的index，也就是前序遍历数组中头元素在中序遍历数组中的index
	var devide int

	// 找到这个divide
	for i, v := range inorder {
		if v == preorder[0] {
			devide = i
			break
		}
	}

	// 找到左右子树的中序遍历数组
	var leftInOrder, rightInOrder []int

	leftInOrder = inorder[:devide]
	if devide < len(inorder)-1 {
		rightInOrder = inorder[devide+1:]
	}

	// 找到左右子树的前序遍历数组
	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	rightPreOrder := preorder[1+len(leftInOrder):]

	// 把左右子树通过递归得到，附到这个节点上
	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)

	// 返回这个节点
	return root
}
