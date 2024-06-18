package main

import "fmt"

func main() {
	nums1 := []int{12, 24, 8, 32}
	nums2 := []int{13, 25, 32, 11}
	res := advantageCount(nums1, nums2)
	fmt.Println(res)
}

// 二分查找
// advantageCount 接收两个整数数组 nums1 和 nums2，返回 nums1 的一个排列，使其对 nums2 的优势最大化。
func advantageCount(nums1 []int, nums2 []int) []int {
	temp := make([]int, len(nums1))       // 创建一个与 nums1 等长的临时数组
	res := make([]int, len(nums1))        // 创建结果数组
	MergeSort(nums1, temp, 0, len(nums1)) // 对 nums1 进行归并排序

	for i, v := range nums2 {
		index := binarySearch(nums1, v) // 对于 nums2 中的每个元素，找到 nums1 中大于它的最小元素的索引
		if index != -1 {
			res[i] = nums1[index] // 如果找到，放入结果数组
			temp := nums1[:index] // 重新构建 nums1，去掉已经使用的元素
			temp = append(temp, nums1[index+1:]...)
			nums1 = temp
		} else {
			res[i] = nums1[0] // 如果没有找到大于的元素，使用 nums1 中最小的元素
			temp := nums1[1:]
			nums1 = temp
		}
	}

	return res
}

// 二分查找 + 使用map存储nums2的索引
func advantageCount1(nums1 []int, nums2 []int) []int {
	temp := make([]int, len(nums1))
	res := make([]int, len(nums1))
	MergeSort(nums1, temp, 0, len(nums1))

	// 如果需要保留原始索引，可以将nums2的索引存储到map中
	indexMap := make(map[int][]int)
	for i, v := range nums2 {
		indexMap[v] = append(indexMap[v], i)
	}

	MergeSort(nums2, temp, 0, len(nums2))

	for _, v := range nums2 {
		index := binarySearch(nums1, v)
		resIndex := indexMap[v][len(indexMap[v])-1]
		indexMap[v] = indexMap[v][:len(indexMap[v])-1]
		if index != -1 {
			res[resIndex] = nums1[index]
			temp := nums1[:index]
			temp = append(temp, nums1[index+1:]...)
			nums1 = temp
		} else {
			res[resIndex] = nums1[0]
			temp := nums1[1:]
			nums1 = temp
		}
	}

	return res
}

// 双指针
func advantageCount2(nums1 []int, nums2 []int) []int {
	temp := make([]int, len(nums1))
	res := make([]int, len(nums1))
	MergeSort(nums1, temp, 0, len(nums1))
	indexMap := make(map[int][]int)
	for i, v := range nums2 {
		indexMap[v] = append(indexMap[v], i)
	}
	MergeSort(nums2, temp, 0, len(nums2))

	left, right := 0, len(nums1)-1
	for i := len(nums2) - 1; i >= 0; i-- {
		resIndex := indexMap[nums2[i]][len(indexMap[nums2[i]])-1]
		indexMap[nums2[i]] = indexMap[nums2[i]][:len(indexMap[nums2[i]])-1]
		if nums1[right] > nums2[i] {
			res[resIndex] = nums1[right]
			right -= 1
		} else {
			res[resIndex] = nums1[left]
			left += 1
		}
	}

	return res
}

// ================================================  工具函数  ===================================================
// 二分查找右边界
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left == len(nums) {
		return -1
	}

	if nums[left] == target {
		left += 1
		if left == len(nums) {
			return -1
		}
	}

	return left
}

// 归并排序
func MergeSort(nums []int, temp []int, left int, right int) {
	if left == right-1 {
		return
	}

	mid := left + (right-left)/2
	MergeSort(nums, temp, left, mid)
	MergeSort(nums, temp, mid, right)

	merge(nums, temp, left, mid, right)
}

func merge(nums []int, temp []int, left int, mid int, right int) {
	copy(temp[left:right], nums[left:right])
	l1, l2 := left, mid
	current := left

	for l1 < mid && l2 < right {
		if temp[l1] <= temp[l2] {
			nums[current] = temp[l1]
			l1 += 1
		} else {
			nums[current] = temp[l2]
			l2 += 1
		}
		current += 1
	}

	for l1 < mid {
		nums[current] = temp[l1]
		l1 += 1
		current += 1
	}

	for l2 < right {
		nums[current] = temp[l2]
		l2 += 1
		current += 1
	}
}
