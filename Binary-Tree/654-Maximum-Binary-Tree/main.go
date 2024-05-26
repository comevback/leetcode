package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var max, index int
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}

	var node *TreeNode = &TreeNode{
		Val: max,
	}

	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])

	return node
}
