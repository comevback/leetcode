# 977. Squares of a Sorted Array
Solved
Easy
Topics
Companies
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.


Example 1:
> Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
> Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.
 

Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

---

# Code
```go
package main

import "fmt"

func main() {
	arr := []int{-1}
	res := sortedSquares(arr)
	fmt.Println(res)
}

func sortedSquares(nums []int) []int {
	// 如果数组第一个元素大于等于0，那么直接平方后返回
	if nums[0] >= 0 {
		for i := range nums {
			nums[i] = nums[i] * nums[i]
		}
		return nums
	}
	// 如果数组最后一个元素小于等于0，那么先反转数组，然后平方后返回
	if nums[len(nums)-1] <= 0 {
		reverse(nums)
		for i := range nums {
			nums[i] = nums[i] * nums[i]
		}
		return nums
	}

	// 如果数组中有正数和负数，那么找到正数和负数的分界点，然后分别向两边遍历，比较大小，小的先平方后加入结果数组，最后返回结果数组
	res := []int{}
	index := 0
	for nums[index] < 0 {
		index += 1
	}

	left := index - 1
	right := index

	for left >= 0 && right < len(nums) {
		if -nums[left] < nums[right] {
			res = append(res, nums[left]*nums[left])
			left -= 1
		} else {
			res = append(res, nums[right]*nums[right])
			right += 1
		}
	}

	for left >= 0 {
		res = append(res, nums[left]*nums[left])
		left -= 1
	}

	for right < len(nums) {
		res = append(res, nums[right]*nums[right])
		right += 1
	}

	return res
}

func reverse(nums []int) {
	head, tail := 0, len(nums)-1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head += 1
		tail -= 1
	}
}
```