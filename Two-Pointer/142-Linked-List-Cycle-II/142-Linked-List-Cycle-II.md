# 142. Linked List Cycle II

Medium

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.


Example 1:
![circularlinkedlist](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)
> Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
![circularlinkedlist_test2](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)
> Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
![circularlinkedlist_test3](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
> Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
 

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.
 

Follow up: Can you solve it using O(1) (i.e. constant) memory?

---

# Code
```go
package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 函数用于检测链表中的环，并返回环的起始节点。
func detectCycle(head *ListNode) *ListNode {
	// 使用两个指针，slow 和 fast，初始都指向头节点。
	slow, fast := head, head

	// 快指针每次移动两步，慢指针每次移动一步。
	for fast != nil {
		fast = fast.Next // 快指针先移动一步
		if fast != nil {
			fast = fast.Next // 如果快指针还未到末尾，则继续移动一步
		}
		slow = slow.Next // 慢指针移动一步

		// 如果快慢指针相遇，说明链表中存在环。
		if fast == slow {
			break
		}
	}

	// 如果快指针走到了链表末尾，说明链表中没有环，返回 nil。
	if fast == nil {
		return nil
	}

	// 如果存在环，需要找到环的起始节点。
	// 将一个指针重新指向头节点，另一个指针保持在相遇点，
	// 然后两个指针都以相同速度移动，当它们再次相遇时，
	// 相遇点就是环的起始节点。
	slow = head
	for fast != slow {
		fast = fast.Next // 快指针移动一步
		slow = slow.Next // 慢指针移动一步

		// 如果两者再次相遇，跳出循环。
		if fast == slow {
			break
		}
	}

	// 返回环的起始节点。
	return fast
}
```