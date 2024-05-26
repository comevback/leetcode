package main

import (
	"fmt"
	"strconv"
)

func main() {
	//root = [0,0,0,0,null,null,0,null,null,0,0]
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  0,
			Left: nil,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	res := findDuplicateSubtrees(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var existed map[string]int
var duplicate []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	existed = make(map[string]int)
	duplicate = []*TreeNode{}

	if root == nil {
		return nil
	}

	dfs(root)
	return duplicate

}

// 1. 我的解法，检查两个子节点
func dfs(root *TreeNode) string {
	if root == nil {
		return ""
	}

	leftStr := dfs(root.Left)
	rightStr := dfs(root.Right)

	// =========================left============================

	if existed[leftStr] == 1 {
		duplicate = append(duplicate, root.Left)
		existed[leftStr] += 1
	}

	if existed[leftStr] == 0 && root.Left != nil {
		existed[leftStr] = 1
	}

	// =========================right============================

	if existed[rightStr] == 1 {
		duplicate = append(duplicate, root.Right)
		existed[rightStr] += 1
	}

	if existed[rightStr] == 0 && root.Right != nil {
		existed[rightStr] = 1
	}

	res := ""
	res += strconv.Itoa(root.Val)
	res = res + "-left->" + leftStr
	res = res + "-right->" + rightStr

	return res
}

// 2. 简化版，只检查本节点，而不是两个子节点
func dfs1(root *TreeNode) string {
	if root == nil {
		return ""
	}

	leftStr := dfs(root.Left)
	rightStr := dfs(root.Right)

	res := ""
	res += strconv.Itoa(root.Val)
	res = res + "-left->" + leftStr
	res = res + "-right->" + rightStr

	if existed[res] == 1 {
		duplicate = append(duplicate, root)
		existed[res] += 1
	}

	if existed[res] == 0 {
		existed[res] = 1
	}

	return res
}
