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
