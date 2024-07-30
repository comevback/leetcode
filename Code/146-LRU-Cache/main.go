package main

import "fmt"

func main() {
	LRUCache := Constructor(2)
	LRUCache.Put(1, 1)
	LRUCache.Put(2, 2)
	fmt.Println(LRUCache.Get(1))
	LRUCache.Put(3, 3)
	fmt.Println(LRUCache.Get(2))
	LRUCache.Put(4, 4)
	fmt.Println(LRUCache.Get(1))
	fmt.Println(LRUCache.Get(3))
	fmt.Println(LRUCache.Get(4))
}

// LRUCache 结构体定义
type LRUCache struct {
	hsMap map[int]*LinkHsMapNode // 哈希表用于快速检索节点
	Head  *LinkHsMapNode         // 双向链表的头节点
	Tail  *LinkHsMapNode         // 双向链表的尾节点
	cap   int                    // LRU 缓存的容量
}

// Constructor 初始化一个新的 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		hsMap: make(map[int]*LinkHsMapNode), // 初始化哈希表
		Head:  nil,                          // 初始时头节点为空
		Tail:  nil,                          // 初始时尾节点为空
		cap:   capacity,                     // 设置缓存容量
	}
}

// Get 方法通过键获取值，如果键存在于缓存中
func (this *LRUCache) Get(key int) int {
	if node, exist := this.hsMap[key]; exist {
		if node != this.Tail { // 如果节点不是尾节点，则需要调整位置
			if node != this.Head {
				node.Pre.Next = node.Next // 从链表中移除当前节点
			} else {
				this.Head = node.Next // 更新头节点
			}
			node.Next.Pre = node.Pre // 断开连接
			node.Next = nil          // 移动到尾部
			node.Pre = this.Tail
			this.Tail.Next = node
			this.Tail = node
		}
		return node.Val // 返回节点值
	} else {
		return -1 // 未找到键
	}
}

// Put 方法插入或更新键值对
func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.hsMap[key]; exist {
		node.Val = value       // 更新节点值
		if node != this.Tail { // 如果节点不是尾节点，调整位置
			if node != this.Head {
				node.Pre.Next = node.Next // 从链表中移除节点
			} else {
				this.Head = node.Next // 更新头节点
			}
			node.Next.Pre = node.Pre
			node.Next = nil
			node.Pre = this.Tail
			this.Tail.Next = node
			this.Tail = node
		}
		return
	}

	newNode := &LinkHsMapNode{ // 创建新节点
		Key: key,
		Val: value,
	}
	this.hsMap[key] = newNode // 加入哈希表

	if len(this.hsMap) <= this.cap { // 如果未达到容量上限
		if this.Head == nil { // 如果还没有节点
			this.Head = newNode
			this.Tail = newNode
			return
		}
		this.Tail.Next = newNode
		newNode.Pre = this.Tail
		this.Tail = newNode
		return
	} else { // 容量已满，需要移除最不常用项
		delete(this.hsMap, this.Head.Key) // 从哈希表中删除头节点
		if this.Head.Next != nil {
			this.Head = this.Head.Next // 更新头节点
			this.Head.Pre = nil
		}
		this.Tail.Next = newNode // 将新节点加入尾部
		newNode.Pre = this.Tail
		this.Tail = newNode
	}
}

type LinkHsMapNode struct {
	Key  int
	Val  int
	Next *LinkHsMapNode
	Pre  *LinkHsMapNode
}

// ------------------------------------------------------------------------------------------------------------
// 移动节点到双向链表的尾部
func (this *LRUCache) moveToTail(node *LinkHsMapNode) {
	if node == this.Tail {
		return // 如果节点已是尾部，无需操作
	}
	// 断开节点的前后连接
	if node != this.Head {
		node.Pre.Next = node.Next
	} else {
		this.Head = node.Next
	}
	node.Next.Pre = node.Pre

	// 将节点加到尾部
	node.Next = nil
	node.Pre = this.Tail
	this.Tail.Next = node
	this.Tail = node
}

// 添加新节点到双向链表的尾部
func (this *LRUCache) addToTail(newNode *LinkHsMapNode) {
	if this.Head == nil {
		this.Head = newNode
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		newNode.Pre = this.Tail
		this.Tail = newNode
	}
}

// 移除双向链表的头节点（最久未使用）
func (this *LRUCache) removeHead() {
	delete(this.hsMap, this.Head.Key)
	this.Head = this.Head.Next
	if this.Head != nil {
		this.Head.Pre = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
