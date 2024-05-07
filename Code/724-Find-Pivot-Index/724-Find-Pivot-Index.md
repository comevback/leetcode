# 724. Find Pivot Index

Easy

Hint
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.


Example 1:
> Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

Example 2:
> Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.

Example 3:
> Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0
 

Constraints:
> 1 <= nums.length <= 104
-1000 <= nums[i] <= 1000
 

Note: This question is the same as 1991: https://leetcode.com/problems/find-the-middle-index-in-array/

---

# Code
```go
package main

func main() {

}

// 1. map
// 一次遍历，map储存所有前后和，二次遍历，对比所有元素的前后和，有相等的就返回该index，无则返回-1
func pivotIndex(nums []int) int {
	// 定义前后和
	var preSum int
	var backSum int
	// 定义每个元素的前后和map
	var preMap map[int]int = make(map[int]int, len(nums))
	var backMap map[int]int = make(map[int]int, len(nums))
	// 双指针i,j分别指向数组头尾
	i, j := 0, len(nums)-1

	// 遍历数组，计算每个元素的前后和
	for i < len(nums)-1 {
		// 前和等于之前的前和+当前值
		preSum += nums[i]
		// 后和等于之前的后和+当前值
		backSum += nums[j]
		// map储存前后和，因为头尾元素前后和都是0，不包括自身，所以map的记录是从头尾第二个元素开始
		preMap[i+1] = preSum
		backMap[j-1] = backSum

		// 双指针移动
		i += 1
		j -= 1
	}

	// 遍历数组，对比每个元素在map中的前后和
	// 有相等的就返回该index
	for i := 0; i < len(nums); i++ {
		if preMap[i] == backMap[i] {
			return i
		}
	}

	// 无则返回-1
	return -1
}

// 2.一次遍历求和
// 二次遍历直接对比前后和
func pivotIndex1(nums []int) int {

	// 求整个数组的和
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 定义前和和后和
	pre := 0
	back := sum

	// 遍历数组，对比每个元素的前后和
	for i, v := range nums {
		// 后和等于之前的后和-当前值
		back -= v
		if pre == back {
			return i
		}
		// 下一个元素的前和才加当前元素值
		pre += v
	}

	// 无则返回-1
	return -1
}
```