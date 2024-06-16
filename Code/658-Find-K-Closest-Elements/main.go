package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 2, 2, 2, 3, 3}
	res := findClosestElements1(arr, 3, 3)
	fmt.Println(res)
}

// 1. 我的方法，二分查找后双指针
func findClosestElements1(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	res := []int{}
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == x {
			right = mid
		} else if arr[mid] > x {
			right = mid
		} else if arr[mid] < x {
			left = mid + 1
		}
	}

	if left == len(arr) {
		left = len(arr) - 1
	}

	leftIdx := left - 1
	rightIdx := left

	for len(res) < k {
		if leftIdx >= 0 && rightIdx < len(arr) {
			if abs(arr[leftIdx]-x) <= abs(arr[rightIdx]-x) {
				res = append([]int{arr[leftIdx]}, res...)
				leftIdx -= 1
			} else {
				res = append(res, arr[rightIdx])
				rightIdx += 1
			}
		} else if leftIdx >= 0 {
			res = append([]int{arr[leftIdx]}, res...)
			leftIdx -= 1
		} else if rightIdx < len(arr) {
			res = append(res, arr[rightIdx])
			rightIdx += 1
		}
	}

	return res
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

// ======================================================  二分查找  ================================================
// 2. 二分查找
func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	// 初始化二分查找边界
	// left 指针表示可能的起始位置的最低边界
	// right 指针表示可能的起始位置的最高边界，它是 len(arr)-k
	// 这样设置 right 边界确保从这个点开始的 k 个元素不会超出数组的范围
	left, right := 0, len(arr)-k

	// 执行二分查找
	for left < right {
		// 计算中间位置
		mid := left + (right-left)/2

		// 比较 x 与当前中点及其 k 个元素后的那个元素的距离
		// 以确定搜索应向左还是向右收缩
		if x-arr[mid] > arr[mid+k]-x {
			// 如果 arr[mid] 与 x 的距离大于 arr[mid+k] 与 x 的距离
			// 表明最接近 x 的 k 个元素应该在当前 mid 点的右边
			left = mid + 1
		} else {
			// 否则，表示 arr[mid] 是靠近 x 的点，或者更接近的点可能在左侧
			right = mid
		}
	}

	// left 最终将停在最接近 x 的 k 个元素的起始索引上
	// 返回从 left 开始的 k 个元素
	return arr[left : left+k]
}
