# 160. Intersection of Two Linked Lists

Easy

Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

For example, the following two linked lists begin to intersect at node c1:


The test cases are generated such that there are no cycles anywhere in the entire linked structure.

Note that the linked lists must retain their original structure after the function returns.

Custom Judge:

The inputs to the judge are given as follows (your program is not given these inputs):

intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
listA - The first linked list.
listB - The second linked list.
skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program. If you correctly return the intersected node, then your solution will be accepted.

 

Example 1:
![alt text](https://assets.leetcode.com/uploads/2021/03/05/160_statement.png)
> Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
- Note that the intersected node's value is not 1 because the nodes with value 1 in A and B (2nd node in A and 3rd node in B) are different node references. In other words, they point to two different locations in memory, while the nodes with value 8 in A and B (3rd node in A and 4th node in B) point to the same location in memory.

Example 2:
![alt text](https://assets.leetcode.com/uploads/2021/03/05/160_example_1_1.png)
> Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.

Example 3:
![alt text](https://assets.leetcode.com/uploads/2021/03/05/160_example_3.png)
> Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
 

Constraints:

The number of nodes of listA is in the m.
The number of nodes of listB is in the n.
1 <= m, n <= 3 * 104
1 <= Node.val <= 105
0 <= skipA < m
0 <= skipB < n
intersectVal is 0 if listA and listB do not intersect.
intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.
 

Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?

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

// 1. 互相连接头尾
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 初始化两个指针，Pa 和 Pb，分别指向两个链表的头部。
	Pa, Pb := headA, headB

	// 继续遍历，直到两个指针相遇。
	for Pa != Pb {
		// 如果指针 Pa 到达链表 A 的末尾，则将其重置为链表 B 的头部。
		// 如果不为空，继续移动到下一个节点。
		if Pa == nil {
			Pa = headB
		} else {
			Pa = Pa.Next
		}

		// 如果指针 Pb 到达链表 B 的末尾，则将其重置为链表 A 的头部。
		// 如果不为空，继续移动到下一个节点。
		if Pb == nil {
			Pb = headA
		} else {
			Pb = Pb.Next
		}
	}

	// 返回相遇点，如果两链表相交，此时 Pa 和 Pb 指向相交节点；如果不相交，Pa 和 Pb 会同时为 nil。
	return Pa
}

// 2.我的方法，较为复杂
// getIntersectionNode 函数用于找出两个链表相交的第一个节点。
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	// 分别初始化两个指针 Pa 和 Pb，开始时分别指向两个链表的头节点。
	Pa, Pb := headA, headB
	// continueA 和 continueB 用于标记每个链表是否已经遍历到尾部并切换到另一个链表头部。
	continueA, continueB := true, true

	// 使用一个无限循环，直到找到相交节点或确定不相交。
	for {
		// 如果两个指针相遇，返回该相遇节点，该节点即为两链表的第一个公共节点。
		if Pa == Pb {
			return Pa
		}

		// 移动指针 Pa。
		Pa = Pa.Next
		// 如果 Pa 到达链表末尾，则根据 continueA 的值决定是否跳到链表 B 的头部开始。
		if Pa == nil {
			if continueA {
				Pa = headB
				continueA = false // 确保只切换一次
			} else {
				// 如果 Pa 已经在链表 B 上遍历过一遍后仍为空，则两链表不相交，返回 nil。
				return nil
			}
		}

		// 移动指针 Pb。
		Pb = Pb.Next
		// 如果 Pb 到达链表末尾，则根据 continueB 的值决定是否跳到链表 A 的头部开始。
		if Pb == nil {
			if continueB {
				Pb = headA
				continueB = false // 确保只切换一次
			} else {
				// 如果 Pb 已经在链表 A 上遍历过一遍后仍为空，则两链表不相交，返回 nil。
				return nil
			}
		}
	}
}
```