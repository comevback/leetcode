/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	var l2 *ListNode = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}

// ================================================================================================================================
// 1. 我的解法：
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位标志add
	add := false
	// 设一个新的链表头 res，一个当前节点 currentNode先指向 res
	var res *ListNode = &ListNode{Val: 0, Next: nil}
	currentNode := res

	// 当 l1 和 l2 都不为空时，循环
	for l1 != nil && l2 != nil {
		// 如果没有进位
		if !add {
			// 如果 l1.Val+l2.Val >= 10，当前节点的值为 l1.Val+l2.Val-10，add = true
			if l1.Val+l2.Val >= 10 {
				currentNode.Val = l1.Val + l2.Val - 10
				add = true
			} else {
				// 如果 l1.Val+l2.Val < 10，当前节点的值为 l1.Val+l2.Val，add = false
				currentNode.Val = l1.Val + l2.Val
			}
			// 如果有进位
		} else {
			// 如果 l1.Val+l2.Val+1 >= 10，当前节点的值为 l1.Val+l2.Val+1-10，add = true
			if l1.Val+l2.Val+1 >= 10 {
				currentNode.Val = l1.Val + l2.Val + 1 - 10
			} else {
				// 如果 l1.Val+l2.Val+1 < 10，当前节点的值为 l1.Val+l2.Val+1，add = false
				currentNode.Val = l1.Val + l2.Val + 1
				add = false
			}
		}
		// l1 和 l2 向后移动
		l1 = l1.Next
		l2 = l2.Next
		// 如果 l1 == nil && l2 == nil && !add，跳出循环，没有下一个节点了
		if l1 == nil && l2 == nil && !add {
			break
		}
		// 如果 l1 和 l2 都不为空，新建一个节点，当前节点的 Next 指向新建的节点，当前节点指向新建的节点
		newListNode := &ListNode{Val: 0, Next: nil}
		currentNode.Next = newListNode
		currentNode = newListNode
	}

	// 当 l1 不为空时，循环，把 l1 的剩余节点加到 res 上
	for l1 != nil {
		if !add {
			currentNode.Val = l1.Val
			currentNode.Next = l1.Next
			break
		} else {
			if l1.Val+1 == 10 {
				currentNode.Val = 0
				newListNode := &ListNode{Val: 0, Next: nil}
				currentNode.Next = newListNode
				currentNode = newListNode
			} else {
				currentNode.Val = l1.Val + 1
				currentNode.Next = l1.Next
				add = false
				break
			}
		}
		l1 = l1.Next
	}

	// 当 l2 不为空时，循环，把 l2 的剩余节点加到 res 上
	for l2 != nil {
		if !add {
			currentNode.Val = l2.Val
			currentNode.Next = l2.Next
			break
		} else {
			if l2.Val+1 == 10 {
				currentNode.Val = 0
				newListNode := &ListNode{Val: 0, Next: nil}
				currentNode.Next = newListNode
				currentNode = newListNode
			} else {
				currentNode.Val = l2.Val + 1
				currentNode.Next = l2.Next
				add = false
				break
			}
		}
		l2 = l2.Next
	}

	// 如果 l1 和 l2 都为空，但是有进位，新建一个节点，当前节点的 Next 指向新建的节点，当前节点指向新建的节点
	if add {
		// 如果currentNode.Val+1 == 10，当前节点的值为 0，新建一个节点，当前节点的 Next 指向新建的节点，当前节点指向新建的节点
		if currentNode.Val+1 == 10 {
			currentNode.Val = 0
			newListNode := &ListNode{Val: 1, Next: nil}
			currentNode.Next = newListNode
			currentNode = newListNode
		} else {
			// 如果currentNode.Val+1 < 10，当前节点的值为 currentNode.Val+1，就是最后一个节点
			currentNode.Val += 1
		}
	} else {
		// 如果 l1 和 l2 都为空，没有进位，删除最后一个节点
		currentNode = nil
	}

	// 返回 res
	return res
}

// ================================================================================================================================
// 2.语法简洁版：
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 定义一个伪头节点，因为这个不是真的要返回的头节点
	var head *ListNode = &ListNode{}

	// 定义current 为这个头节点，在循环中，current不断变换为current.Next，也就是下一个节点
	current := head

	// 是否进位，如果0不进位，如果1就是进位
	addOne := 0

	// 循环条件，如果有l1，l2有一个不为nil，或者还有进位，则继续循环
	for l1 != nil || l2 != nil || addOne != 0 {
		// 首先新的一个节点的值定义为val，看有没有进位，有进位先加上
		val := addOne

		// 如果l1不为nil，val加上l1.Val，l1变成下一个节点
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		// 如果l2不为nil，val加上l2.Val，l2变成下一个节点
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		// 如果val大于等于10，就进位，新节点的赋值是val-10，addOne变为1，否则新节点赋值val，addOne变为0
		// 也可以定义addOne = val/10
		if val >= 10 {
			val = val - 10
			addOne = 1
		} else {
			addOne = 0
		}

		// current.Next就是我们这一次循环要得到的节点，只要给Val赋值val，不用管Next，Next在下一次循环中定义。
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	// 最后返回伪头节点的下一个节点，也就是真正的头节点
	return head.Next
}
