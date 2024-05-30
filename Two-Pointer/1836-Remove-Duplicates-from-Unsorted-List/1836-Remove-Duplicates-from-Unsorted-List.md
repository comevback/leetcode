# 1836. 从未排序的链表中移除重复元素
给定一个链表的第一个节点 head ，找到链表中所有出现多于一次的元素，并删除这些元素所在的节点。

返回删除后的链表。

示例 1:


输入: head = [1,2,3,2]
输出: [1,3]
解释: 2 在链表中出现了两次，所以所有的 2 都需要被删除。删除了所有的 2 之后，我们还剩下 [1,3] 。
示例 2:


输入: head = [2,1,1,2]
输出: []
解释: 2 和 1 都出现了两次。所有元素都需要被删除。
示例 3:


输入: head = [3,2,2,1,3,2,4]
输出: [1,4]
解释: 3 出现了两次，且 2 出现了三次。移除了所有的 3 和 2 后，我们还剩下 [1,4] 。
提示：

链表中节点个数的范围是 [1, 105]
1 <= Node.val <= 105

---

# Code
```go
package main

import "fmt"

func main() {
	// 创建链表 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
	head := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}

	fmt.Println("Original list:")
	printList(head)

	// 删除重复节点
	head = removeDuplicateNodes(head)

	fmt.Println("List after removing duplicates:")
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两次遍历，第一次记录每个元素出现次数，第二次找到出现次数等于1的连接
func removeDuplicateNodes(head *ListNode) *ListNode {
	var Repeated map[int]int = make(map[int]int)
	current := head

	for current != nil {
		Repeated[current.Val] += 1
		current = current.Next
	}

	var dummy *ListNode = &ListNode{}
	dummy.Next = head
	deleted, current := dummy, head

	for current != nil {
		if Repeated[current.Val] == 1 {
			deleted.Next = current
			deleted = deleted.Next
		}
		current = current.Next
	}

	deleted.Next = nil
	return dummy.Next
}
```