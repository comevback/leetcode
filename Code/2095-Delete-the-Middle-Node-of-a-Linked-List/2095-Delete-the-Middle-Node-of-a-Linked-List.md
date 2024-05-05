# 2095. Delete the Middle Node of a Linked List

Medium

Hint
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
 

Example 1:
![example1.png](example1.png)
> Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node. 

Example 2:
![example2.png](example2.png)
> Input: head = [1,2,3,4]
Output: [1,2,4]
Explanation:
The above figure represents the given linked list.
For n = 4, node 2 with value 3 is the middle node, which is marked in red.

Example 3:
![example3.png](example3.png)
> Input: head = [2,1]
Output: [2]
Explanation:
The above figure represents the given linked list.
For n = 2, node 1 with value 1 is the middle node, which is marked in red.
Node 0 with value 2 is the only node remaining after removing node 1.
 

Constraints:
> The number of nodes in the list is in the range [1, 105].
1 <= Node.val <= 105

---

## 思路
双指针，设置一快一慢两个指针，类型都是*ListNode，初始化都为head，当fast.Next不为nil时，继续循环
每次循环，快慢指针都向后走一步，fast = fast.Next，slow = slow.Next
如果在fast走一步后，fast.Next为nil，说明链表长度为奇数，此时slow指向的就是中间节点，删除slow即可
如果此时fast.Next不为nil，fast再走一步。

## 代码
```go
package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	printLink(node1)
	res := deleteMiddle(node1)
	printLink(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteMiddle(head *ListNode) *ListNode {

	// 如果只有一个节点，直接返回nil
	if head.Next == nil {
		return nil
	}

	// 定义快慢指针
	var slow, fast *ListNode = head, head

	// 每次循环，快慢指针都向后移动一步，如果快指针的下一个节点为空，说明慢指针指向的是中间节点
	// 如果快指针的下一个节点不为空，说明慢指针指向的是中间节点的前一个节点，快指针继续向后移动一步
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	// 当快指针无法再向后移动时，这时候慢指针指向的就是中间节点
	// 当中间节点的下一个节点不为空时，将中间节点替换为下一个节点，相当于删除了中间节点
	if slow.Next != nil {
		slow.Val = slow.Next.Val
		temp := slow.Next.Next
		slow.Next.Next = nil
		slow.Next = temp
	} else {
		// 当中间节点的下一个节点为空时，说明中间节点是最后一个节点，直接将中间节点置为nil
		head.Next = nil
	}

	// 返回链表的头节点
	return head
}

// 打印链表方法
func printLink(head *ListNode) {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	fmt.Println(arr)
}
```