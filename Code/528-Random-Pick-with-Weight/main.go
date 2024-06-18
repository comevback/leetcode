package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 3, 4, 5, 6}
	obj := Constructor(nums)
	res := obj.PickIndex()
	fmt.Println(res)
}

// Solution 结构体包含一个整数切片 Arr，用来存储权重的累积和。
type Solution struct {
	Arr []int
}

// Constructor 是 Solution 结构体的构造函数，它接受一个权重数组 w，并返回一个 Solution 实例。
// 这个函数计算权重的累积和并存储在 Arr 中。
func Constructor(w []int) Solution {
	var weights []int = make([]int, len(w))
	sum := 0
	for i, v := range w {
		sum += v
		weights[i] = sum
	}

	return Solution{
		Arr: weights,
	}
}

// PickIndex 是 Solution 类的方法，用于根据权重数组随机选择一个索引。
// 它首先生成一个 [0, Arr 最后一个元素) 范围内的随机数。
// 然后使用二分查找确定随机数落在哪个权重区间。
func (this *Solution) PickIndex() int {
	r := rand.Intn(this.Arr[len(this.Arr)-1])

	left, right := 0, len(this.Arr)
	for left < right {
		mid := left + (right-left)/2
		if this.Arr[mid] == r {
			return mid + 1
		} else if this.Arr[mid] < r {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := PickIndex();
 */
