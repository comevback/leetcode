217. Contains Duplicate

Easy

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

Example 1:
> Input: nums = [1,2,3,1]
Output: true

Example 2:
> Input: nums = [1,2,3,4]
Output: false

Example 3:
> Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
 

Constraints:
> 1 <= nums.length <= 105
-109 <= nums[i] <= 109

---

```go
package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 5, 6, 2, 7, 9, 5, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

// ================================================================================================================================
// 1.用哈希map，把每个遇到的值都存进去，每次遍历新的值时查找是否已存在map里，如果有，返回true
func containsDuplicate(nums []int) bool {
	var hsMap map[int]bool = make(map[int]bool)

	for _, v := range nums {
		if hsMap[v] {
			return true
		}
		hsMap[v] = true
	}

	return false
}

// ================================================================================================================================
// 2.暴力循环 O(n^2) - O(1)
func containsDuplicate1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[i] == nums[j] && i != j {
				return true
			}
		}
	}
	return false
}

// ================================================================================================================================
// 3.排序后查找相邻数是否一样, O(nlogn) - O(1)
// func containsDuplicate2(nums []int) bool {
// 	nums = sort(nums)
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] == nums[i+1] {
// 			return true
// 		}
// 	}
// 	return false
// }

```