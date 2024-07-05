package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	subRoot := &TreeNode{Val: 1}
	subRoot.Left = &TreeNode{Val: 2}

	res := isSubtree(root, subRoot)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 比较
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if same(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// 针对单个节点，看以这个节点为根的子树是否与subRoot完全相等
func same(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return same(root.Left, subRoot.Left) && same(root.Right, subRoot.Right)
}

// **************************************************  KMP无法解决  ******************************************************
// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	arr1 := []int{}
// 	travel(root, &arr1)

// 	arr2 := []int{}
// 	travel(subRoot, &arr2)

// 	revArr1 := []int{}
// 	travelReverse(root, &revArr1)

// 	revArr2 := []int{}
// 	travelReverse(subRoot, &revArr2)

// 	inOrderArr1 := []int{}
// 	travelInOrder(root, &inOrderArr1)

// 	inOrderArr2 := []int{}
// 	travelInOrder(subRoot, &inOrderArr2)

// 	res1 := KMP(arr1, arr2)
// 	res2 := KMP(revArr1, revArr2)
// 	res3 := KMP(inOrderArr1, inOrderArr2)
// 	if res1 == -1 || res2 == -1 || res3 == -1 {
// 		return false
// 	}

// 	return true
// }

// func travel(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	travel(root.Left, arr)
// 	travel(root.Right, arr)
// 	*arr = append(*arr, root.Val)
// }

// func travelInOrder(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	travelInOrder(root.Left, arr)
// 	*arr = append(*arr, root.Val)
// 	travelInOrder(root.Right, arr)
// }

// func travelReverse(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}

// 	travelReverse(root.Right, arr)
// 	travelReverse(root.Left, arr)
// 	*arr = append(*arr, root.Val)
// }

// func KMP(arr []int, needle []int) int {
// 	LPSlist := LPS(needle)
// 	posi := 0

// 	for i := 0; i < len(arr); i++ {
// 		if needle[posi] == arr[i] {
// 			if posi == len(LPSlist)-1 {
// 				return i - posi
// 			}
// 			posi += 1
// 		} else {
// 			for posi > 0 && needle[posi] != arr[i] {
// 				posi = LPSlist[posi-1]
// 			}

// 			if needle[posi] == arr[i] {
// 				posi += 1
// 			}
// 		}
// 	}

// 	return -1
// }

// func LPS(needle []int) []int {
// 	prefixList := make([]int, len(needle))
// 	length := 0

// 	for i := 1; i < len(needle); i++ {
// 		if length > 0 && prefixList[length] != needle[i] {
// 			for length > 0 && prefixList[length] != needle[i] {
// 				length = prefixList[length-1]
// 			}

// 			if prefixList[length] == needle[i] {
// 				length += 1
// 				prefixList[i] = length
// 			}
// 		}
// 	}

// 	return prefixList
// }
