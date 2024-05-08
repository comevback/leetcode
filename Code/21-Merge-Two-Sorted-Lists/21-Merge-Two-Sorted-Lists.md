# 21. Merge Two Sorted Lists

**本题与 [Merge-Sorted-Arrays.md](../../Memos/Merge-Sorted-Arrays.md) 类似** 

Easy

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

 

Example 1:
![merge_ex1](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)
> Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
> Input: list1 = [], list2 = []
Output: []

Example 3:
> Input: list1 = [], list2 = [0]
Output: [0]
 

Constraints:
> The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

---

# 思路
## 1. 取较小值
取两链表头中较小的一个，循环，知道其中一个链表为空，将另一个链表剩余部分接到结果链表后面

## 2. 递归
递归地比较两个链表的头节点，将较小的节点作为当前节点，然后递归地将较小节点的next指针指向下一个较小节点，直到其中一个链表为空，返回另一个链表

---

# 代码

```go
package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ============================================ 本题与 Merge-Sorted-Arrays.md 类似 ================================================

// =========================================================  取最小值  ===========================================================
// 1. 取两者最小值
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个新的链表头
	var newHead *ListNode = &ListNode{}
	// 创建一个current指针，指向newHead
	current := newHead
	// 遍历两个链表，将较小的节点接到current后面
	for list1 != nil && list2 != nil {
		// 如果list1.Val小，将list1接到current后面，current和list1都后移
		if list1.Val < list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			// 如果list2.Val小，将list2接到current后面，current和list2都后移
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}

	// 如果list1或list2有一个为空，直接将另一个链表接到current后面
	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	// 返回结果链表，也就是newHead的下一个节点
	return newHead.Next
}

// =========================================================  递归  ===========================================================
// 2. 递归
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	//　边界处理，如果两个链表都为空，返回nil，如果其中一个为空，返回另一个
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	// 如果list1.Val小，返回一个节点，Val为list1.Val，Next为mergeTwoLists1(list1.Next, list2)，list2于此类似
	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists1(list1.Next, list2),
		}
	} else {
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists1(list1, list2.Next),
		}
	}
}
```