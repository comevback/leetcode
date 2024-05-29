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
