package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 类似于 2-Add-Two-Numbers，但是链表是正序的，所以需要先翻转链表，再相加
// ========================================  翻转链表方法  ========================================
// 1.我的解法，翻转数组，按位相加，从后到前赋值给新数组
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 翻转两个链表，得到翻转后的两个链表头，代表了两个数的个位数
	head1 := reverseLink(l1)
	head2 := reverseLink(l2)

	// 新链表头，为nil
	var newHead *ListNode
	// 进位符
	add := 0

	// 如果有链表不为空，或者有进位符时，继续循环
	for head1 != nil || head2 != nil || add != 0 {
		// 设定结果链表的新节点值为val，初始化为0
		val := 0

		// 首先加上进位符
		val += add

		// 如果链表1节点不为空，加上链表1节点的值
		if head1 != nil {
			val += head1.Val
			head1 = head1.Next
		}

		// 如果链表2节点不为空，加上链表2节点的值
		if head2 != nil {
			val += head2.Val
			head2 = head2.Next
		}

		// 如果val大于等于10，说明需要进位，val减去10，add加1
		if val >= 10 {
			val -= 10
			add = 1
		} else {
			// 否则，进位符为0，val不变
			add = 0
		}

		// 定义一个新节点，值为val，Next指向新链表头
		newNode := &ListNode{
			Val:  val,
			Next: newHead,
		}
		// 新链表头变为新节点
		newHead = newNode
	}

	// 返回新链表头
	return newHead
}

// ========================================  栈方法  ========================================
// 2.栈方法
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义两个栈
	stack1, stack2 := newStack(), newStack()

	// 将两个链表的值压入栈
	for l1 != nil {
		stack1.push(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2.push(l2.Val)
		l2 = l2.Next
	}

	// 新链表头
	var newHead *ListNode
	// 进位符
	add := 0
	// 如果两个栈不都为空，或者有进位符时，继续循环
	for stack1.head != nil || stack2.head != nil || add != 0 {
		// 设定结果链表的新节点值为val，初始化为0
		val := 0
		// 首先加上进位符
		val += add

		// 如果栈1不为空，弹出栈顶元素，加到val上
		if stack1.head != nil {
			val += stack1.pop().Val
		}
		// 如果栈2不为空，弹出栈顶元素，加到val上
		if stack2.head != nil {
			val += stack2.pop().Val
		}

		// 如果val大于等于10，说明需要进位，val减去10，add加1，否则进位符为0，val不变
		if val >= 10 {
			val = val - 10
			add = 1
		} else {
			add = 0
		}

		// 定义一个新节点，值为val，Next指向新链表头
		newNode := &ListNode{
			Val:  val,
			Next: newHead,
		}
		// 新链表头变为新节点
		newHead = newNode
	}
	// 返回新链表头
	return newHead
}

// ========================================  栈定义和方法  ========================================
type Stack struct {
	head *ListNode
}

// 栈的构造函数
func newStack() *Stack {
	return &Stack{
		head: nil,
	}
}

// 栈的push方法
func (stack *Stack) push(value int) {
	var newNode = &ListNode{
		Val:  value,
		Next: stack.head,
	}
	stack.head = newNode
}

// 栈的pop方法
func (stack *Stack) pop() *ListNode {
	if stack.head == nil {
		return nil
	} else {
		res := stack.head
		stack.head = stack.head.Next
		return res
	}
}

// 栈的peek方法
func (stack *Stack) peek() *ListNode {
	return stack.head
}

// ========================================  链表方法  ========================================
// 反转链表  方法见 （206-Reverse-Linked-List）
func reverseLink(head *ListNode) *ListNode {
	// 新链表头
	var newHead *ListNode

	// 遍历链表，将每个节点的Next指向新链表头，新链表头变为当前节点
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	// 返回新链表头
	return newHead
}

// 打印链表为数组
func printLink(head *ListNode) {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Println(res)
}
