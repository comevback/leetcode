# 169. Majority Element

Easy

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

Example 1:
> Input: nums = [3,2,3]
Output: 3

Example 2:
> Input: nums = [2,2,1,1,1,2,2]
Output: 2
 

Constraints:
> n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
 

Follow-up: Could you solve the problem in linear time and in O(1) space?

## My first Try (√) 用map计算出现次数 O(n) O(n)

```go
// 1.用map来保存次数
func majorityElement(nums []int) int {
	marjo := len(nums) / 2
	var hsMap = make(map[int]int)
	for _, v := range nums {
		hsMap[v] += 1
		if hsMap[v] > marjo {
			return v
		}
	}
	return 0
}
```

## Boyer-Moore 投票算法 O(n) O(1)

```go
// 2.Boyer-Moore 投票算法
func majorityElement1(nums []int) int {
	cand := 0
	count := 0
	// 如果当前count是0，设当前元素为候选人
	for _, v := range nums {
		if count == 0 {
			cand = v
		}
		// 如果当前元素是候选人，票数加一，反之减一
		if v == cand {
			count += 1
		} else {
			count -= 1
		}
	}
	// 最后还能撑到当选的，就过半数了
	return cand
}
```