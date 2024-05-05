# 2130. Maximum Twin Sum of a Linked List（链表局部翻转）

> 如果要求前后对应的值，可以考虑翻转链表。但不能全部翻转，因为会改变整个链表的结构，所以只翻转后半部分。翻转方法见（206-Reverse-Linked-List）。

Medium

Hint
In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.

Example 1:
![eg1drawio](https://assets.leetcode.com/uploads/2021/12/03/eg1drawio.png)
> Input: head = [5,4,2,1]
Output: 6
Explanation:
Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
There are no other nodes with twins in the linked list.
Thus, the maximum twin sum of the linked list is 6. 

Example 2:
![eg2drawio](https://assets.leetcode.com/uploads/2021/12/03/eg2drawio.png)
> Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7. 

Example 3:
![eg3drawio](https://assets.leetcode.com/uploads/2021/12/03/eg3drawio.png)
> Input: head = [1,100000]
Output: 100001
Explanation:
There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.
 

Constraints:
> The number of nodes in the list is an even integer in the range [2, 105].
1 <= Node.val <= 105

---

## 思路
1. 用数组或map来存储链表的值，然后遍历数组或map，找到对应值相加最大的值。
2. 翻转链表，但是只翻转后半部分，然后同时遍历原链表的前半部分和翻转后的链表，找到对应值相加最大的值。翻转链表的方法见（206-Reverse-Linked-List）。

## 代码
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var head *ListNode = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	// pairSum(head)
	res := pairSum1(head)
	println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 定义数组 O(n)-O(n)
func pairSum(head *ListNode) int {
	// 定义一个数组
	var arr []int

	// 把链表里的每个值按顺序加入到数组中
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	// 定义最大的twin sum为maxSum
	var maxSum int = math.MinInt32
	// 遍历数组，找配对的两个值加起来，如果大于maxSum，则替换为maxSum
	for i := 0; i < len(arr); i++ {
		if arr[i]+arr[len(arr)-1-i] > maxSum {
			maxSum = arr[i] + arr[len(arr)-1-i]
		}
	}

	// 返回最大值maxSum
	return maxSum
}

// ========================================================================================================================

// 2. 翻转链表（后半部分）（双指针）O(n)-O(1)
func pairSum1(head *ListNode) int {

	// 定义两个指针，一个快指针，一个慢指针
	var slow, fast *ListNode = head, head.Next

	// 快指针走两步，慢指针走一步，当快指针走到最后一个元素时，慢指针走到中间元素，此时慢指针的下一个元素就是后半部分的头
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转后半部分链表，返回翻转后的头
	reversedHead := reverseList(slow.Next) // 用两种翻转方法都可以

	// 定义最大的twin sum为maxSum
	var maxSum int = math.MinInt32

	// head前半部分正常，后半部分已经被翻转
	// reverseHead所代表的被翻转的后半部分，长度为原链表一半，所以用reverseHead != nil来判断是否遍历完
	for reversedHead != nil {
		// 如果head.Val+reversedHead.Val大于maxSum，则替换为maxSum
		if head.Val+reversedHead.Val > maxSum {
			maxSum = head.Val + reversedHead.Val
		}
		// head和reversedHead分别往后移动
		head = head.Next
		reversedHead = reversedHead.Next
	}

	// 返回最大值maxSum
	return maxSum
}

// ========================================================================================================================
// Functions

// 翻转链表1
// 详细解释见(206-Reverse-Linked-List)
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	return newHead
}

// 翻转链表2
// 详细解释见(206-Reverse-Linked-List)
func reverseList2(head *ListNode) *ListNode {

	var arr []int

	temp1 := head
	temp2 := head

	for temp1 != nil {
		arr = append(arr, temp1.Val)
		temp1 = temp1.Next
	}

	for i := 0; i < len(arr); i++ {
		temp2.Val = arr[len(arr)-1-i]
		temp2 = temp2.Next
	}

	return head
}

// 打印链表
func printLink(head *ListNode) {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	fmt.Println(arr)
}
```