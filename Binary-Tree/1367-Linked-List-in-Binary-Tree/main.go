package main

import "fmt"

func main() {
	// 构建链表
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}

	// 构建二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 1}

	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Left.Right.Left = &TreeNode{Val: 1}
	root.Right.Left.Right.Right = &TreeNode{Val: 3}

	res := isSubPath1(head, root)
	fmt.Println(res) // 输出结果
}

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ************************************************  一般解法  *****************************************************

// isSubPath 检查链表是否是二叉树的子路径
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	// 如果当前节点值相同，检查是否存在匹配路径
	if head.Val == root.Val {
		if check(root, head) {
			return true
		}
	}

	// 递归检查左子树和右子树
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// check 检查从当前二叉树节点开始是否存在匹配链表的路径
func check(root *TreeNode, head *ListNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	// 如果当前节点值相同，递归检查左子树和右子树
	if root.Val == head.Val {
		return check(root.Left, head.Next) || check(root.Right, head.Next)
	}

	return false
}

// ************************************************  KMP 解法  *****************************************************

// 全局变量存储所有从根到叶子的路径
var Arrs [][]int

// isSubPath1 使用 KMP 算法检查链表是否是二叉树的子路径
func isSubPath1(head *ListNode, root *TreeNode) bool {
	Arrs = [][]int{}
	travelArr(root, []int{}) // 获取所有路径
	listArr := []int{}
	current := head
	for current != nil {
		listArr = append(listArr, current.Val)
		current = current.Next
	}

	// 使用 KMP 算法检查链表数组是否是任何一个路径的子数组
	for _, v := range Arrs {
		if KMP(v, listArr) != -1 {
			return true
		}
	}
	return false
}

// travelArr 获取二叉树中所有从根到叶子的路径
func travelArr(root *TreeNode, arr []int) {
	if root == nil {
		return
	}

	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		Arrs = append(Arrs, copyArr)
		return
	}

	travelArr(root.Left, arr)
	travelArr(root.Right, arr)
}

// KMP 使用 KMP 算法查找 needle 在 arr 中的位置
func KMP(arr []int, needle []int) int {
	LPSlist := LPS(needle)
	start := 0

	for i := 0; i < len(arr); i++ {
		if needle[start] == arr[i] {
			if start == len(needle)-1 {
				return i - start
			}
			start += 1
		} else {
			for start > 0 && needle[start] != arr[i] {
				start = LPSlist[start-1]
			}

			if needle[start] == arr[i] {
				start += 1
			}
		}
	}

	return -1
}

// LPS 计算 needle 的最长前缀后缀数组
func LPS(needle []int) []int {
	prefixList := make([]int, len(needle))
	length := 0

	for i := 1; i < len(needle); i++ {
		for length > 0 && needle[length] != needle[i] {
			length = prefixList[length-1]
		}

		if needle[length] == needle[i] {
			length += 1
			prefixList[i] = length
		}
	}

	return prefixList
}
