package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var biggest int

func getMinimumDifference(root *TreeNode) int {
	biggest = 0

	if root == nil {
		return 0
	}

	left := getMinimumDifference(root.Left)
	right := getMinimumDifference(root.Right)

	diff := abs(left - right)
	if diff > biggest {

	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
