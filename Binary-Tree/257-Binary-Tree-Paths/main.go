package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	res := binaryTreePaths_628(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	// 遍历思路
	// paths = []string{}
	// if root == nil {
	// 	return []string{}
	// }

	// collectPath(root, "")
	// return paths

	// 分解思路
	res := getPath(root)
	return res
}

// 1. 遍历思路
func collectPath(root *TreeNode, path string) {
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	}

	path += "->"

	if root.Left != nil {
		collectPath(root.Left, path)
	}

	if root.Right != nil {
		collectPath(root.Right, path)
	}
}

// 2. 分解思路
func getPath(root *TreeNode) []string {
	paths := []string{}

	if root.Left == nil && root.Right == nil {
		str := strconv.Itoa(root.Val)
		paths = append(paths, str)
		return paths
	}

	paths = append(paths, getPath(root.Left)...)
	paths = append(paths, getPath(root.Right)...)

	for i := range paths {
		paths[i] = strconv.Itoa(root.Val) + "->" + paths[i]
	}

	return paths
}

// ******************************************  review 6.28  **********************************************
var paths_628 []string

func binaryTreePaths_628(root *TreeNode) []string {
	paths_628 = []string{}
	genePath(root, "")
	return paths_628
}

func genePath(root *TreeNode, path string) {
	if root == nil {
		return
	}
	valStr := strconv.Itoa(root.Val)
	path += valStr

	if root.Left == nil && root.Right == nil {
		paths_628 = append(paths_628, path)
		return
	}

	path += "->"

	if root.Left != nil {
		genePath(root.Left, path)
	}

	if root.Right != nil {
		genePath(root.Right, path)
	}
}
