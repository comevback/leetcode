package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	var devide int
	for i := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			devide = i
			break
		}
	}

	rightInOrder := inorder[devide+1:]
	leftInOrder := inorder[:devide]

	rightPostOrder := postorder[len(leftInOrder) : len(postorder)-1]
	leftPostOrder := postorder[:len(leftInOrder)]

	root.Right = buildTree(rightInOrder, rightPostOrder)
	root.Left = buildTree(leftInOrder, leftPostOrder)

	return root
}
