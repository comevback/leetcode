package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 1}
	res := containsNearbyAlmostDuplicate(nums, 3, 0)
	fmt.Println(res)
}

// 1. 时间复杂度太高
func containsNearbyAlmostDuplicate1(nums []int, indexDiff int, valueDiff int) bool {
	left, right := 0, 0
	countMap := make(map[int]int)

	for right < len(nums) {
		if countMap[nums[right]] > 0 {
			return true
		} else {
			for i := nums[right] - valueDiff; i <= nums[right]+valueDiff; i++ {
				countMap[i] += 1
			}
		}

		right += 1
		if right-left > indexDiff {
			for i := nums[left] - valueDiff; i <= nums[left]+valueDiff; i++ {
				countMap[i] -= 1
			}
			left += 1
		}
	}
	return false
}

// =======================================================================================================================

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) == 0 || indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	right := 0
	window := make([]int, 0, indexDiff+1)

	for right < len(nums) {
		// 在 window 中找到最接近 nums[right] 的元素
		pos := sort.SearchInts(window, nums[right])
		if pos < len(window) && abs(window[pos]-nums[right]) <= valueDiff {
			return true
		}
		if pos > 0 && abs(window[pos-1]-nums[right]) <= valueDiff {
			return true
		}

		// 将新元素添加到 window
		window = append(window, nums[right])

		// 保持 window 的长度不超过 indexDiff
		if len(window) > indexDiff {
			// 移除最早的元素
			window = window[1:]
		}
	}
	return false
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

// 寻找左右相邻的index
func searchIndex(arr []int, num int, leftidx int, rightidx int) (int, int) {
	if leftidx >= rightidx {
		// 边界处理：当搜索范围缩小到单个元素时
		if num > arr[leftidx] {
			return leftidx, leftidx + 1
		} else {
			return leftidx - 1, leftidx
		}
	}

	mid := (leftidx + rightidx) / 2

	if num > arr[mid] {
		return searchIndex(arr, num, mid+1, rightidx)
	} else if num < arr[mid] {
		return searchIndex(arr, num, leftidx, mid)
	} else {
		return mid, mid
	}
}

// 插入排序
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	} else {
		mid := len(nums) / 2
		left := MergeSort(nums[:mid])
		right := MergeSort(nums[mid:])
		return merge(left, right)
	}
}

func merge(nums1 []int, nums2 []int) []int {
	p1, p2 := 0, 0
	res := make([]int, 0, len(nums1)+len(nums2))

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] <= nums2[p2] {
			res = append(res, nums1[p1])
			p1 += 1
		} else {
			res = append(res, nums2[p2])
			p2 += 1
		}
	}

	if p1 < len(nums1) {
		res = append(res, nums1[p1:]...)
	}

	if p2 < len(nums2) {
		res = append(res, nums2[p2:]...)
	}

	return res
}
