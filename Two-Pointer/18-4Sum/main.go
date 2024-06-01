package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(nums, 0, 0)
	fmt.Println(res)
}

// fourSum 四数之和
func fourSum(nums []int, start int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	QuickSort(nums)

	res := [][]int{}
	pre := math.MinInt

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if i > 0 && val == pre {
			continue
		}
		preRes := threeSum(nums, i+1, target-nums[i])
		for i := range preRes {
			preRes[i] = append(preRes[i], val)
		}
		res = append(res, preRes...)
		pre = nums[i]
	}
	return res
}

// threeSum 三数之和
func threeSum(nums []int, start int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// QuickSort(nums)

	res := [][]int{}
	pre := math.MinInt

	for i := start; i < len(nums); i++ {
		val := nums[i]
		if i > 0 && val == pre {
			continue
		}
		twoRes := TwoSum(nums, i+1, target-nums[i])
		for i := range twoRes {
			twoRes[i] = append(twoRes[i], val)
		}
		res = append(res, twoRes...)
		pre = nums[i]
	}
	return res
}

// TwoSum 两数之和
func TwoSum(nums []int, start int, target int) [][]int {
	left, right := start, len(nums)-1
	res := [][]int{}

	for left < right {
		if nums[left]+nums[right] < target {
			now := nums[left]
			for left < right && nums[left] == now {
				left += 1
			}
		} else if nums[left]+nums[right] > target {
			now := nums[right]
			for left < right && nums[right] == now {
				right -= 1
			}
		} else {
			res = append(res, []int{nums[left], nums[right]})
			nowLeft := nums[left]
			for left < right && nums[left] == nowLeft {
				left += 1
			}
			nowRight := nums[right]
			for left < right && nums[right] == nowRight {
				right -= 1
			}
		}
	}

	return res
}

// ==================================================================================================================
// QuickSort 快速排序
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	key := len(nums) - 1
	head := 0

	for head < key {
		if nums[head] > nums[key] {
			if head+1 == key {
				nums[head], nums[key] = nums[key], nums[head]
			} else {
				nums[head], nums[key], nums[key-1] = nums[key-1], nums[head], nums[key]
			}
			key -= 1
		} else {
			head += 1
		}
	}

	QuickSort(nums[:key])
	if key+1 < len(nums) {
		QuickSort(nums[key+1:])
	}
}
