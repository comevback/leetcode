package main

func main() {

}

// ======================================== 1. 定义链表结构 ========================================

type LinkedList struct {
	head *LinkedNode
	tail *LinkedNode
}

// ======================================== 2. 定义链表节点结构 ========================================
type LinkedNode struct {
	value any
	Next  *LinkedNode
}

// ======================================== 3. append方法 ========================================
// append方法 - LinkedNode
func (node *LinkedNode) Append(newNode *LinkedNode) {
	if node.Next != nil {
		node.Next.Append(newNode)
	} else {
		node.Next = newNode
	}
}

// append方法 - LinkedList
func (lkList *LinkedList) Append(newNode *LinkedNode) {
	if lkList.tail != nil {
		lkList.tail.Next = newNode
		lkList.tail = newNode
	} else {
		lkList.head = newNode
		lkList.tail = newNode
	}
}

// ======================================== 4. prepend方法 ========================================
// prepend方法 - LinkedNode
func (node *LinkedNode) Prepend(newNode *LinkedNode) *LinkedNode {
	newNode.Next = node
	return newNode
}

// prepend方法 - LinkedList
func (lkList *LinkedList) Prepend(newNode *LinkedNode) {
	newNode.Next = lkList.head
	lkList.head = newNode
}

// ======================================== 5. insert方法 ========================================
// Insert方法 - LinkedNode
func (node *LinkedNode) Insert(newNode *LinkedNode, index int) {
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
func (lkList *LinkedList) Insert(newNode *LinkedNode, index int) {
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
// delete方法 - LinkedNode
func (node *LinkedNode) remove(deleteNode *LinkedNode) {
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
func (lkList *LinkedList) remove(deleteNode *LinkedNode) {

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
