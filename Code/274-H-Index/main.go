package main

import "fmt"

func main() {
	// 示例引用次数数组
	cit := []int{3, 0, 6, 1, 5}
	// 计算 h 指数
	res := hIndex(cit)
	// 打印 h 指数
	fmt.Println(res)
}

// hIndex 计算给定引用次数数组的 h 指数
func hIndex(citations []int) int {
	res := 0
	length := len(citations)
	// 临时数组，用于合并排序
	temp := make([]int, len(citations))
	// 对引用次数数组进行合并排序
	mergeSort(citations, temp, 0, length-1)
	// 计算 h 指数
	for i, v := range citations {
		numsLeft := length - i
		if v >= numsLeft {
			res = numsLeft
			return res
		}
	}
	return 0
}

// mergeSort 对数组进行合并排序
func mergeSort(nums []int, temp []int, left int, right int) {
	if left >= right {
		return
	}

	// 计算中间索引
	mid := left + (right-left)/2
	// 递归排序左半部分
	mergeSort(nums, temp, left, mid)
	// 递归排序右半部分
	mergeSort(nums, temp, mid+1, right)
	// 合并排序结果
	sortIndex(nums, temp, left, mid, right)
}

// sortIndex 合并两个已排序的子数组
func sortIndex(nums []int, temp []int, left int, mid int, right int) {
	// 复制待排序部分到临时数组
	copy(temp[left:right+1], nums[left:right+1])
	var i, j int = left, mid + 1
	var current int = left

	// 合并两个子数组
	for i < mid+1 && j < right+1 {
		if temp[i] <= temp[j] {
			nums[current] = temp[i]
			i += 1
		} else {
			nums[current] = temp[j]
			j += 1
		}
		current += 1
	}

	// 复制剩余的左半部分元素
	for i < mid+1 {
		nums[current] = temp[i]
		i += 1
		current += 1
	}

	// 复制剩余的右半部分元素
	for j < right+1 {
		nums[current] = temp[j]
		j += 1
		current += 1
	}
}
