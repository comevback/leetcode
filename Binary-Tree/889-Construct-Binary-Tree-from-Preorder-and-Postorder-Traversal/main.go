package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(preorder) == 0 {
		return nil
	}

	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	lefthead := preorder[1]

	var divide int
	for i := 1; i < len(postorder); i++ {
		if postorder[i] == lefthead {
			divide = i
			break
		}
	}

	leftPostOrder := postorder[:divide+1]
	rightPostOrder := postorder[divide+1 : len(postorder)-1]

	leftPreOrder := preorder[1 : 1+len(leftPostOrder)]
	rightPreOrder := preorder[1+len(leftPostOrder):]

	root.Left = constructFromPrePost(leftPreOrder, leftPostOrder)
	root.Right = constructFromPrePost(rightPreOrder, rightPostOrder)

	return root
}
