package linkedlist

import "fmt"

func main() {

}

// ======================================== 1. 定义链表和节点 ========================================

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	Val  any
	Next *ListNode
}

// ======================================== 2. 构造方法  ========================================

// 链表的构造函数
func NewLinkedList(nums []int) *LinkedList {

	// 如果数组为空，直接返回空链表
	if len(nums) == 0 {
		return nil
	}

	// 定义链表头和当前节点，头节点值为数组第一个元素
	var head *ListNode = &ListNode{Val: nums[0]}
	var current *ListNode = &ListNode{}
	// 初始时头节点的下一个节点为当前节点
	head.Next = current

	// 遍历数组，构造链表，把每个节点的值赋值为数组当前元素的值
	for i := 1; i < len(nums); i++ {
		// 如果是最后一个元素，就直接赋值给当前节点的值，不要新增节点
		if i == len(nums)-1 {
			current.Val = nums[i]
			break
		}
		current.Val = nums[i]
		newNode := &ListNode{}
		current.Next = newNode
		current = current.Next
	}

	// 链表中，头节点就是链表的头，当前节点就是链表的尾
	var linkList *LinkedList = &LinkedList{
		head: head,
		tail: current,
	}

	// 返回链表
	return linkList
}

// ======================================== 3. append方法 ========================================
// append方法 - ListNode
func (node *ListNode) Append(newNode *ListNode) {
	if node.Next != nil {
		node.Next.Append(newNode)
	} else {
		node.Next = newNode
	}
}

// append方法 - LinkedList
func (lkList *LinkedList) Append(newNode *ListNode) {
	if lkList.tail != nil {
		lkList.tail.Next = newNode
		lkList.tail = newNode
	} else {
		lkList.head = newNode
		lkList.tail = newNode
	}
}

// ======================================== 4. prepend方法 ========================================
// prepend方法 - ListNode
func (node *ListNode) Prepend(newNode *ListNode) *ListNode {
	newNode.Next = node
	return newNode
}

// prepend方法 - LinkedList
func (lkList *LinkedList) Prepend(newNode *ListNode) {
	newNode.Next = lkList.head
	lkList.head = newNode
}

// ======================================== 5. insert方法 ========================================
// Insert方法 - ListNode
func (node *ListNode) Insert(newNode *ListNode, index int) {
	if index == 0 {
		node.Prepend(newNode)
	}

	if index >= 1 && node.Next == nil {
		panic("not of index")
	}

	if index == 1 {
		newNode.Next = node.Next
		node.Next = newNode
	} else {
		node.Next.Insert(newNode, index-1)
	}
}

// Insert方法 - LinkedList
func (lkList *LinkedList) Insert(newNode *ListNode, index int) {
	if index == 0 {
		lkList.Prepend(newNode)
	} else {
		node := lkList.head
		for index > 1 {
			if node.Next == nil {
				panic("index out of range")
			}
			index -= 1
			node = node.Next
		}
		if node.Next != nil {
			newNode.Next = node.Next
			node.Next = newNode
		} else {
			newNode.Next = nil
			node.Next = newNode
			lkList.tail = newNode
		}
	}
}

// ======================================== 6. delete方法 ========================================
// delete方法 - ListNode
func (node *ListNode) remove(deleteNode *ListNode) {
	// 检查当前节点的下一个节点是否是要删除的节点
	if node.Next == deleteNode {
		// 使当前节点直接指向要删除节点的下一个节点，从而删除之
		node.Next = deleteNode.Next
		deleteNode.Next = nil // 帮助垃圾收集
	} else if node.Next != nil {
		// 如果不是，递归到下一个节点
		node.Next.remove(deleteNode)
	} else {
		// 如果已经到链表尾部还没有找到，可能会抛出异常或进行错误处理
		panic("Node not found")
	}
}

// delete方法 - LinkedList
func (lkList *LinkedList) remove(deleteNode *ListNode) {

	//
	if lkList.head == deleteNode {
		lkList.head = lkList.head.Next
		if lkList.head == nil {
			lkList.tail = nil
		}
	}

	node := lkList.head
	for node.Next != nil {
		if node.Next == deleteNode {
			node.Next = deleteNode.Next
			deleteNode.Next = nil
			return
		}
		node = node.Next
	}
	panic("node not found")
}

// ======================================== 7. 打印链表方法  ========================================

// 打印链表为数组
func printList(head *ListNode) {
	var arr []any
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	fmt.Println(arr)
}

// 翻转链表  具体实现 （206-reverse-linked-list）
func reverseLink(head *ListNode) *ListNode {
	// 如果链表为空，直接返回
	if head == nil {
		return head
	}
	// 定义新链表头
	var newHead *ListNode

	// 遍历链表，把每个节点的Next指向新链表头，然后更新新链表头为当前节点
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}

	// 返回新链表头
	return newHead
}
