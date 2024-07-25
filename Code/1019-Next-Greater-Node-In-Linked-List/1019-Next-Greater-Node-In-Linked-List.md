# 1019. Next Greater Node In Linked List

Solved
Medium
Topics
Companies
Hint
You are given the head of a linked list with n nodes.

For each node in the list, find the value of the next greater node. That is, for each node, find the value of the first node that is next to it and has a strictly larger value than it.

Return an integer array answer where answer[i] is the value of the next greater node of the ith node (1-indexed). If the ith node does not have a next greater node, set answer[i] = 0.

Example 1:

Input: head = [2,1,5]
Output: [5,5,0]
Example 2:

Input: head = [2,7,4,3,5]
Output: [7,0,5,5,0]

Constraints:

The number of nodes in the list is n.
1 <= n <= 104
1 <= Node.val <= 109

---

# Code

```go
package main

import "fmt"

func main() {
	head := &ListNode{Val: 3}     // 创建链表的头节点，值为 3
	head.Next = &ListNode{Val: 3} // 创建链表的第二个节点，值也为 3

	res := nextLargerNodes(head) // 调用函数来寻找链表中每个节点的下一个更大节点的值
	fmt.Println(res)             // 打印结果
}

// ListNode 定义了链表的节点结构。
type ListNode struct {
	Val  int       // 节点存储的整数值
	Next *ListNode // 指向下一个节点的指针
}

// nextLargerNodes 接受链表的头节点，并返回一个数组，其中每个元素表示链表中对应节点的下一个更大元素的值。
func nextLargerNodes(head *ListNode) []int {
	arr := []int{}  // 用于存储链表的值
	current := head // 使用 current 遍历链表
	// 遍历链表，将所有节点的值存入 arr
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}

	res := make([]int, len(arr)) // 创建结果数组，初始值为 0
	temp := []int{}              // 用于存储还未找到下一个更大元素的节点值

	// 从后向前遍历 arr，寻找每个元素的下一个更大元素
	for i := len(arr) - 1; i >= 0; i-- {
		// 移除 temp 中所有小于或等于当前元素 arr[i] 的元素
		for len(temp) > 0 && arr[i] >= temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		}

		// 如果 temp 为空，说明没有更大的元素，否则 temp 的最后一个元素就是 arr[i] 的下一个更大元素
		if len(temp) == 0 {
			res[i] = 0
		} else {
			res[i] = temp[len(temp)-1]
		}

		// 将当前元素 arr[i] 添加到 temp 中，以便后续比较
		temp = append(temp, arr[i])
	}

	return res // 返回结果数组
}
```
