# 35. Search Insert Position

Solved
Easy
Topics
Companies
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104

---

# Code

```go
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
```
