# 503. Next Greater Element II

Solved
Medium
Topics
Companies
Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.

Example 1:

Input: nums = [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number.
The second 1's next greater number needs to search circularly, which is also 2.
Example 2:

Input: nums = [1,2,3,4,3]
Output: [2,3,4,-1,4]

Constraints:

1 <= nums.length <= 104
-109 <= nums[i] <= 109

---

# Code

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3}
	res := nextGreaterElements(nums)
	fmt.Println(res)
}

func nextGreaterElements(nums []int) []int {
	nextHigher := make([]int, len(nums)) // 创建一个数组来保存每个元素的下一个更大元素
	temp := []int{}                      // 使用一个栈来帮助找出下一个更大的元素

	for i := len(nums) - 1; i >= 0; i-- { // 从后向前遍历数组
		for len(temp) > 0 && nums[i] >= temp[len(temp)-1] {
			temp = temp[:len(temp)-1] // 弹出栈中小于等于当前元素的所有元素
		}

		if len(temp) == 0 {
			nextHigher[i] = -1 // 如果栈为空，则当前元素没有更大的下一个元素
		} else {
			nextHigher[i] = temp[len(temp)-1] // 栈顶元素是下一个更大的元素
		}
		temp = append(temp, nums[i]) // 当前元素入栈
	}

	// 重复一次相同的过程，但不正确处理环形数组的特性
	for i := len(nums) - 1; i >= 0; i-- {
		for len(temp) > 0 && nums[i] >= temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		}

		if len(temp) == 0 {
			nextHigher[i] = -1
		} else {
			nextHigher[i] = temp[len(temp)-1]
		}
		temp = append(temp, nums[i])
	}

	return nextHigher // 返回结果
}

// 将数组长度扩展为两倍进行模拟。使用取模运算
func nextGreaterElements1(nums []int) []int {
	n := len(nums)        // 获取数组的长度
	res := make([]int, n) // 创建结果数组
	s := make([]int, 0)   // 使用栈来存储元素

	// 循环两倍长度来模拟环形数组
	for i := 2*n - 1; i >= 0; i-- {
		// 通过取模操作来处理环形数组特性
		for len(s) > 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1] // 弹出栈中小于等于当前元素的所有元素
		}

		if len(s) == 0 {
			res[i%n] = -1 // 如果栈为空，则当前元素没有更大的下一个元素
		} else {
			res[i%n] = s[len(s)-1] // 栈顶元素是下一个更大的元素
		}

		s = append(s, nums[i%n]) // 当前元素入栈
	}

	return res // 返回结果
}
```
