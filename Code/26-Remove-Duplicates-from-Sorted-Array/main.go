package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	// 设下标为1，当前元素currentNum为nums[0]，因为第一个元素肯定是要保存的，所以修改后的nums[k]从第二个元素开始查询
	k := 1
	var currentNum int = nums[0]

	// 循环nums，如果这个元素和当前元素不等，放入nums[k]中，同时把这个元素变成当前元素，k+1，如果相同则略过，
	for _, val := range nums {
		if currentNum != val {
			currentNum = val
			nums[k] = val
			k += 1
		}
	}

	// 另一种循环方式
	// for i := 0; i < len(nums); i++ {
	// 	if currentNum != nums[i] {
	// 		currentNum = nums[i]
	// 		nums[k] = nums[i]
	// 		k += 1
	// 	}

	// 	fmt.Println(nums)
	// }

	// 返回k的，也就是修改后的切片长度
	return k
}

// 循环所有可能的数，如果这个数出现在nums里，则添加进nums[k]
func removeDuplicates2(nums []int) int {

	k := 0

	for i := -100; i <= 100; i++ {
		// 用二分查找法
		if binSearch(nums, i) != -1 {
			nums[k] = i
			k += 1
		}
	}

	return k
}

// 二分查找法实现
func binSearch(nums []int, target int) int {

	// 定义一左一右两个下标
	l, r := 0, len(nums)-1

	// 如果左下标还小于右下标，继续循环
	for l <= r {
		// 得到中间下标
		m := (l + (r - l)) / 2
		// 如果中间下标对应的值等于要找的值，直接返回中间下标
		if target == nums[m] {
			return m
			// 如果中间下标对应的值小于要找的值，要找的值可能位于左侧，所以右下标变为中间下标-1
		} else if target < nums[m] {
			r = m - 1
			// 如果中间下标对应的值大于要找的值，要找的值可能位于右侧，所以左下标变为中间下标+1
		} else {
			l = m + 1
		}
	}

	// 如果循环结束还找不到，返回-1
	return -1
}
