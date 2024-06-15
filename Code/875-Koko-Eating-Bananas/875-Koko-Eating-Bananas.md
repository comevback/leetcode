# 875. Koko Eating Bananas
Solved
Medium
Topics
Companies
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

 

Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23
 

Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109

# Code
```go
package main

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	// sort.Ints(piles)  // 如果排序，right 可以直接取到piles里的最大值，也就是最后一个元素
	left := 1
	right := 1000000001 // 10^9 + 1 因为题目要求的是1 <= piles[i] <= 10^9，如果之前排序了，这里可以直接取到piles里的最大值

	// 使用二分查找来确定最小的合适速度
	for left < right {
		mid := left + (right-left)/2 // 计算中间速度
		// 调用 countHours 来计算以当前速度 mid 吃完所有香蕉所需的总小时数
		if countHours(piles, mid) <= h {
			right = mid // 如果当前速度所需小时数小于等于h，试图找到更小的速度，将搜索范围减半
		} else {
			left = mid + 1 // 如果当前速度所需小时数大于h，增加速度
		}
	}

	return left // 返回最小速度
}

// countHours 计算给定速度 k 下，吃完所有香蕉堆所需的总小时数
func countHours(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i] / k // 计算当前堆的香蕉数除以速度得到的整数小时数
		if nums[i]%k > 0 {
			res += 1 // 如果有余数，意味着还需要额外一个小时来吃完这堆的剩余香蕉
		}
	}
	return res // 返回总小时数
}
```