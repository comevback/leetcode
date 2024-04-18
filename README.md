# 2. Add Two Numbers

Medium

> You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 

Example 1:
> Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

> Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

> Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
 

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

# My First Code

思路：把链表转换成整数，整数相加，再把整数转换为一个顺序相反的数组，用这个数组生产一个链表。

错误点：如果链表太长，就会整数溢出，得到错的答案，应该直接在链表中位相加，

```go
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

	//这种转换成整数，相加，再转换为数组，再转换为链表的方式， 在整数相加时，如果遇到极大的数字，会溢出，所以不可取。
	//-------------------------------

	var arr1 []int = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	var arr2 []int = []int{5, 6, 4}

	node1 := getNode(arr1)
	node2 := getNode(arr2)

	num1 := getNumber(node1)
	num2 := getNumber(node2)

	fmt.Println(num1)
	fmt.Println(num2)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := getNumber(l1)
	num2 := getNumber(l2)

	sum := num1 + num2
	fmt.Println(sum)
	arr := getList(sum)
	fmt.Println(arr)
	res := getNode(arr)
	return res
}

// 输入链表头，转换为整数
func getNumber(l *ListNode) int {

	//递归，如果Next不为nil，说明不是最后一个数，下一个数比这个数更大一位，所以要乘10。 如果为nil，说明是最后一个数，直接返回。
	if l.Next != nil {
		return l.Val + 10*getNumber(l.Next)
	} else {
		return l.Val
	}
}

// 输入整数，转换为这个数的个位数
func getLastNum(i int) int {
	if i >= 10 {
		return i % 10
	} else {
		return i
	}
}

// 输入整数，转换为一个切片，切片的顺序是和整数每一位上相反的， 例如 243 => [3, 4, 2]
func getList(num int) []int {
	var res []int = make([]int, 0)

	//如果这个数是0，直接返回[0]
	if num == 0 {
		return []int{0}
	}

	//循环，只要不小于1，把这个数的个位数填进数组，然后减掉个卫视，除以10
	for i := num; i >= 1; i = (i - getLastNum(i)) / 10 {
		res = append(res, getLastNum(i))
	}

	//得到数组
	return res
}

// 输入一个整数类型的切片， 转换为一个链表，
func getNode(li []int) *ListNode {

	//创建一个链表中的元素，类型为指针，值为nil
	var node *ListNode = nil

	//从切片最后一位开始向前遍历
	for i := len(li) - 1; i >= 0; i-- {
		//这个current是一个链表中的元素，值是数组的最后一个，Next也是指向nil，所以指向之前定义的node
		current := &ListNode{
			Val:  li[i],
			Next: node,
		}
		//此时把node变为刚才定义的current
		node = current
	}

	//最后得到的node就是最新生产的一个元素，也就是链表的第一个元素，就是链表头
	return node
}

```
