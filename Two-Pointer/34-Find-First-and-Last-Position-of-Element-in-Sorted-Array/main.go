package main

import "fmt"

func main() {
	nums := []int{2, 2}         // 测试数组
	res := searchRange(nums, 3) // 调用 searchRange 函数查找目标值3
	fmt.Println(res)            // 打印结果
}

// searchRange 返回目标值在给定数组中的起始和结束位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1} // 如果数组为空，返回 [-1, -1]
	}
	res := bounds(nums, target) // 调用 bounds 函数获取边界
	return res
}

// findIndex 使用递归二分查找法找到目标值的最左和最右索引
func findIndex(nums []int, target int, left int, right int) (int, int) {
	if left > right {
		return -1, -1 // 基本情况：如果左索引大于右索引，返回 [-1, -1]
	}
	var leftIdx, rightIdx int
	mid := (left + right) / 2 // 计算中间索引
	if nums[mid] < target {
		return findIndex(nums, target, mid+1, right) // 目标值在右半部分
	} else if nums[mid] > target {
		return findIndex(nums, target, left, mid-1) // 目标值在左半部分
	} else if nums[mid] == target {
		leftIdx, rightIdx = mid, mid
		for leftIdx > 0 && nums[leftIdx-1] == target {
			leftIdx -= 1 // 向左扩展，直到找到第一个不等于目标值的元素
		}

		for rightIdx < len(nums)-1 && nums[rightIdx+1] == target {
			rightIdx += 1 // 向右扩展，直到找到最后一个不等于目标值的元素
		}

		return leftIdx, rightIdx // 返回最左和最右的索引
	}

	return -1, -1 // 其他情况，返回 [-1, -1]
}

// bounds 使用两次二分查找定位目标值的左右边界
func bounds(nums []int, target int) []int {
	left1, right1, left2, right2 := 0, len(nums), 0, len(nums)

	// 第一次查找，寻找左边界
	for left1 < right1 {
		mid := left1 + (right1-left1)/2
		if nums[mid] == target {
			right1 = mid // 收缩右边界
		} else if nums[mid] < target {
			left1 = mid + 1 // 目标在右侧，移动左边界
		} else if nums[mid] > target {
			right1 = mid // 目标在左侧，收缩右边界
		}
	}

	// 第二次查找，寻找右边界
	for left2 < right2 {
		mid := left2 + (right2-left2)/2
		if nums[mid] == target {
			left2 = mid + 1 // 目标值在右半部分，移动左边界
		} else if nums[mid] < target {
			left2 = mid + 1 // 目标在右侧，移动左边界
		} else if nums[mid] > target {
			right2 = mid // 目标在左侧，收缩右边界
		}
	}

	if left1 == len(nums) || nums[left1] != target {
		return []int{-1, -1} // 如果没有找到目标值，返回 [-1, -1]
	}

	return []int{left1, left2 - 1} // 返回左边界和右边界的索引
}

// =============================================================================
// review 6.17
func searchRange1(nums []int, target int) []int {
	left, right := 0, len(nums)
	res := []int{}

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left != len(nums) && nums[left] == target {
		res = append(res, left)
	} else {
		return []int{-1, -1}
	}

	left, right = 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	res = append(res, left)
	return res
}
