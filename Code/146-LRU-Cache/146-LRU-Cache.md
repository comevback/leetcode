# 146. LRU Cache（双向哈希链表）

Solved
Medium
Topics
Companies
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.

Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1); // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2); // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1); // return -1 (not found)
lRUCache.get(3); // return 3
lRUCache.get(4); // return 4

Constraints:

1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 \* 105 calls will be made to get and put.

---

# Code

```go
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
```
