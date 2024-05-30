# 82. Remove Duplicates from Sorted List II
Solved
Medium
Topics
Companies
Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.


Example 1:
![alt text](https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg)
> Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
![alt text](https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg)
> Input: head = [1,1,1,2,3]
Output: [2,3]


Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.

---

# Code
```go
package main

func main() {
	var head *ListNode = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}

	res := deleteDuplicates(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteDuplicates(head *ListNode) *ListNode {
	// 边界条件：如果链表为空或者只有一个节点，直接返回
	if head.Next == nil || head == nil {
		return head
	}
	// 边界条件：如果只有两个节点，判断两个节点是否相等，如果相等，返回nil，否则返回head
	if head.Next.Next == nil {
		if head.Val != head.Next.Val {
			return head
		} else {
			return nil
		}
	}

	// dummy节点，用于处理头节点就是重复节点的情况，赋值为101，因为链表的值范围是-100～100
	var dummy *ListNode = &ListNode{Val: 101}
	dummy.Next = head          // dummy节点指向head
	fast, slow := dummy, dummy // 快慢指针都指向dummy节点
	var ok *ListNode           // 用于记录下一个不重复的节点

	// 当fast.Next.Next不为空时，判断fast.Next和fast.Next.Next是否相等，如果不相等，说明fast.Next是不重复的节点，将slow.Next指向fast.Next，slow进一步
	for fast.Next.Next != nil {
		if fast.Next.Val != fast.Val && fast.Next.Next.Val != fast.Next.Val {
			ok = fast.Next
			slow.Next = ok
			slow = slow.Next
		}
		fast = fast.Next
	}

	// 最后一个节点在上面的循环中没有处理，所以在这里处理
	if fast.Next.Val != fast.Val {
		ok = fast.Next
		slow.Next = ok
		slow = slow.Next
	}

	// 最后要断开slow.Next，否则可能把后面重复的节点也带上
	slow.Next = nil

	return dummy.Next
}

// ********************************************************************************************************************
// 逻辑更清晰的双指针解法
func deleteDuplicates1(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	slow, fast := dummy, head
	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			// ** 一旦发现重复节点，跳过所有重复节点
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}
			fast = fast.Next //** 此时进入到一个不同的值的节点，但是这个节点是否重复还不知道，所以下一轮循环来判断
			// 此时 fast 跳过了这一段重复元素
			if fast == nil {
				slow.Next = nil
			}
			// 不过下一段元素也可能重复，等下一轮 for 循环判断
		} else {
			// 不是重复节点，接到 dummy 后面
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	return dummy.Next
}

// 递归解法
func deleteDuplicates2(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		// 如果头结点和身后节点的值不同，则对之后的链表去重即可
		head.Next = deleteDuplicates(head.Next)
		return head
	}
	// 如果如果头结点和身后节点的值相同，则说明从 head 开始存在若干重复节点
	// 越过重复节点，找到 head 之后那个不重复的节点
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	// 直接返回那个不重复节点开头的链表的去重结果，就把重复节点删掉了
	return deleteDuplicates(head.Next)
}
```