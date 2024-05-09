package main

import (
	"fmt"
	"strconv"
)

func main() {
	operations := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints1(operations))
}

// 1. ====================================================  数组处理  ====================================================
func calPoints(operations []string) int {
	// 创建一个数组
	arr := []int{}
	// 遍历输入
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			// 如果是+，则将数组最后两个元素相加
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
		case "D":
			// 如果是D，则将数组最后一个元素的两倍加入数组
			arr = append(arr, 2*arr[len(arr)-1])
		case "C":
			// 如果是C，则删除数组最后一个元素
			arr = arr[:len(arr)-1]
		default:
			// 如果是数字，则将该数字加入数组
			num, err := strconv.Atoi(operations[i])
			if err != nil {
				panic(err)
			}
			arr = append(arr, num)
		}
	}

	// 计算数组的和
	sum := 0
	for _, v := range arr {
		sum += v
	}
	// 返回和
	return sum
}

// 2. ====================================================  栈方法  ====================================================
func calPoints1(operations []string) int {
	// 创建一个栈
	stack := NewStack()
	// 遍历输入
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			// 如果是+，则将栈顶两个元素弹出，相加，再把弹出的两个元素和相加的新值推入栈
			num1, num2 := stack.pop(), stack.pop()
			stack.push(num2)
			stack.push(num1)
			stack.push(num1 + num2)
		case "D":
			// 如果是D，查看栈顶元素，将栈顶元素的两倍推入栈
			num := stack.peek()
			stack.push(num * 2)
		case "C":
			// 如果是C，将栈顶元素弹出
			stack.pop()
		default:
			// 如果是数字，则将该数字推入栈
			num, err := strconv.Atoi(operations[i])
			if err != nil {
				panic(err)
			}
			stack.push(num)
		}
	}
	// 计算栈中元素的和
	sum := 0
	for stack.head != nil {
		sum += stack.pop()
	}
	// 返回和
	return sum
}

// ====================================================  栈定义和方法  ====================================================
// 定义链表节点
type ListNode struct {
	Val  int
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
func (stack *Stack) push(value int) {
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}
	stack.head = newNode
}

// 栈的pop方法
func (stack *Stack) pop() int {
	if stack.head == nil {
		panic("empty stack")
	} else {
		res := stack.head.Val
		stack.head = stack.head.Next
		return res
	}
}

// 栈的peek方法
func (stack *Stack) peek() int {
	if stack.head != nil {
		return stack.head.Val
	} else {
		panic("empty stack")
	}
}
