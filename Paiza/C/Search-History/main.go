package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	useStack()
}

// =================================================  Stack Implementation  ====================================================

// 1.用Stack实现
func useStack() {
	// 读取输入，输入数字为搜索的次数
	var num int
	fmt.Scan(&num)
	// 创建一个栈
	stack := NewStack()

	// 创建一个扫描器，持续读取num次输入
	scanner := bufio.NewScanner(os.Stdin)
	// 读取输入
	for scanner.Scan() {
		// 每一行输入作为字符串
		str := scanner.Text()
		// 查找栈中是否有该字符串
		f := stack.find(str)
		// 如果有，删除该字符串
		if f != nil {
			stack.delete(f)
		}

		// 将该字符串压入栈
		stack.push(str)
		// 次数减一，如果次数小于1，退出循环
		num -= 1
		if num < 1 {
			break
		}
	}

	// 打印栈
	stack.print()
}

// 定义链表节点
type ListNode struct {
	Val  string
	Next *ListNode
}

// 定义栈
type Stack struct {
	head *ListNode
}

// 栈的构造函数
func NewStack() *Stack {
	return &Stack{}
}

// 栈的push方法
func (stack *Stack) push(value string) {
	// 创建新节点
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}
	// 新节点作为栈顶
	stack.head = newNode
}

// 栈的find方法，查找返回值为value的节点
func (stack *Stack) find(value string) *ListNode {
	// 创建一个指针指向栈顶
	current := stack.head

	// 遍历栈
	for current != nil {
		// 如果找到了值为value的节点，返回该节点
		if value == current.Val {
			return current
		}
		// 指针后移
		current = current.Next
	}

	// 没有找到返回nil
	return nil
}

// 栈的delete方法，删除节点
func (stack *Stack) delete(node *ListNode) {

	// 如果要删除的节点是栈顶节点，直接将栈顶指针指向下一个节点
	if node == stack.head {
		stack.head = stack.head.Next
		return
	}

	// 创建一个指针指向栈顶
	current := stack.head
	// 遍历栈
	for current.Next != nil {
		// 如果找到了要删除的节点，将该节点的前一个节点的Next指向该节点的Next
		if node == current.Next {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

// 栈的print方法，打印栈
func (stack *Stack) print() {
	current := stack.head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// =================================================  Array Implementation  ====================================================
// 2.
