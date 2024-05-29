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
