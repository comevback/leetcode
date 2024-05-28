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

// ========================================================  review  ======================================================
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{}
	var current = dummy

	current1, current2 := list1, list2

	for current1 != nil && current2 != nil {
		if current1.Val <= current2.Val {
			current.Next = current1
			current1 = current1.Next
		} else {
			current.Next = current2
			current2 = current2.Next
		}
		current = current.Next
	}

	if current1 != nil {
		current.Next = current1
	}

	if current2 != nil {
		current.Next = current2
	}

	return dummy.Next
}
