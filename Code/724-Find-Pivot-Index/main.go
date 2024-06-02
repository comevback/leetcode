package main

import "fmt"

func main() {
	nums := []int{2, 1, -1}
	res := pivotIndex2(nums)
	fmt.Println(res)
}

// 1. map
// 一次遍历，map储存所有前后和，二次遍历，对比所有元素的前后和，有相等的就返回该index，无则返回-1
func pivotIndex(nums []int) int {
	// 定义前后和
	var preSum int
	var backSum int
	// 定义每个元素的前后和map
	var preMap map[int]int = make(map[int]int, len(nums))
	var backMap map[int]int = make(map[int]int, len(nums))
	// 双指针i,j分别指向数组头尾
	i, j := 0, len(nums)-1

	// 遍历数组，计算每个元素的前后和
	for i < len(nums)-1 {
		// 前和等于之前的前和+当前值
		preSum += nums[i]
		// 后和等于之前的后和+当前值
		backSum += nums[j]
		// map储存前后和，因为头尾元素前后和都是0，不包括自身，所以map的记录是从头尾第二个元素开始
		preMap[i+1] = preSum
		backMap[j-1] = backSum

		// 双指针移动
		i += 1
		j -= 1
	}

	// 遍历数组，对比每个元素在map中的前后和
	// 有相等的就返回该index
	for i := 0; i < len(nums); i++ {
		if preMap[i] == backMap[i] {
			return i
		}
	}

	// 无则返回-1
	return -1
}

// 2.一次遍历求和
// 二次遍历直接对比前后和
func pivotIndex1(nums []int) int {

	// 求整个数组的和
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 定义前和和后和
	pre := 0
	back := sum

	// 遍历数组，对比每个元素的前后和
	for i, v := range nums {
		// 后和等于之前的后和-当前值
		back -= v
		if pre == back {
			return i
		}
		// 下一个元素的前和才加当前元素值
		pre += v
	}

	// 无则返回-1
	return -1
}

// ========================================================================================================
// review
func pivotIndex2(nums []int) int {
	sum := 0
	prefix := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		prefix[i] = sum
		sum += nums[i]
	}

	index := -1
	for i := 0; i < len(nums); i++ {
		if sum-prefix[i]-nums[i] == prefix[i] {
			index = i
			break
		}
	}

	return index
}
