package main

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}
	reorderList1(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ************************************************* 2.第一种思路：建立一个反向链表，交叉组合 ************************************************
// 1.第一种想法
func reorderList(head *ListNode) {
	// 使用快慢指针找到链表的中点
	left, right := head, head
	for right != nil {
		right = right.Next
		if right != nil {
			right = right.Next
		} else {
			break
		}
		left = left.Next
	}

	// 反转链表并获取长度
	reversed, length := getReverse(head)
	curLen := 1
	current := head

	// 按照重排顺序重新链接链表节点
	for curLen < length {
		temp := current.Next    // 保存原链表的下一个节点
		tempRe := reversed.Next // 保存反转链表的下一个节点
		current.Next = reversed // 将当前节点指向反转链表的当前节点
		current = current.Next  // 移动当前指针到下一个位置
		curLen += 1

		if curLen == length { // 如果已经处理到最后一个节点，停止
			current.Next = nil
			break
		}

		current.Next = temp    // 将当前节点指向原链表的下一个节点
		current = current.Next // 移动当前指针到下一个位置
		curLen += 1

		if curLen == length { // 如果已经处理到最后一个节点，停止
			current.Next = nil
			break
		}

		reversed = tempRe // 移动反转链表的指针到下一个位置
	}
}

// getReverse 函数反转链表并返回新的头节点及链表长度
func getReverse(head *ListNode) (*ListNode, int) {
	var newHead *ListNode
	arr := []int{}

	// 将链表节点值存入数组中
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	length := len(arr)

	// 从数组中依次取出节点值，创建新的反转链表
	for len(arr) > 0 {
		node := &ListNode{Val: arr[0]}
		arr = arr[1:]
		node.Next = newHead
		newHead = node
	}

	return newHead, length
}

// ************************************************* 2.第二种思路：中间切开链表，翻转后半部分 ************************************************
// 2.第二种思路
// 应该找到链表的中间位置，将链表分为前后两部分，然后反转后半部分，并最终交替连接前半部分和反转后的后半部分。
// 实现方法：
// 找到链表的中间节点。
// 将链表分为前后两部分。
// 反转后半部分链表。
// 交替连接前半部分和反转后的后半部分。

func reverse(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

// reorderList1 函数按照要求重排链表
func reorderList1(head *ListNode) {
	// 使用快慢指针找到链表的中点
	left, right := head, head
	for right != nil {
		right = right.Next
		if right != nil {
			right = right.Next
		} else {
			break
		}
		left = left.Next
	}

	// 将链表分成两部分并反转后半部分
	newList := left.Next
	left.Next = nil
	newList = reverse(newList)

	// 交替合并两部分链表
	for newList != nil {
		// 保存原链表和反转链表的下一个节点
		temp := head.Next
		tempRe := newList.Next

		// 将新节点插入到原链表中
		head.Next = newList
		head = head.Next

		// 将原链表节点插入到新节点之后
		head.Next = temp
		head = head.Next

		// 移动到下一个新节点
		newList = tempRe
	}
}

// ******************************************************* 3.第三种思路：栈 *************************************************************
func reorderList2(head *ListNode) {
	// 先把所有节点装进栈里，得到倒序结果
	type stack []*ListNode
	var stk stack
	p := head
	for p != nil {
		stk = append(stk, p)
		p = p.Next
	}

	p = head
	for p != nil {
		// 链表尾部的节点
		lastNode := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		next := p.Next
		if lastNode == next || lastNode.Next == next {
			// 结束条件，链表节点数为奇数或偶数时均适用
			lastNode.Next = nil
			break
		}
		p.Next = lastNode
		lastNode.Next = next
		p = next
	}
}
