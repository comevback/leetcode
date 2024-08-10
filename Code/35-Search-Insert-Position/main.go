package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7}
	res := searchInsert(nums, 6)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for right >= left {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	// if left >= len(nums) {
	// 	return left
	// } else if right < 0 {
	// 	return 0
	// } else if nums[right] > target {
	// 	return right
	// } else if nums[left] < target {
	// 	return left + 1
	// } else {
	// 	return left
	// }

	// 只需要返回left即可覆盖所有情况
	return left
}
