# 380. Insert Delete GetRandom O(1)
Solved
Medium
Topics
Companies
Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity. 

Example 1:

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
 

Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.

---

# Code
```go
package main

import "math/rand"

// main函数演示了RandomizedSet类的使用
func main() {
	obj := Constructor() // 创建RandomizedSet对象
	obj.Insert(0)        // 插入元素0
	obj.Insert(1)        // 插入元素1
	obj.Remove(0)        // 移除元素0
	obj.Insert(2)        // 插入元素2
	obj.Remove(1)        // 移除元素1
	obj.GetRandom()      // 获取一个随机元素
}

// RandomizedSet结构体，包含一个哈希表和一个动态数组
type RandomizedSet struct {
	hsMap map[int]int // 哈希表存储元素值和其在数组中的索引
	arr   []int       // 数组存储元素值
}

// Constructor初始化一个新的RandomizedSet对象
func Constructor() RandomizedSet {
	return RandomizedSet{
		hsMap: make(map[int]int),
		arr:   []int{},
	}
}

// Insert方法将元素val插入到集合中。如果元素已存在，返回false，否则插入并返回true。
func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.hsMap[val]; !exist {
		this.hsMap[val] = len(this.arr)  // 记录元素的索引
		this.arr = append(this.arr, val) // 将元素添加到数组
		return true
	} else {
		return false
	}
}

// Remove方法从集合中移除元素val。如果元素存在，则移除并返回true，否则返回false。
func (this *RandomizedSet) Remove(val int) bool {
	if index, exist := this.hsMap[val]; exist {
		last := this.arr[len(this.arr)-1]     // 获取数组最后一个元素
		this.arr[index] = last                // 将最后一个元素移到要删除的元素位置
		this.hsMap[last] = index              // 更新哈希表中最后一个元素的索引
		this.arr = this.arr[:len(this.arr)-1] // 删除数组最后一个元素
		delete(this.hsMap, val)               // 从哈希表删除元素
		return true
	} else {
		return false
	}
}

// GetRandom方法返回集合中的一个随机元素。
func (this *RandomizedSet) GetRandom() int {
	if len(this.arr) == 0 {
		return -1 // 如果数组为空，返回-1
	}
	randomNum := rand.Intn(len(this.arr)) // 生成一个随机索引
	return this.arr[randomNum]            // 返回该索引对应的元素
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
```