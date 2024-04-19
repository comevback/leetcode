package main

import "fmt"

func main() {

	// -------------------------------
	// 测试用数据

	l3 := ListNode{
		Val: 3,
	}

	l2 := ListNode{
		Val:  4,
		Next: &l3,
	}

	l1 := ListNode{
		Val:  2,
		Next: &l2,
	}

	l6 := ListNode{
		Val: 4,
	}

	l5 := ListNode{
		Val:  6,
		Next: &l6,
	}

	l4 := ListNode{
		Val:  5,
		Next: &l5,
	}

	res := addTwoNumbers(&l1, &l4)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	// 创建最后一个节点
// 	tail := &ListNode{Val: 0}
// 	// 是否进位变量
// 	addOne := 0

// 	for {

// 		// 这一次循环里Val要考虑的值，大于等于10则取 added - 10，进位，小于10则取 added。
// 		added := l1.Val + l2.Val + addOne

// 		// 根据added 是否大于等于10，是否进位，而选择这个尾节点的值
// 		if added >= 10 {
// 			current := &ListNode{
// 				Val:  added - 10,
// 				Next: tail,
// 			}
// 			addOne = 1
// 			tail = current
// 		} else {
// 			current := &ListNode{
// 				Val:  added,
// 				Next: tail,
// 			}
// 			addOne = 0
// 			tail = current
// 		}

// 		// 只有两个链表都结束时，才退出循环，如果最后一位有一个是9，并且之前进位了，那再在尾部添加一个{Val: 1}的尾节点，否则返回最后的尾节点。
// 		if l1.Next == nil && l2.Next == nil {
// 			if added == 10 {
// 				current := &ListNode{
// 					Val: added,
// 					Next: &ListNode{
// 						Val:  1,
// 						Next: tail,
// 					},
// 				}
// 				res := &ListNode{
// 					Val:  1,
// 					Next: current,
// 				}
// 				return res

// 			} else {
// 				current := &ListNode{
// 					Val:  added,
// 					Next: tail,
// 				}
// 				addOne = 0
// 				return current
// 			}
// 			// 如果其中有一个链表后面已经没有节点，则给它下一个节点赋值，让循环可以继续延续下去
// 		} else if l1.Next == nil {
// 			l1.Next = &ListNode{
// 				Val:  0,
// 				Next: nil,
// 			}
// 		} else if l2.Next == nil {
// 			l2.Next = &ListNode{
// 				Val:  0,
// 				Next: nil,
// 			}
// 		}

// 		l1 = l1.Next
// 		l2 = l2.Next
// 	}
// }

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
