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

// **************************************************************************************************************
// review 7.5
func buildTree_75(inorder []int, postorder []int) *TreeNode {
	return build_75(inorder, postorder)
}

func build_75(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	newNode := &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	leftPostOrder := []int{}
	leftInOrder := []int{}
	rightPostOrder := []int{}
	rightInOrder := []int{}

	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			leftInOrder = inorder[:i]
			if i+1 < len(inorder) {
				rightInOrder = inorder[i+1:]
			}
		}
	}

	rightPostOrder = postorder[len(postorder)-1-len(rightInOrder) : len(postorder)-1]
	leftPostOrder = postorder[:len(postorder)-1-len(rightInOrder)]

	newNode.Left = build_75(leftInOrder, leftPostOrder)
	newNode.Right = build_75(rightInOrder, rightPostOrder)

	return newNode
}
