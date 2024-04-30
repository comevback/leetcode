# 手写linkedList

---

# 完整实现

```go
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

```

---

# 踩坑记录

### 结构和append方法
```go
package main

func main() {

}

type LinkedNode struct {
	value any
	Next  *LinkedNode
}

func (node *LinkedNode) append(newNode LinkedNode) {
	if node.Next != nil {
		node.Next.append(newNode)
	} else {
		node.Next = &newNode
	}
}
```

### 错误点：
在 append 方法中，你正确地递归查找列表末尾来添加新节点，但是直接传递 LinkedNode 而不是指针可能会导致问题。应当传递指向 LinkedNode 的指针以避免复制整个节点。
下面是一个改进后的版本：

```go
Copy code
type LinkedNode struct {
    value any
    Next  *LinkedNode
}

// append 方法接收一个指向 LinkedNode 的指针，而不是 LinkedNode本身
func (node *LinkedNode) append(newNode *LinkedNode) {
    if node.Next != nil {
        node.Next.append(newNode)
    } else {
        node.Next = newNode // 这里不需要取地址符
    }
}
```
在这个改进版本中，append 方法接收一个指向 LinkedNode 的指针，这样就可以直接将新节点链接到链表的末尾，而不是复制节点。这样做更有效，也避免了不必要的内存使用。


### prepend 方法
```go
// prepend方法
func (node *LinkedNode) Prepend(newNode *LinkedNode) {
	newNode.Next = node
}
```

### 错误点
在 prepend 方法中，你试图将新节点插入到链表的开头，但是你忘记了更新链表的头指针。下面是一个改进后的版本：

```go
// prepend方法
func (node *LinkedNode) Prepend(newNode *LinkedNode) *LinkedNode{
	newNode.Next = node
    return newNode
}
```

### Insert方法
```go
// Insert方法
func (node *LinkedNode) Insert(newNode *LinkedNode, index int) {
	// 如果index为0，将新节点插入到链表的开头
    if index == 0 {
		node.Prepend(newNode)
	}

    // 如果index大于0，将新节点插入到链表的中间
	if index >= 1 && node.Next == nil {
		panic("not of index")
	}

    // 递归调用Insert方法
    // 如果index大于0，将新节点插入到链表的中间
	if index == 1 {
		newNode.Next = node.Next
		node.Next = newNode
    // 如果index大于1，递归调用Insert方法
	} else {
		node.Next.Insert(newNode, index-1)
	}
}
```

### delete（remove）方法
```go
// delete方法
func (node *LinkedNode) remove(deleteNode *LinkedNode) {
	if node == deleteNode {
		deleteNode = deleteNode.Next
	} else {
		node.Next.remove(deleteNode)
	}
}
```

### 错误点
在 remove 方法中，你试图删除链表中的一个节点，但是你忘记了更新链表的头指针。下面是一个改进后的版本：

```go
// delete方法
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
```