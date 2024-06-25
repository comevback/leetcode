package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}

	res := build(1, n)
	return res
}

func build(lo int, hi int) []*TreeNode {
	res := []*TreeNode{}
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		leftTrees := build(lo, i-1)
		rightTrees := build(i+1, hi)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}

	return res
}
