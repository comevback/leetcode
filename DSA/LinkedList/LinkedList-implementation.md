# 手写linkedList

---

# 完整实现

```go
package main

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
```


---

# 翻转链表方法

[206-Reverse-Linked-List.md](206-Reverse-Linked-List/206-Reverse-Linked-List.md)

## 思路

> 1. 定义一个新的链表，头节点为new，初始化为nil
> 2. 把当前链表的head节点的下一个节点head.Next保存给临时变量temp，方便之后head转移为head.Next。
> 3. head.Next指向new，即head.Next = new，这样head节点就从原链表中脱离出来，接入到新链表中。
![206-Reverse-Linked-List-1.jpg](/DSA-implementation/LinkedList/206-Reverse-Linked-List/206-Reverse-Linked-List-1.jpg)
> 4. 把这里的head变成new，即new = head，这样new就还是新链表的头节点。
> 5. 把head变成temp，即head = temp，等于开启了下一轮的循环。原本的head节点已经接入到新链表中。
![206-Reverse-Linked-List-2](/DSA-implementation/LinkedList/206-Reverse-Linked-List/206-Reverse-Linked-List-2.jpg) 
> 6. 重复2-5步骤，直到head为nil，即原链表的尾节点。
![206-Reverse-Linked-List-3](/DSA-implementation/LinkedList/206-Reverse-Linked-List/206-Reverse-Linked-List-3.jpg)
![reverse-linkedlist.png](/DSA-implementation/LinkedList/206-Reverse-Linked-List/reverse-linkedlist.png)

---

## 代码实现

```go
package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printLinnkedList(head)
	res := reverseList(head)
	printLinnkedList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.对调linkNode的值Val
// 没有真正改变链表的链接结构，只重新分配了值
func reverseList1(head *ListNode) *ListNode {

	// 如果链表为nil或只有一个头元素，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	// 两个指针，一个用于遍历链表得到每个元素的值，一个用于修改链表里每个元素的值
	var temp1 *ListNode = head
	var temp2 *ListNode = head

	// 用于存储链表的值
	var arr []int

	// 遍历链表，将链表的值存储到arr数组中
	for temp1 != nil {
		arr = append(arr, temp1.Val)
		temp1 = temp1.Next
	}

	// 第二次从头遍历链表，将arr数组的值倒序赋值给链表
	count := len(arr) - 1
	for temp2 != nil {
		temp2.Val = arr[count]
		count -= 1
		temp2 = temp2.Next
	}

	// 返回修改值后的原链表
	return head
}

// =======================================================================================================================
// 2.正确的链表反转
func reverseList(head *ListNode) *ListNode {

	// 如果链表为nil或只有一个头元素，直接返回head
	if head == nil || head.Next == nil {
		return head
	}

	// 设一个新的数组，头元素newHead暂时为nil
	var newHead *ListNode = nil

	// 当head.Next不为nil时，继续循环
	for head != nil {
		// 保存head.Next的值
		next := head.Next
		// head.Next指向newHead
		head.Next = newHead
		// newHead前进为head
		newHead = head
		// head前进为next
		head = next
	}

	// 返回newHead
	return newHead
}
```

---

# 注意点

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