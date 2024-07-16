package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	node  *TreeNode
	arr   []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	arr := []int{}
	travel(root, &arr)
	return BSTIterator{
		node:  root,
		arr:   arr,
		index: 0,
	}
}

func (this *BSTIterator) Next() int {
	res := this.arr[this.index]
	this.index += 1
	return res
}

func (this *BSTIterator) HasNext() bool {
	if this.index < len(this.arr) {
		return true
	} else {
		return false
	}
}

func travel(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	travel(root.Left, arr)
	(*arr) = append((*arr), root.Val)
	travel(root.Right, arr)
}
